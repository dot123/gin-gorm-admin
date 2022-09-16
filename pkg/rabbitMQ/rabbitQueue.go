package rabbitMQ

import (
	"sync"
	"sync/atomic"
)

/**
rabbitMq 通用队列
*/
type Node struct {
	data *rChannel
	next *Node
}

type ChannelQueue struct {
	head  *Node
	end   *Node
	l     int32
	rlock sync.Mutex
}

func NewChannelQueue() *ChannelQueue {
	q := &ChannelQueue{head: nil, end: nil, l: 0}
	return q
}

func (q *ChannelQueue) Add(data *rChannel) {
	q.rlock.Lock()
	defer q.rlock.Unlock()
	n := &Node{data: data, next: nil}
	atomic.AddInt32(&q.l, 1)
	if q.end == nil {
		q.head = n
		q.end = n
	} else {
		q.end.next = n
		q.end = n
	}
	return
}

func (q *ChannelQueue) Pop() (*rChannel, bool) {
	q.rlock.Lock()
	defer q.rlock.Unlock()
	if q.head == nil {
		return nil, false
	}
	atomic.AddInt32(&q.l, -1)
	data := q.head.data
	q.head = q.head.next
	if q.head == nil {
		q.end = nil
	}
	return data, true
}

func (q *ChannelQueue) Count() int32 {
	return q.l
}
