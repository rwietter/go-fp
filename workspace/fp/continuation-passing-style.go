package fp

import (
	"fmt"
	"time"
)

func asyncAdd(a, b int, ch chan int) {
	result := a + b
	time.Sleep(2 * time.Second)
	ch <- result
}

func ContinuationPassingStyle() {
	ch := make(chan int)
	defer close(ch)

	go asyncAdd(3, 4, ch)

	fmt.Print("Running...")

	result := <-ch
	fmt.Println(" -> Async Add:", result)
}

// -----------------------------------------------
// ------------- WaitGroup Example ---------------
// -----------------------------------------------
/*
func asyncAdd(a, b int, cont func(int)) {
	go func() {
		result := a + b
		time.Sleep(2 * time.Second)
		cont(result)
	}()
}

func ContinuationPassingStyle() {
	var wg sync.WaitGroup
	wg.Add(1)

	asyncAdd(3, 4, func(result int) {
		fmt.Println(" -> Async Add:", result)
		wg.Done()
	})

	fmt.Print("Running...")
	wg.Wait()
}
*/
