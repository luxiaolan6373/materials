package main

import "fmt"

func main() {

	//len 变量的长度,跳过

	//new 用来分配内存,主要用来分配值类型 比如int float32 等 返回指针
	num1 := 100
	fmt.Printf("num1的类型%T num1的值=%v num1的地址%v", num1, num1, &num1)
	fmt.Println("")
	num2 := new(int) //类型为指针类型
	*num2 = 100      //给指向值赋值
	fmt.Printf("num2的类型%T num2的值=%v num2的地址%v num2指针指向的值%v", num2, num2, &num2, *num2)

	//make 用来分配内存,主要分配引用类型 比如channel  map slice 后面讲

}
