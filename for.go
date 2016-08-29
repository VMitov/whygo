package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// while
	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Println(sum)
	}

	// forever
	// for {
	// }
}
