package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Cache struct {
	m    sync.Mutex
	data map[string]Data
}

type Data struct {
	age int
}

var cache = Cache{
	data: map[string]Data{
		"Kirill": Data{age: 34},
		"Anna":   Data{age: 23},
		"Alex":   Data{age: 19},
		"Kate":   Data{age: 48},
	},
}

func getAge(name string) *Data {
	cache.m.Lock()
	data, ok := cache.data[name]
	cache.m.Unlock()

	if ok {
		fmt.Printf("CACHED")
		return &data
	}

	cache.m.Lock()
	personAge := getFromDb(name)
	entry := Data{age: personAge}
	cache.data[name] = entry
	cache.m.Unlock()

	return &entry

}

func getFromDb(n string) int {
	fmt.Printf("Getting %s age from db\n", n)
	time.Sleep(time.Second * 2)

	return rand.Intn(80-23+1) + 23
}

func main() {
	names := []string{"Kirill", "Anna", "Alice", "Alex", "Bob", "Charlie"}
	var wg sync.WaitGroup

	for _, n := range names {
		wg.Add(1)
		go func(name string, waitGroup *sync.WaitGroup) {
			age := getAge(name)
			fmt.Printf("Got %s's age (%d)\n", name, age)
			waitGroup.Done()

		}(n, &wg)
	}

	for _, n := range names {
		wg.Add(1)
		go func(name string, waitGroup *sync.WaitGroup) {
			age := getAge(name)
			fmt.Printf("#2 Got %s's age (%d)\n", name, age)
			waitGroup.Done()

		}(n, &wg)
	}

	wg.Wait()
}
