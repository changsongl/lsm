package main

import (
	"github.com/OneOfOne/xxhash"
	"sync"
)

type MemData interface {
	Size() int
	Set(key, value string) (string, bool)
	Get(key string) (string, bool)
}

type HashFunc func(key string) int

type memData struct {
	Partition int
	size      int

	Maps  []map[string]string
	Locks []sync.RWMutex
	Hash  HashFunc
}

func (m *memData) Set(key, value string) (string, bool) {
	i := m.Hash(key)
	m.Locks[i].Lock()
	defer m.Locks[i].Unlock()

	vLen, kLen := len(value), len(key)
	oldValue, exist := m.Maps[i][key]
	if exist {
		m.size += vLen - len(oldValue)
	} else {
		m.size += vLen + kLen
	}

	m.Maps[i][key] = value

	return oldValue, exist
}

func (m *memData) Get(key string) (string, bool) {
	i := m.Hash(key)
	m.Locks[i].RLock()
	defer m.Locks[i].RUnlock()

	v, exist := m.Maps[i][key]
	return v, exist
}

func (m *memData) Size() int {
	return m.size
}

func NewMemData(partition int) MemData {
	maps := make([]map[string]string, 0, partition)
	for i := 0; i < partition; i++ {
		maps[i] = make(map[string]string)
	}

	data := &memData{
		Partition: partition,
		Maps:      maps,
		Locks:     make([]sync.RWMutex, partition),
		Hash: func(key string) int {
			h := xxhash.New64()
			_, err := h.Write([]byte(key))
			if err != nil {
				//TODO
			}

			return int(h.Sum64() % uint64(partition))
		},
	}
	return data
}
