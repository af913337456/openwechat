package xrobot

import (
	"fmt"
	"sync"
	"time"
)

var onceRunQueueMap = sync.Map{}

func mapKey(sender string, receiver string) string {
	return fmt.Sprintf("%s_%s_map", sender, receiver)
}

func GetOnce(key string, nilThenCreate bool) *sync.Once {
	val, ok := onceRunQueueMap.Load(key)
	if ok {
		if val != nil {
			return val.(*sync.Once)
		}
	}
	if nilThenCreate {
		newOnce := &sync.Once{}
		onceRunQueueMap.Store(key, newOnce)
		return newOnce
	}
	return nil
}

var queueMap = sync.Map{}
var queueClearOnce = sync.Once{}

func ClearOlderQueue() {
	queueClearOnce.Do(func() {
		limitSec := 60 * 60 * 24 * 5 // 5 days
		go func() {
			for {
				time.Sleep(time.Minute)
				queueMap.Range(func(key, value any) bool {
					if value != nil {
						queue := value.(*Queue)
						if queue.IsEmpty() && (time.Now().Unix()-queue.CreateTime > int64(limitSec)) {
							queueMap.Delete(key)
						}
					}
					return true
				})
			}
		}()
	})
}

func GetQueue(key string) *Queue {
	val, ok := queueMap.Load(key)
	if ok && val != nil {
		return val.(*Queue)
	}
	newQueue := &Queue{
		CreateTime: time.Now().Unix(),
	}
	queueMap.Store(key, newQueue)
	return newQueue
}

type QueueItem struct {
	// todo note 其他字段在这里扩展
	//DataId     string
	//Receiver   string
	ReplyMsg XRobotMsg
	//MsgContent string
	//MsgTime    string
}
type Queue struct {
	CreateTime int64
	items      []*QueueItem
}

func (q *Queue) Enqueue(item *QueueItem) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() *QueueItem {
	if len(q.items) == 0 {
		return nil
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}
