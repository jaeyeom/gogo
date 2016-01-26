package conmap

import "sync"

type Resource int

type Accessor struct {
	R *Resource
	L *sync.Mutex
}

func (acc *Accessor) Use() {
	// do something
	acc.L.Lock()
	// Use acc.R
	acc.L.Unlock()
	// Do something else
}

type ConcurrentMap struct {
	M map[string]string
	L *sync.RWMutex
}

func (m ConcurrentMap) Get(key string) string {
	m.L.RLock()
	defer m.L.RUnlock()
	return m.M[key]
}

func (m ConcurrentMap) Set(key, value string) {
	m.L.Lock()
	m.M[key] = value
	m.L.Unlock()
}
