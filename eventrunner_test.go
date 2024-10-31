package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

type eventOpen struct {
	session    *Session
	sessionIDs *[]string
	wg         *sync.WaitGroup
	mu         *sync.Mutex
}

func (e *eventOpen) Handle() {
	defer e.mu.Unlock()

	e.mu.Lock()
	*e.sessionIDs = append(*e.sessionIDs, e.session.id)
	e.wg.Done()
}

func TestEventRunner(t *testing.T) {
	er := NewEventRunner(1024)
	defer er.Close()

	go er.Run()

	userCount := 1000

	wg := &sync.WaitGroup{}
	wg.Add(userCount)

	mu := &sync.Mutex{}

	sessionIDs := make([]string, 0, userCount)

	for i := 0; i < userCount; i++ {
		session := NewSession(nil, 1024)
		er.Send(&eventOpen{
			session,
			&sessionIDs,
			wg,
			mu,
		})
	}

	wg.Wait()

	require.Equal(t, len(sessionIDs), userCount)
}
