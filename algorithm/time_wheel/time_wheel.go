package timewheel

import (
	"container/list"
	"errors"
	"fmt"
	"sync"
	"time"
)

type TimeWheel struct {
	interval          time.Duration
	slots             []*list.List
	ticker            *time.Ticker
	currentPos        int
	slotNums          int
	addTaskChannel    chan *Task
	removeTaskChannel chan *Task
	stopChannel       chan bool
	taskRecords       *sync.Map
	job               Job
	isRunning         bool
}

type Job func(interface{})

type Task struct {
	key         interface{}
	interval    time.Duration
	createdTime time.Time
	pos         int
	circle      int
	job         Job
	times       int
}

var tw *TimeWheel
var once sync.Once

// ErrDuplicateTaskKey is an error for duplicate task key
var ErrDuplicateTaskKey = errors.New("Duplicate task key")

// ErrTaskKeyNotFount is an error when task key is not found
var ErrTaskKeyNotFount = errors.New("Task key doesn't existed in task list, please check your input")

func CreateTimeWheel(interval time.Duration, slotNums int, job Job) *TimeWheel {
	once.Do(func() {
		tw = New(interval, slotNums, job)
	})
	return tw
}

func GetTimeWheel() *TimeWheel {
	return tw
}

func New(interval time.Duration, slotNums int, job Job) *TimeWheel {
	if interval <= 0 || slotNums <= 0 {
		return nil
	}
	tw := &TimeWheel{
		interval:          interval,
		slots:             make([]*list.List, slotNums),
		currentPos:        0,
		slotNums:          slotNums,
		addTaskChannel:    make(chan *Task),
		removeTaskChannel: make(chan *Task),
		stopChannel:       make(chan bool),
		taskRecords:       &sync.Map{},
		job:               job,
		isRunning:         false,
	}

	tw.initSlots()
	return tw
}

func (tw *TimeWheel) Start() {
	tw.ticker = time.NewTicker(tw.interval)
	go tw.start()
	tw.isRunning = true
}

func (tw *TimeWheel) Stop() {
	tw.stopChannel <- true
	tw.isRunning = false
}

func (tw *TimeWheel) IsRunning() bool {
	return tw.isRunning
}

func (tw *TimeWheel) AddTask(interval time.Duration, key interface{}, createdTime time.Time, times int, job Job) error {
	if interval <= 0 || key == nil {
		return errors.New("Invalid task params")
	}

	// 检查Task.Key是否已经存在
	_, ok := tw.taskRecords.Load(key)
	if ok {
		return ErrDuplicateTaskKey
	}

	tw.addTaskChannel <- &Task{
		key:         key,
		interval:    interval,
		createdTime: createdTime,
		job:         job,
		times:       times,
	}

	return nil
}

func (tw *TimeWheel) RemoveTask(key interface{}) error {
	if key == nil {
		return nil
	}

	// 检查该Task是否存在
	val, ok := tw.taskRecords.Load(key)
	if !ok {
		return ErrTaskKeyNotFount
	}

	task := val.(*list.Element).Value.(*Task)
	tw.removeTaskChannel <- task
	return nil
}

func (tw *TimeWheel) initSlots() {
	for i := 0; i < tw.slotNums; i++ {
		tw.slots[i] = list.New()
	}
}

func (tw *TimeWheel) start() {
	for {
		select {
		case <-tw.ticker.C:
			tw.checkAndRunTask()
		case task := <-tw.addTaskChannel:
			tw.addTask(task, false)
		case task := <-tw.removeTaskChannel:
			tw.removeTask(task)
		case <-tw.stopChannel:
			tw.ticker.Stop()
			return
		}
	}
}

func (tw *TimeWheel) checkAndRunTask() {
	currentList := tw.slots[tw.currentPos]

	if currentList != nil {
		for item := currentList.Front(); item != nil; {
			task := item.Value.(*Task)
			if task.circle > 0 {
				task.circle--
				item = item.Next()
				continue
			}

			if task.job != nil {
				go task.job(task.key)
			} else if tw.job != nil {
				go tw.job(task.key)
			} else {
				fmt.Println(fmt.Sprintf("The task %d don't have job to run", task.key))
			}

			next := item.Next()
			tw.taskRecords.Delete(task.key)
			currentList.Remove(item)
			item = next

			if task.times != 0 {
				if task.times < 0 {
					tw.addTask(task, true)
				} else {
					task.times--
					tw.addTask(task, true)
				}

			}
		}
	}

	if tw.currentPos == tw.slotNums-1 {
		tw.currentPos = 0
	} else {
		tw.currentPos++
	}
}

func (tw *TimeWheel) addTask(task *Task, byInterval bool) {
	var pos, circle int
	if byInterval {
		pos, circle = tw.getPosAndCircleByInterval(task.interval)
	} else {
		pos, circle = tw.getPosAndCircleByCreatedTime(task.createdTime, task.interval, task.key)
	}

	task.circle = circle
	task.pos = pos

	element := tw.slots[pos].PushBack(task)
	tw.taskRecords.Store(task.key, element)
}

func (tw *TimeWheel) removeTask(task *Task) {
	val, _ := tw.taskRecords.Load(task.key)
	tw.taskRecords.Delete(task.key)

	currentList := tw.slots[task.pos]
	currentList.Remove(val.(*list.Element))
}

func (tw *TimeWheel) getPosAndCircleByInterval(d time.Duration) (int, int) {
	delaySeconds := int(d.Seconds())
	intervalSeconds := int(tw.interval.Seconds())
	circle := delaySeconds / intervalSeconds / tw.slotNums
	pos := (tw.currentPos + delaySeconds/intervalSeconds) % tw.slotNums
	if pos == tw.currentPos && circle != 0 {
		circle--
	}
	return pos, circle
}

func (tw *TimeWheel) getPosAndCircleByCreatedTime(createdTime time.Time, d time.Duration, key interface{}) (int, int) {
	passedTime := time.Since(createdTime)
	passedSeconds := int(passedTime.Seconds())
	delaySeconds := int(d.Seconds())
	intervalSeconds := int(tw.interval.Seconds())

	circle := delaySeconds / intervalSeconds / tw.slotNums
	pos := (tw.currentPos + (delaySeconds-(passedSeconds%delaySeconds))/intervalSeconds) % tw.slotNums

	if pos == tw.currentPos && circle != 0 {
		circle--
	}

	return pos, circle
}
