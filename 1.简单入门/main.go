package main

import (
	"fmt"
	"time"
)

func jiujiu() {

	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %d  ", i, j, i*j)
		}
		fmt.Println("\n")

	}

}

func add(num1 float64, num2 float64) float64 {

	return num1 + num2

}
func taowa(n int) {
	if n > 2 {
		n--
		taowa(n)
	} else {
		fmt.Println("n=", n)
	}

}

func for_item() {
	var str = "fasfa"
	for i, value := range str {

		fmt.Println("for_item", i, string(value))

	}

}
func main() {
	old_time := time.Now().Unix()
	var st string

	fmt.Scanln(&st)
	print(st)
	jiujiu()
	println("加法", int(add(1, 2)))
	taowa(10)

	println("用时:", time.Now().Unix()-old_time)

	for_item()

}
