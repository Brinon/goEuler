package level1

import (
	"fmt"
)

func Problem1() {
	s := 0
	for i := 1; i < 1000; i++ {
		if i%3 == 0 {
			s += i
		} else if i%5 == 0 {
			s += i
		}
	}
	fmt.Println("Problem 1:")
	fmt.Println(s)
}
