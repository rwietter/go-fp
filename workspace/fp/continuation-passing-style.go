package fp

import (
	"fmt"
	"time"
)

func asyncAdd(a, b int, cont func(int)) {
	go func() {
		result := a + b
		time.Sleep(2 * time.Second)
		cont(result)
	}()
}

func ContinuationPassingStyle() {
	asyncAdd(3, 4, func(result int) {
		fmt.Println(" -> Async Add:", result)
	})

	fmt.Print("Running...")
	time.Sleep(3 * time.Second)
}
