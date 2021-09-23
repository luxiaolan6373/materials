package main

import "fmt"

func main() {
	/* 数组是同一类型的一组数据 成员数固定不能变化
	一个养鸡场有6只鸡,它们的体重分别是3kg,5kg,1kg,4kg,2kg,50kg,请问总体重是多少 平均是多少
	*/

	var hens [6]float64

	hens[0] = 3.0
	hens[1] = 5.0
	hens[2] = 1.0
	hens[3] = 3.5
	hens[4] = 4.0
	hens[5] = 50.0
	totaWeight := 0.0

	for _, value := range hens {
		totaWeight += value

	}

	fmt.Printf("总体重为%v 平均值%.2f", totaWeight, totaWeight/float64(len(hens)))

	var intArr [3]int
	//当我们定义完数组后,其实数组的各个元素有默认值 0
	fmt.Println(intArr)
	//第一个成员就是首地址 也就是数组的地址
	//后面的成员就是+8 int占用的字节为8
	fmt.Printf("intArr地址=%p intArr[0]地址%p intArr[1]地址%p\n", &intArr, &intArr[0], &intArr[1])

	//有以下几种赋值方式
	var arr2 = [...]int{1, 2, 3, 4, 5} //不填成员算,通过赋值自动计算成员数
	fmt.Println(arr2)
	fmt.Printf("%F", arr2)

	var arr3 = [5]int{7, 8, 9, 4, 5}
	fmt.Println(arr3)

	arr4 := [...]int{0: 7, 2: 8, 3: 9, 4: 4, 1: 5}
	fmt.Println(arr4)

	arr5 := [...]string{0: "A", 2: "B", 3: "C", 4: "T", 1: "I"}
	fmt.Println(arr5)

	//数组运用案例
	//打印大写字母 26个
	var myChars [26]byte

	for i := 0; i < len(myChars); i++ {

		myChars[i] = 'A' + byte(i)
		fmt.Printf("%c", myChars[i])

	}
	fmt.Println("")

	//找最大值
	arr6 := [...]int{1, 4, 6, 8, 043, 534}
	maxVal := arr6[0]
	maxValIndex := 0
	for i := 1; i < len(arr6); i++ {
		if maxVal < arr6[i] {
			maxVal = arr6[i]
			maxValIndex = i

		}
	}

	fmt.Printf("maxVal=%v maxValIndex=%v", maxVal, maxValIndex)

}
