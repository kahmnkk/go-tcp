package ds

import "sync"

type EventLoop[T any] struct {
	queue   chan T
	handler func(T)
	wg      sync.WaitGroup
}

func NewEventLoop[T any](queueCapacity int) *EventLoop[T] {
	e := &EventLoop[T]{
		queue:   make(chan T, queueCapacity),
		handler: func(T) {},
	}
	return e
}

func (e *EventLoop[T]) Run() {
	e.wg.Add(1)
	go func() {
		defer e.wg.Done()

		for event := range e.queue {
			e.handler(event)
		}
	}()
}

func (e *EventLoop[T]) Close() {
	close(e.queue)
	e.wg.Wait()
}

func (e *EventLoop[T]) Send(event T) {
	e.queue <- event
}

func (e *EventLoop[T]) SetHandler(handler func(T)) {
	e.handler = handler
}
