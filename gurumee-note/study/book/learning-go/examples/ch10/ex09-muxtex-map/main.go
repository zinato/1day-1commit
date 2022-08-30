package main

import (
	"fmt"
	"sync"
)

type MutexScoreboardManager struct {
	l          sync.RWMutex
	scoreboard map[string]int
}

func NewMutexScoreboardManager() *MutexScoreboardManager {
	return &MutexScoreboardManager{
		scoreboard: map[string]int{},
	}
}

func (msm *MutexScoreboardManager) Update(name string, val int) {
	msm.l.Lock()
	defer msm.l.Unlock()

	msm.scoreboard[name] = val
}

func (msm *MutexScoreboardManager) Read(name string) (int, bool) {
	msm.l.RLock()
	defer msm.l.RUnlock()

	val, ok := msm.scoreboard[name]
	return val, ok
}

func main() {
	msm := NewMutexScoreboardManager()

	name := "hello"
	val := 1
	msm.Update(name, val)

	val, ok := msm.Read(name)
	if !ok {
		fmt.Println(name, " is not exist")
	} else {
		fmt.Println(val)
	}

	name = "stay"
	val, ok = msm.Read(name)
	if !ok {
		fmt.Println(name, " is not exist")
	} else {
		fmt.Println(val)
	}
}
