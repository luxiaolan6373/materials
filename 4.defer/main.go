package main

import "fmt"

func sum(n1 int, n2 int) int {
	//执行defer时,暂时不执行,会将defer后面的语句压入独立的defer栈中 先入后出
	//压入的时候会把参数的值一起存入栈中
	//比如你打开了一个文件,,你第二句直接就调用defer 关闭文件()..它会自动在函数调用完之后才会执行关闭
	//这种机制就不用担心忘记了关闭文件了,你可以打开后直接调用defer 关闭文件
	defer fmt.Println("ok1 n1=", n1) //3
	defer fmt.Println("ok2 n2=", n2) //2

	res := n1 + n2
	fmt.Println("ok3 res=", res) //1
	return res
}
func main() {

	res := sum(10, 20)
	fmt.Println("res=", res) //4
}
