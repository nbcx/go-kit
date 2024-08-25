package co

import (
	"sync"
)

type Options[T any] func(*Task[T])

type Task[T any] struct {
	wg        sync.WaitGroup
	data      chan T
	call      func(T)
	closeChan chan struct{}
}

func NewTask[T any](opts ...Options[T]) *Task[T] {
	h := &Task[T]{
		closeChan: make(chan struct{}, 1),
	}
	h.wg.Add(1) // start add
	for _, v := range opts {
		v(h)
	}
	if h.call != nil {
		h.data = make(chan T, 10)
		go h.receive()
	}
	return h
}

func WithCallback[T any](call func(T)) Options[T] {
	return func(t *Task[T]) {
		t.call = call
	}
}

func (p *Task[T]) receive() {
	defer close(p.data)
	for {
		select {
		case value := <-p.data:
			p.call(value)
			p.wg.Done() // // exec call add end
		case <-p.closeChan:
			return
		}
	}
}

// 接口请求
func (p *Task[T]) Exec(f func() T) {
	p.wg.Add(1) // exec add
	go func() {
		defer p.wg.Done() // exec add end
		t := f()
		if p.call != nil {
			p.wg.Add(1) // exec call add
			p.data <- t
		}
	}()
}

// 等待所有请求完成
func (p *Task[T]) Wait() {
	defer close(p.closeChan)
	p.wg.Done() // start add end
	p.wg.Wait()
	p.closeChan <- struct{}{}
}
