package main

import (
	"fmt"
	"time"
)

func main() {
	//1.获取当前时间
	now := time.Now()
	fmt.Println(now)

	//获取其它相关信息
	fmt.Println(now.Year())       //取年
	fmt.Println(int(now.Month())) //取月
	fmt.Println(now.Day())        //取日
	fmt.Println(now.Hour())       //取小时
	fmt.Println(now.Minute())     //取分钟
	fmt.Println(now.Second())     //取秒

	//格式化日期和时间
	//第一种
	fmt.Printf("第一种格式化日期: %d-%d-%d %d:%d:%d", now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())
	dataStr := fmt.Sprintf("%d-%d-%d %d:%d:%d", now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println("\n格式化日期到字符串", dataStr)
	//第二种 必须这样填 也可以用yyyy mm dd  hh mm ss 来输入,会自动变成日期
	fmt.Println("第二种格式化:", now.Format("2006/06/02 15:01:05"))
	fmt.Println("第二种格式化:", now.Format("2006-06-02"))
	fmt.Println("第二种格式化:", now.Format("15时01分05秒"))

	//休眠 延时  常量有这些
	/*
		const (
		    Nanosecond  Duration = 1
		    Microsecond          = 1000 * Nanosecond
		    Millisecond          = 1000 * Microsecond
		    Second               = 1000 * Millisecond
		    Minute               = 60 * Second
		    Hour                 = 60 * Minute
		)
	*/
	time.Sleep(2 * time.Second)

	//获取时间戳 时间截
	fmt.Println("获取时间戳", time.Now().Unix())     //时间戳到秒
	fmt.Println("获取时间戳", time.Now().UnixNano()) //时间戳到纳秒

}
