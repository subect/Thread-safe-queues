package main

import "sync"

type SliceQueue struct {
	queue []interface{}
	mu    sync.Mutex
}

// NewSliceQueue 初始化
func (safeQueue *SliceQueue) NewSliceQueue(n int) (q *SliceQueue) {
	return &SliceQueue{
		queue: make([]interface{}, 0, n),
	}
}

// EnQueue 入队
func (safeQueue *SliceQueue) EnQueue(v interface{}) {
	safeQueue.mu.Lock()
	defer safeQueue.mu.Unlock()
	safeQueue.queue = append(safeQueue.queue, v)
}

// OutQueue 出队
func (safeQueue *SliceQueue) OutQueue() interface{} {
	safeQueue.mu.Lock()
	defer safeQueue.mu.Unlock()
	if len(safeQueue.queue) == 0 {
		return nil
	}
	v := safeQueue.queue[0]
	safeQueue.queue = safeQueue.queue[1:]
	return v
}
