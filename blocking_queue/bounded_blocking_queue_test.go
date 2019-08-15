package blocking_queue

import (
	"fmt"
	"testing"
	"time"
)

func produce(q BlockingQueue, i int, interval int){
	start, end := i * 10, (i + 1) * 10
	for i := start; i < end; i++{
		q.Put(i)

		if interval > 0{
			time.Sleep(time.Duration(interval) * time.Second)
		}
	}
}

func consume(q BlockingQueue, interval int){
	for {
		fmt.Println(q.Take().(int))

		if interval > 0{
			time.Sleep(time.Duration(interval) * time.Second)
		}
	}
}

func TestChanQueueFull(t *testing.T) {
	q := NewQueueWithChan(3)

	for i := 0; i < 3; i++{
		go produce(q, i, 0)
	}

	for i := 0; i< 1; i++{
		go consume(q, 1)
	}

	time.Sleep(time.Minute * 1)
}

func TestChanQueueEmpty(t *testing.T) {
	q := NewQueueWithChan(3)

	for i := 0; i < 1; i++{
		go produce(q, i, 1)
	}

	for i := 0; i<10;i++{
		go consume(q, 0)
	}

	time.Sleep(time.Second * 15)
}

func TestChanQueue(t *testing.T){
	q := NewQueueWithChan(3)

	for i := 0; i < 10; i++{
		go produce(q, i, 0)
	}

	for i := 0; i<10;i++{
		go consume(q, 0)
	}

	time.Sleep(time.Second * 10)
}

func TestCondQueueFull(t *testing.T) {
	q := NewQueueWithCond(3)

	for i := 0; i < 3; i++{
		go produce(q, i, 0)
	}

	for i := 0; i< 1; i++{
		go consume(q, 1)
	}

	time.Sleep(time.Minute * 1)
}

func TestCondQueueEmpty(t *testing.T) {
	q := NewQueueWithCond(3)

	for i := 0; i < 1; i++{
		go produce(q, i, 1)
	}

	for i := 0; i<10;i++{
		go consume(q, 0)
	}

	time.Sleep(time.Second * 15)
}

func TestCondQueue(t *testing.T){
	q := NewQueueWithCond(3)

	for i := 0; i < 10; i++{
		go produce(q, i, 0)
	}

	for i := 0; i<10;i++{
		go consume(q, 0)
	}

	time.Sleep(time.Second * 10)
}