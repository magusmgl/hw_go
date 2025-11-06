package main

import (
	"fmt"
	"math/rand"
)

func main() {
	reader(pow(writer()))
}

func writer() <-chan int {
	ch := make(chan int)
	sl := make([]int, 10)

	for i := 0; i < len(sl); i++ {
		sl[i] = rand.Intn(100)
	}

	go func() {
		defer close(ch)

		for _, v := range sl {
			ch <- v
		}
	}()

	return ch
}

func pow(ch <-chan int) <-chan int {
	powCh := make(chan int)

	go func() {
		defer close(powCh)

		for v := range ch {
			powCh <- v * v
		}
	}()

	return powCh
}

func reader(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}
