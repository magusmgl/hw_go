package main

import (
	"context"
	"fmt"
	"gorutine/miner"
	"gorutine/postman"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	startTime := time.Now()
	var coal atomic.Int64
	var mails []string
	wg := &sync.WaitGroup{}
	mtx := &sync.Mutex{}

	minerContext, minerCancel := context.WithCancel(context.Background())
	postmanContext, postmanCancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(time.Second * 5)
		postmanCancel()
	}()

	go func() {
		time.Sleep(time.Second * 3)
		minerCancel()
	}()

	coalTransferPoint := miner.MinerPool(minerContext, 1000)
	mailTransferPoint := postman.PostmanPool(postmanContext, 1000)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range coalTransferPoint {
			coal.Add(int64(v))
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range mailTransferPoint {
			mtx.Lock()
			mails = append(mails, v)
			mtx.Unlock()
		}
	}()
	// isCoalClosed := false
	// isMailClosed := false
	// for !isMailClosed || !isCoalClosed {
	// 	select {
	// 	case c, ok := <-coalTransferPoint:
	// 		if !ok {
	// 			isCoalClosed = true
	// 			continue
	// 		}
	// 		coal += c
	// 	case m, ok := <-mailTransferPoint:
	// 		if !ok {
	// 			isMailClosed = true
	// 			continue
	// 		}
	// 		mails = append(mails, m)
	// 		fmt.Println("total mails: ", mails)
	// 	}
	// }
	wg.Wait()

	fmt.Println("Duration: ", time.Since(startTime))
	fmt.Println("Добыто угля ", coal.Load())
	mtx.Lock()
	fmt.Println("Получено писем ", len(mails))
	mtx.Unlock()
}
