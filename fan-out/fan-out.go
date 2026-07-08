package main

import (
	"fmt"
	"sync"
)

func fanOut(in <-chan int, numWorkers int, process func(int) int) []<-chan int {
	outs := make([]<-chan int, numWorkers)

	for i := 0; i < numWorkers; i++ {
		out := make(chan int)
		outs[i] = out

		go func(out chan<- int) {
			defer close(out)
			for n := range in {
				out <- process(n)
			}
		}(out)
	}

	return outs
}

func main() {
	// готовим входной канал
	in := make(chan int)
	go func() {
		defer close(in)
		for i := 1; i <= 10; i++ {
			in <- i
		}
	}()

	// распараллеливание обработки одного канала между 3 воркерами.
	outs := fanOut(in, 3, func(n int) int {
		return n * n
	})

	// читаем все выходные каналы
	var wg sync.WaitGroup
	for _, out := range outs {
		wg.Add(1)
		go func(ch <-chan int) {
			defer wg.Done()
			for n := range ch {
				fmt.Println(n)
			}
		}(out)
	}
	wg.Wait()
}
