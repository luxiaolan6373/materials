package main

var (
	//第三种 全局匿名函数
	fun1 = func() {
		print(123)
	}
)

func main() {
	//第一种
	func() {
		print(123)
	}()
	//第二种
	a := func() {
		print(123)
	}
	a()

}
