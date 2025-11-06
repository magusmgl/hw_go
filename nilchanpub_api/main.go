package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type PaymentData struct {
	Description string    `json:"description"`
	Amount      int       `json:"amount"`
	Address     string    `json:"address"`
	FullName    string    `json:"fullName"`
	Time        time.Time `json:"date"`
}

func (p PaymentData) Println() {
	fmt.Println("Description: ", p.Description)
	fmt.Println("Amount: ", p.Amount)
	fmt.Println("Address: ", p.Address)
	fmt.Println("Full name: ", p.FullName)
}

func (p PaymentData) ValidateRequest() bool {
	if p.Description == "" ||
		p.Address == "" ||
		p.FullName == "" {
		return false
	}
	if p.Amount == 0 {
		return false
	}
	return true
}

type TransferData struct {
	TransferSum int    `json:"transferSum"`
	FullName    string `json:"fullName"`
	Time        time.Time
}

func (t TransferData) ValidateRequest() bool {
	if t.TransferSum == 0 {
		return false
	}
	if t.FullName == "" {
		return false
	}
	return true
}

func (t TransferData) Println() {
	fmt.Println(t.TransferSum)
	fmt.Println(t.FullName)
	fmt.Println(t.Time)
}

type PaymentDataResponse struct {
	Money          int           `json:"money"`
	PaymentHistory []PaymentData `json:"paymentData"`
}

type TrasferDataResponce struct {
	CurrentBank     int            `json:"currentBank"`
	CurrentMoney    int            `json:"currentMoney"`
	HistoryTransfer []TransferData `json:"historyTransfer"`
}

var money int = 1000
var bank int = 0
var mtx = sync.Mutex{}
var paymentHistory = make([]PaymentData, 0)
var transferHistory = make([]TransferData, 0)

func baseHandler(w http.ResponseWriter, r *http.Request) {
	fooParam := r.URL.Query().Get("foo")
	booParam := r.URL.Query().Get("boo")

	fmt.Println("fooParam: ", fooParam)
	fmt.Println("booParam: ", booParam)
}

