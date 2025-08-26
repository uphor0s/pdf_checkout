package main

import (
	"os"
	"strconv"
	"sync"
)

const counterFile = "counter"

var (
	mu sync.Mutex
)

func loadCounter() int {
	data, err := os.ReadFile(counterFile)
	if err != nil {
		return 0 // file missing or unreadable → start fresh
	}
	n, err := strconv.Atoi(string(data))
	if err != nil {
		return 0
	}
	return n
}

func saveCounter(n int) {
	_ = os.WriteFile(counterFile, []byte(strconv.Itoa(n)), 0644)
}

func increment() {
	mu.Lock()
	defer mu.Unlock()
	cntr := loadCounter()
	cntr++
	saveCounter(cntr) // you can also batch writes for performance

}
