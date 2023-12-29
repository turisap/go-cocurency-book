package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Cache struct {
	m    sync.RWMutex
	data map[string]Data
}

type Data struct {
	age int
}

var cache = Cache{
	data: map[string]Data{
		"Kirill": Data{age: 34},
		"Alice":  Data{age: 18},
	},
}

func getAge(name string) (cached bool, d *Data) {

	cache.m.Lock()
	data, ok := cache.data[name]
	cache.m.Unlock()

	if ok {
		entry := Data{age: data.age}
		return true, &entry
	}

	cache.m.RLock()
	personAge := getFromDb(name)
	entry := Data{age: personAge}
	cache.data[name] = entry
	cache.m.RUnlock()

	return false, &entry

}

func getFromDb(n string) int {
	fmt.Printf("Getting %s age from db\n", n)
	time.Sleep(time.Millisecond * 500)

	return rand.Intn(80-23+1) + 23
}

func main() {
	names := []string{"Anna", "Alice", "Alex", "Kirill"}
	var wg = sync.WaitGroup{}

	for _, n := range names {
		wg.Add(1)
		go printAge(n, &wg)
	}

	for _, n := range names {
		wg.Add(1)
		go printAge(n, &wg)
	}

	wg.Wait()

}

func printAge(name string, wg *sync.WaitGroup) {
	cached, age := getAge(name)
	fmt.Printf("[CACHED: %t]Got %s's age (%d)\n", cached, name, age.age)
	wg.Done()
}
