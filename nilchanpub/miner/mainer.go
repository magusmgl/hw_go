package miner

import (
	"context"
	"sync"
	"time"
)

func Miner(
	ctx context.Context,
	wg *sync.WaitGroup,
	transferPoint chan<- int,
	number int,
	power int) {

	defer wg.Done()

	for {
		println("Я шахтер номер ", number, ". Начал добывать уголь!")

		select {
		case <-ctx.Done():
			println("Я шахтер номер ", number, ". Завершил рабочий день!")
			return
		case <-time.After(1 * time.Second):
			println("Я шахтер номер ", number, ". Добыл уголь: ", power)
		}

		select {
		case <-ctx.Done():
			println("Я шахтер номер ", number, ". Завершил рабочий день!")
			return
		case transferPoint <- power:
			println("Я шахтер номер ", number, ".Передал уголь: ", power)
		}
	}
	// for {
	// 	select {
	// 	case <-ctx.Done():
	// 		println("Я шахтер номер ", number, ". Завершил рабочий день!")
	// 		return
	// 	default:
	// 		println("Я шахтер номер ", number, ". Начал добывать уголь!")
	// 		time.Sleep(1 * time.Second)
	// 		println("Я шахтер номер ", number, ". Добыл уголь!")
	// 		transferPoint <- power
	// 		println("Я шахтер номер ", number, ".Передал уголь!")
	// 	}
	// }
}

func MinerPool(ctx context.Context, mainerCount int) <-chan int {
	coolTransferPoint := make(chan int)
	wg := &sync.WaitGroup{}

	for i := 1; i <= mainerCount; i++ {
		wg.Add(1)
		go Miner(ctx, wg, coolTransferPoint, i, i*10)
	}

	go func() {
		wg.Wait()
		close(coolTransferPoint)
	}()
	return coolTransferPoint
}
