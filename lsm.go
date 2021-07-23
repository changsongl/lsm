package main

type LSM interface {
	Set(key, value string)
	Get(key string) (string, bool)
	Del(key string) bool
}

type lsm struct {
	mem MemTable
}

func NewLSM() LSM {
	l := &lsm{
		mem: NewMemTable(),
	}

	return l
}

func (l *lsm) Set(key, value string) {
	l.mem.Set(key, value)
}

func (l *lsm) Get(key string) (string, bool) {
	return l.mem.Get(key)
}

func (l *lsm) Del(key string) bool {
	return l.mem.Del(key)
}
