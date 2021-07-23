package main

import "sync"

type MemTable interface {
	Set(key, value string)
	Get(key string) (string, bool)
	Del(key string) bool

	HasImmutable() bool
	GetImmutable() MemData
}

const (
	deleteValue   = ""
	maxSize       = 1024 * 1024
	dataPartition = 4
)

type mem struct {
	sync.Mutex

	Data          MemData
	ImmutableData MemData

	MaxSize       int
	DataPartition int
}

func (m *mem) Set(key, value string) {
	defer m.AfterEvent()
	m.Data.Set(key, value)
}

func (m *mem) Get(key string) (string, bool) {
	return m.Data.Get(key)
}

func (m *mem) Del(key string) bool {
	defer m.AfterEvent()
	oldValue, exist := m.Data.Set(key, deleteValue)
	return exist && oldValue != deleteValue
}

func (m *mem) HasImmutable() bool {
	return m.ImmutableData != nil
}

func (m *mem) GetImmutable() MemData {
	return m.ImmutableData
}

func (m *mem) AfterEvent() {
	if m.HasImmutable() {
		return
	} else if m.MaxSize <= m.Data.Size() {
		return
	}

	m.Lock()
	defer m.Unlock()
	if m.HasImmutable() {
		return
	}

	m.ImmutableData = m.Data
	m.Data = NewMemData(m.DataPartition)
}

func NewMemTable() MemTable {
	m := &mem{
		Data:          NewMemData(dataPartition),
		MaxSize:       maxSize,
		DataPartition: dataPartition,
	}
	return m
}
