package main

import (
	"fmt"
	"sync"
)

func main() {
	//ch1.Deadlock()
	//ch1.LiveLock()
	//ch1.GoodAndBadWorker()
	//ch1.RaceCondition()
	//ch1.RaceConditionSolution()

	//ch3.MemConsumed()

	var wg sync.WaitGroup
	for _, text := range []string{"a", "b", "c"} {
		wg.Add(1)
		go func(text string) {
			defer wg.Done()
			fmt.Println(text)
		}(text)
	}
	wg.Wait()
}
