// solution for counting even ended numbers resulted from multiplying four digit
// numbers.
package main

import (
	"fmt"
)

func main() {
	count := 0
	for a := 1000; a <= 9999; a++ {
		for b := a; b <= 9999; b++ {
			n := a * b
			s := fmt.Sprintf("%d", n)
			if s[0] == s[len(s)-1] {
				count++
			}
		}
	}

	fmt.Println(count)
}
