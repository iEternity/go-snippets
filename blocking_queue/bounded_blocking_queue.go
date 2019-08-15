package blocking_queue

import (
	"container/list"
	"sync"
)

type BlockingQueue interface {
	Put(interface{})
	Take() interface{}
}

type BlockingQueueWithChan struct {
	ch chan interface{}
}

func NewQueueWithChan(maxSize uint32) BlockingQueue {
	return &BlockingQueueWithChan{ch: make(chan interface{}, maxSize)}
}

func (self *BlockingQueueWithChan) Put(v interface{}){
	self.ch <- v
}

func (self *BlockingQueueWithChan) Take() interface{}{
	return <-self.ch
}

type BlockingQueueWithCond struct{
	queue *list.List
	fullCondition *sync.Cond
	emptyCondition *sync.Cond
	maxSize uint32
	isFull bool
	isEmpty bool
}

func NewQueueWithCond(maxSize uint32) BlockingQueue {
	lock := &sync.Mutex{}
	return &BlockingQueueWithCond{
		queue: list.New(),
		fullCondition: sync.NewCond(lock),
		emptyCondition: sync.NewCond(lock),
		maxSize: maxSize,
		isFull: false,
		isEmpty: true,
	}
}

func (self *BlockingQueueWithCond) Put(v interface{}) {
	self.fullCondition.L.Lock()
	defer self.fullCondition.L.Unlock()

	for self.isFull {
		self.fullCondition.Wait()
	}

	self.queue.PushBack(v)
	self.isEmpty = false

	if self.queue.Len() >= int(self.maxSize){
		self.isFull = true
	}

	self.emptyCondition.Signal()
}

func (self *BlockingQueueWithCond) Take() interface{} {
	self.emptyCondition.L.Lock()
	defer self.emptyCondition.L.Unlock()

	for self.isEmpty {
		self.emptyCondition.Wait()
	}

	element := self.queue.Front()
	self.queue.Remove(element)

	self.isFull = false
	if self.queue.Len() == 0 {
		self.isEmpty = true
	}

	self.fullCondition.Signal()

	return element.Value
}