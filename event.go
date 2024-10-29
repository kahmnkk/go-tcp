package main

type Event interface {
	Handle()
}

type EventOpen struct {
	session *Session
}

func (e *EventOpen) Handle() {
	// s.sessionMap.Set(e.session.id, e.session)
}

type EventClose struct {
	session *Session
}

func (e *EventClose) Handle() {
	// s.sessionMap.Del(e.session.id)
}

type EventPacket struct {
	session *Session
	message string
}

func (e *EventPacket) Handle() {
	server.router.RouteMessage(e.session, e.message)
}
