package main

import "fmt"
import "strings"

//闭包就是一个函数和与其相关环境组合的一个整体(实体)
//累加器
func AddUpper() func(int) int {
	var n int = 10
	return func(x int) int {
		n = n + x
		return n
	}
}

//闭包练习 检查字符串是否有后缀,没有的话就加上后缀
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		//如果name 没有指定后缀名,否则按原来的名字返回
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		} else {
			return name
		}
	}
}
func main() {
	//闭包类似于类

	f := AddUpper()   //这里初始化一次 这里已经执行了 var n int=10
	fmt.Println(f(1)) //这个仅仅是调用return 所以不会运算var n int=10这句
	fmt.Println(f(2)) //这个仅仅是调用return 所以不会运算var n int=10这句
	//如果这样做就每次都会初始化
	fmt.Println(AddUpper()(1))
	fmt.Println(AddUpper()(1))

	//闭包练习调用
	ff := makeSuffix(".jpg")
	println(ff("dfsdfsdfds.jpg"))
	println(ff("12345"))
	println(ff("hhkh"))

}
