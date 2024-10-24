package ds

import "sync"

type ConcurrentMap[K comparable, V any] struct {
	m  map[K]V
	mu sync.RWMutex
}

func NewConcurrentMap[K comparable, V any]() *ConcurrentMap[K, V] {
	return &ConcurrentMap[K, V]{
		m: make(map[K]V),
	}
}

func (cm *ConcurrentMap[K, V]) Get(key K) (V, bool) {
	cm.mu.RLock()
	defer cm.mu.Unlock()

	v, ok := cm.m[key]
	return v, ok
}

func (cm *ConcurrentMap[K, V]) Set(key K, value V) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	cm.m[key] = value
}

func (cm *ConcurrentMap[K, V]) Del(key K) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	delete(cm.m, key)
}

func (cm *ConcurrentMap[K, V]) Len() int {
	cm.mu.RLock()
	defer cm.mu.Unlock()

	return len(cm.m)
}
