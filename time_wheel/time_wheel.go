package time_wheel

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

type TimeWheel struct {
	interval    time.Duration
	slotNums    int
	currentPos  int
	ticker      *time.Ticker
	slots       []*list.List
	taskRecords *sync.Map
	isRunning   bool
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

func (tw *TimeWheel) Start() {
	tw.ticker = time.NewTicker(tw.interval)
	go tw.start()
	tw.isRunning = true
}

func (tw *TimeWheel) start() {
	for {
		select {
		case <-tw.ticker.C:
			tw.checkAndRunTask()
		}
	}
}

// 添加任务的内部函数
// @param task       Task  Task对象
// @param byInterval bool  生成Task在时间轮盘位置和圈数的方式，true表示利用Task.interval来生成，false表示利用Task.createTime生成
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

func (tw *TimeWheel) getPosAndCircleByInterval(d time.Duration) (int, int) {
	delaySeconds := int(d.Seconds())
	intervalSeconds := int(tw.interval.Seconds())
	circle := delaySeconds / intervalSeconds / tw.slotNums
	pos := (tw.currentPos + delaySeconds/intervalSeconds) % tw.slotNums

	// 计算的位置和当前位置重叠时，因为当前位置已经走过了，circle需要减一
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

	// 计算的位置和当前位置重叠时，因为当前位置已经走过了，circle需要减一
	if pos == tw.currentPos && circle != 0 {
		circle--
	}
	return pos, circle
}

func (tw *TimeWheel) removeTask(task *Task) {
	val, _ := tw.taskRecords.Load(task.key)
	tw.taskRecords.Delete(task.key)
	currentList := tw.slots[task.pos]
	currentList.Remove(val.(*list.Element))
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

			} else {
				tw.taskRecords.Delete(task.key)
			}
		}
	}

	if tw.currentPos == tw.slotNums-1 {
		tw.currentPos = 0
	} else {
		tw.currentPos++
	}
}
