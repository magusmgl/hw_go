package main

import "fmt"

func main() {
	a := 5
	fmt.Println(&a)
	double(&a)
	fmt.Println(a)

	b := [4]int{1, 2, 3, 4}
	reverse(&b)
	fmt.Println(b)

}

func double(num *int) {
	*num = *num * 2
}

func reverse(arr *[4]int) {
	for index, value := range *arr {
		(*arr)[len(arr)-1-index] = value
	}
}
