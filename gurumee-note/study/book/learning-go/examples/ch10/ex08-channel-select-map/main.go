package main

import "fmt"

func scoreboardManager(in <-chan func(map[string]int), done <-chan struct{}) {
	scoreboard := map[string]int{}
	for {
		select {
		case <-done:
			return
		case f := <-in:
			f(scoreboard)
		}
	}
}

type ChannelScoreboardManager chan func(map[string]int)

func NewChannelScoreboardManager() (ChannelScoreboardManager, func()) {
	ch := make(ChannelScoreboardManager)
	done := make(chan struct{})
	go scoreboardManager(ch, done)
	return ch, func() {
		close(done)
	}
}

func (csm ChannelScoreboardManager) Update(name string, val int) {
	csm <- func(m map[string]int) {
		m[name] = val
	}
}

func (csm ChannelScoreboardManager) Read(name string) (int, bool) {
	var out int
	var ok bool
	done := make(chan struct{})
	csm <- func(m map[string]int) {
		out, ok = m[name]
		close(done)
	}
	<-done
	return out, ok
}

func main() {
	csm, cancel := NewChannelScoreboardManager()
	defer cancel()

	name := "hello"
	val := 1
	csm.Update(name, val)

	val, ok := csm.Read(name)
	if !ok {
		fmt.Println(name, " is not exist")
	} else {
		fmt.Println(val)
	}

	name = "stay"
	val, ok = csm.Read(name)
	if !ok {
		fmt.Println(name, " is not exist")
	} else {
		fmt.Println(val)
	}
}
