package main

import (
	"fmt"
	"sync"
)

const (
	amoutGoroutines = 5
)

func main() {
	var wg sync.WaitGroup
	printer := func(id int) {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			str := fmt.Sprintf("%d: Мой ID: %d", i, id)
			fmt.Println(str)
		}
	}

	wg.Add(amoutGoroutines)
	for i := 0; i < amoutGoroutines; i++ {
		go printer(i)
	}

	wg.Wait()
}
