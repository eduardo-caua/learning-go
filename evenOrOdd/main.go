package main

import "fmt"

func main() {
	values := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, v := range values {
		if v%2 == 0 {
			fmt.Println(v, "is even")
		} else {
			fmt.Println(v, "is odd")
		}
	}
}