func payHandler(w http.ResponseWriter, r *http.Request) {

	// for k, v := range r.Header {
	// 	fmt.Println("k: ", k, "----V: ", v)
	// }

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		msg := "method not allowed"
		if _, err := w.Write([]byte(msg)); err != nil {
			fmt.Println("fail to write responce", err.Error())
		}
		return
	}

	fmt.Println("HTTP method ", r.Method)

	// httpRequestBody, err := io.ReadAll(r.Body)
	// if err != nil {
	// 	fmt.Println("err", err.Error())
	// 	w.WriteHeader(http.StatusInternalServerError)
	// }

	// if err := json.Unmarshal(httpRequestBody, &payment); err != nil {
	// 	fmt.Println("err: ", err)
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }

	var payment PaymentData
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		msg := "fail to read HTTP body: " + err.Error()
		w.Write([]byte(msg))
		fmt.Println(msg)
		return
	}
	if !payment.ValidateRequest() {
		w.WriteHeader(http.StatusBadRequest)
		msg := "wrong requet body"
		w.Write([]byte(msg))
		return
	}
	payment.Time = time.Now()
	payment.Println()

	mtx.Lock()
	if money-payment.Amount < 0 {
		w.WriteHeader(http.StatusBadRequest)
		msg := "Оплата не возможна, недостаточно средств"
		if _, err := w.Write([]byte(msg)); err != nil {
			fmt.Println("fail to wrtite responce: ", err.Error())
		}
		return
	}

	money -= payment.Amount
	paymentHistory = append(paymentHistory, payment)
	mtx.Unlock()

	var request = PaymentDataResponse{
		Money:          money,
		PaymentHistory: paymentHistory,
	}
	b, err := json.Marshal(request)
	if err != nil {
		fmt.Println("err", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(b); err != nil {
		fmt.Println("err: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	// fmt.Println("Total money: ", money)
	// fmt.Println("History payment: ", paymentHistory)

	// httpRequestBody, err := io.ReadAll(r.Body)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	msg := "fail to read HTTP body:" + err.Error()
	// 	fmt.Println(msg)
	// 	if _, err = w.Write([]byte(msg)); err != nil {
	// 		w.WriteHeader(http.StatusInternalServerError)
	// 		fmt.Println("fail to write responce: ", err.Error())
	// 		return
	// 	}
	// 	return
	// }

	// httpRequestBodyString := string(httpRequestBody)

	// paymentAmount, err := strconv.Atoi(httpRequestBodyString)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	msg := "fail to convert amount: " + err.Error()
	// 	fmt.Println(msg)
	// 	if _, err = w.Write([]byte(msg)); err != nil {
	// 		w.WriteHeader(http.StatusInternalServerError)
	// 		fmt.Println("fail to write responce: ", err.Error())
	// 		return
	// 	}
	// 	return
	// }

	// mtx.Lock()
	// if money-paymentAmount >= 0 {
	// 	money -= paymentAmount
	// 	msg := "Оплата прошла успешно: " + strconv.Itoa(money)
	// 	_, err = w.Write([]byte(msg))
	// 	if err != nil {
	// 		fmt.Println("fail to write responce")
	// 	}
	// }
	// mtx.Unlock()

}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		msg := "method not allowed"
		if _, err := w.Write([]byte(msg)); err != nil {
			fmt.Println("fail to write responce", err.Error())
		}
		return
	}
	fmt.Println("HTTP method ", r.Method)
	var responceData TransferData

	if err := json.NewDecoder(r.Body).Decode(&responceData); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		msg := "fail to read request Body: " + err.Error()
		if _, err = w.Write([]byte(msg)); err != nil {
			fmt.Println("fail to write answer: ", err.Error())
		}
		return
	}

	if !responceData.ValidateRequest() {
		w.WriteHeader(http.StatusBadRequest)
		msg := "bad reguest"
		if _, err := w.Write([]byte(msg)); err != nil {
			msg := "failed to write respoce: " + err.Error()
			println(msg)
		}
		return
	}

	responceData.Println()

	mtx.Lock()
	if money-responceData.TransferSum < 0 {
		w.WriteHeader(http.StatusBadRequest)
		msg := "сумма перевода больше доступных средств"
		w.Write([]byte(msg))
		return
	}
	money -= responceData.TransferSum
	bank += responceData.TransferSum
	responceData.Time = time.Now()
	transferHistory = append(transferHistory, responceData)

	mtx.Unlock()

	responce := TrasferDataResponce{
		CurrentBank:     bank,
		CurrentMoney:    money,
		HistoryTransfer: transferHistory,
	}

	b, err := json.Marshal(responce)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		msg := "failed to write responce"
		w.Write([]byte(msg))
		return
	}

	if _, err := w.Write(b); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		msg := "failed to write responce"

		w.Write([]byte(msg))
		return
	}

	// httpRequestBody, err := io.ReadAll(r.Body)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// }

	// httpRequestBodyString := string(httpRequestBody)

	// saveAmount, err := strconv.Atoi(httpRequestBodyString)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	msg := "fail to convert saveAmount: " + err.Error()
	// 	fmt.Println(msg)
	// 	if _, err = w.Write([]byte(msg)); err != nil {
	// 		w.WriteHeader(http.StatusInternalServerError)
	// 		fmt.Println("fail to write responce: ", err.Error())
	// 		return
	// 	}
	// 	return
	// }

	// mtx.Lock()
	// if money-saveAmount >= 0 {
	// 	money -= saveAmount
	// 	bank += saveAmount
	// 	msg := "Деньги успешно переведены в копилку " + strconv.Itoa(saveAmount)
	// 	_, err = w.Write([]byte(msg))
	// 	if err != nil {
	// 		fmt.Println("fail to write responce")
	// 	}
	// }
	// mtx.Unlock()
}

func main() {
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/save", saveHandler)
	http.HandleFunc("/", baseHandler)

	err := http.ListenAndServe(":5061", nil)
	if err != nil {
		fmt.Println("HTTP server error: ", err)
	}
}
