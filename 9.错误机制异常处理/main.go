package main

import (
	"errors"
	"fmt"
)

func test() {
	//使用defer_recover 来捕获和处理异常,函数结束或者异常时运行这里
	defer func() {
		err := recover() //内置函数,可捕获异常
		if err != nil {
			fmt.Println("发生错误:", err)
		}
	}()

	num1 := 10
	num2 := 0
	res := num1 / num2

	fmt.Println("res=", res)

}

//自定义异常
func diy_err(name string) error {

	if name == "confing.ini" {
		return nil
	} else {
		//返回一个自定义错误
		return errors.New("读取文件类型错误..")
	}

}

//自定义错误演示
func test2(name string) {
	err := diy_err(name)
	if err != nil {
		//如果发生错误 就输出错误,并终止程序
		panic(err)

	}
	fmt.Println("test02()继续执行....")

}

func main() {

	test()
	fmt.Println("下面的代码")

	//自定义错误演示
	test2("123")
	fmt.Println("下面的代码")

}
