package level1

import (
	"fmt"

	"github.com/brinon/goEuler/common"
)

func Problem4() {
	max := -1
	for i := 999; i > 99; i-- {
		for j := 999; j > 99; j-- {
			if i*j > max && common.IsPalindrome(i*j) {
				max = i * j
			}
		}
	}
	fmt.Println("Problem 4:", max)
}
