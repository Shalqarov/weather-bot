package pkg

import "sync"

type UMap struct {
	sync.RWMutex
	mp map[int64]struct{}
}

func (m *UMap) Set(k int64) {
	defer m.Unlock()
	m.Lock()
	m.mp[k] = struct{}{}
}

func (m *UMap) Get(k int64) bool {
	defer m.RUnlock()
	m.RLock()
	_, ok := m.mp[k]
	return ok
}

func (m *UMap) GetAllKeys() []int64 {
	defer m.RUnlock()
	m.RLock()
	res := make([]int64, 0, len(m.mp))
	for k, _ := range m.mp {
		res = append(res, k)
	}
	return res
}

func NewUMap(capacity int) *UMap {
	m := &UMap{
		mp: make(map[int64]struct{}, capacity),
	}
	return m
}
