package main

import "fmt"

func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func main() {
	for n := range generator(1, 2, 3, 4, 5) {
		fmt.Println(n)
	}
}
