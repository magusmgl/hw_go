package postman

import (
	"context"
	"sync"
	"time"
)

func Postman(
	ctx context.Context,
	wg *sync.WaitGroup,
	transferPoint chan<- string,
	mail string,
	number int) {

	defer wg.Done()

	for {
		println("Я почтальон номер ", number, ". Взял письмо!")

		select {
		case <-ctx.Done():
			println("Я почтальон номер ", number, "Завершил рабочий день!")
			return
		case <-time.After(1 * time.Second):
			println("Я почтальон номер ", number, ". Донес письмо до почты: ", mail)
		}

		select {
		case <-ctx.Done():
			println("Я почтальон номер ", number, "Завершил рабочий день!")
			return
		case transferPoint <- mail:
			println("Я почтальон номер ", number, ". Передал письмо: ", mail)
		}
	}
	// for {
	// 	select {
	// 	case <-ctx.Done():
	// 		println("Я почтальон номер ", number, "Завершил рабочий день!")
	// 		return
	// 	default:
	// 		println("Я почтальон номер ", number, ". Взял письмо!")
	// 		time.Sleep(1 * time.Second)
	// 		println("Я почтальон номер ", number, ". Донес письмо до почты: ", mail)
	// 		transferPoint <- mail
	// 		println("Я почтальон номер ", number, ". Передал письмо: ", mail)
	// 	}
	// }
}

func PostmanPool(ctx context.Context, postmanCount int) <-chan string {
	mailTransferPoint := make(chan string)
	wg := &sync.WaitGroup{}

	for i := 1; i <= postmanCount; i++ {
		wg.Add(1)
		go Postman(ctx, wg, mailTransferPoint, postmanToMail(i), i)
	}

	go func() {
		wg.Wait()
		close(mailTransferPoint)
	}()

	return mailTransferPoint
}

func postmanToMail(postmanNumber int) string {
	ptm := map[int]string{
		1: "dasdasd",
		2: "zczxcx",
		3: "erwerew",
	}

	mail, ok := ptm[postmanNumber]
	if !ok {
		return "AAAAA"
	}
	return mail
}
