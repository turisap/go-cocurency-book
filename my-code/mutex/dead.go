package main

import "sync"

var m sync.Mutex

func f() {

	m.Lock()
	defer m.Unlock()
	// process
}

func g() {
	m.Lock()
	defer m.Unlock()
	go f() // Deadlock (FIXED with GO)
}

func main() {
	g()
}
