package main

import (
	"gotcp/internal/ds"
)

type EventRunner struct {
	eventLoop *ds.EventLoop[Event]
}

func NewEventRunner(queueCapacity int) *EventRunner {
	runner := &EventRunner{}
	runner.eventLoop = ds.NewEventLoop[Event](queueCapacity)
	runner.eventLoop.SetHandler(runner.Handler)

	return runner
}

func (eh *EventRunner) Handler(event Event) {
	event.Handle()
}

func (eh *EventRunner) Run() {
	eh.eventLoop.Run()
}

func (eh *EventRunner) Close() {
	eh.eventLoop.Close()
}

func (eh *EventRunner) Send(event Event) {
	eh.eventLoop.Send(event)
}
