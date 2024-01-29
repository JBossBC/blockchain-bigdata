package core

import "sync"

type cachePool struct {
	data map[uintptr]any
	rw   sync.RWMutex
}

func newCachePool() *cachePool {
	var res = new(cachePool)
	res.rw = sync.RWMutex{}
	res.data = make(map[uintptr]any)
	return res
}

func (cp *cachePool) get(ptr uintptr) any {
	cp.rw.RLock()
	defer cp.rw.RUnlock()
	return cp.data[ptr]
}

func (cp *cachePool) query(ptr uintptr) (data any, ok bool) {
	cp.rw.RLock()
	defer cp.rw.RUnlock()
	data, ok = cp.data[ptr]
	return data, ok
}

func (cp *cachePool) put(ptr uintptr, date any) {
	cp.rw.Lock()
	defer cp.rw.Unlock()
	cp.data[ptr] = date
}
