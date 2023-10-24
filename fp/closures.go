package fp

import "fmt"

func closure() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func Closure() {
	increment := closure()

	fmt.Println("Closure:", increment(), increment(), increment()) // 1 2 3
}
