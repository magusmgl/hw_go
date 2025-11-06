package main

import (
	"fmt"
	"sync"
	"tasks/cache"
)

// --4.1 ----------------------------------------
type ConcurrentMap struct {
	data map[string]string
	m    sync.RWMutex
}

func newConcurretnMap() *ConcurrentMap {
	return &ConcurrentMap{
		data: make(map[string]string),
		m:    sync.RWMutex{},
	}
}

func (cm *ConcurrentMap) GetOrCreate(key string, value string) string {
	cm.m.RLock()
	v, ok := cm.data[key]
	cm.m.RUnlock()
	if ok {
		return v
	}

	cm.m.Lock()
	defer cm.m.Unlock()
	if v, ok := cm.data[key]; ok {
		return v
	}
	cm.data[key] = value
	return value
}

//-------------------------------------------------------------

func main() {

	// fmt.Println(uniqN(10))

	//-------------------------------------
	//4.1
	// cm := newConcurretnMap()

	// wg := sync.WaitGroup{}
	// wg.Add(2)

	// go func() {
	// 	defer wg.Done()
	// 	val := cm.GetOrCreate("key", "value1")
	// 	fmt.Println("Goroutine 1 got, ", val)
	// }()

	// go func() {
	// 	defer wg.Done()
	// 	val := cm.GetOrCreate("key", "value2")
	// 	fmt.Println("Goroutine 2 got, ", val)
	// }()

	// wg.Wait()
	//--------------------------------------
	//4.2
	// stockStream := stock.UpdateProductStock()

	// var stockHistory []map[string]int

	// for stock := range stockStream {
	// 	// stock := <-stockStream
	// 	stockHistory = append(stockHistory, stock)
	// }
	// for i, stock := range stockHistory {
	// 	fmt.Printf("Iteration %d: %v\n", i+1, stock)
	// }

	//-------------------------------------
	//4.3
	// wc := wordcount.NewWordCounter(3)

	// words := []string{"apple", "banana", "apple", "orange", "grape", "banana", "kiwi"}
	// for _, word := range words {
	// 	wc.CountWord(word)
	// }

	// fmt.Println("Количество слов: ", wc.Counts)
	//_----------------------

	//---4.4
	// oldMap := map[string][]string{
	// 	"group1": {"apple", "banana"},
	// 	"group":  {"carrot"},
	// }

	// newValues := []string{"banana", "cherry"}

	// key := "group"

	// MergeToMap(oldMap, key, newValues)

	// 5.3
	// creditCard := &payment.CreditCardProcessor{
	// 	Limit: 100.0,
	// }
	// payment.ExecutePayment(creditCard, 50.0)
	// ----------------------------------------

	// 5.6
	cache := &cache.Cache{
		Data: make(map[string]any),
	}

	cache.Store("name", "Alice")
	cache.Store("age", 25)

	name, ok := cache.Load("name")
	if !ok {
		fmt.Println("Name not found")
		return
	}
	nameStr, ok := name.(string)
	if !ok {
		fmt.Println("Error cast to string")
		return
	}

	fmt.Println("Name: ", nameStr)

	age, ok := cache.Load("age")
	if !ok {
		fmt.Println("Age not found ")
		return
	}
	ageInt, ok := age.(int)
	if !ok {
		fmt.Println("Error cast to int")
		return
	}
	fmt.Println("Age: ", ageInt)

}

// 4.4
func MergeToMap(oldMap map[string][]string, key string, newValue []string) {
	unique := make(map[string]struct{})
	for _, v := range oldMap[key] {
		unique[v] = struct{}{}
	}

	for _, v := range newValue {
		if _, ok := unique[v]; !ok {
			oldMap[key] = append(oldMap[key], v)
			unique[v] = struct{}{}
		}
	}
	fmt.Println(oldMap[key])
}

//-----------------------------------------

// func uniqN(n int) []int {
// 	m := make(map[int]struct{}, n)
// 	res := make([]int, 0, 10)
// 	for len(res) <= n {
// 		num := rand.Intn(10)
// 		if _, ok := m[num]; !ok {
// 			res = append(res, num)
// 			m[num] = struct{}{}
// 		}
// 	}
// 	return res
// }
