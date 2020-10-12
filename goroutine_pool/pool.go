package main

// Task task of worker
type Task struct {
	function func(...interface{})
	args     []interface{}
}

// Pool pool struct
type Pool struct {
	task    chan Task
	limiter chan struct{}
}

// new worker
func (pool *Pool) newWorker(task Task) {
	defer func() { <-pool.limiter }()

	for {
		task.function(task.args...)
		task = <-pool.task
	}
}

// AddTask add task to pool
func (pool *Pool) AddTask(task Task) {
	select {
	case pool.task <- task:
	case pool.limiter <- struct{}{}:
		go pool.newWorker(task)
	}
}

// New create pool
func New(size int) *Pool {
	return &Pool{
		task:    make(chan Task),
		limiter: make(chan struct{}, size),
	}
}
