package main

import "fmt"

func main() {

	//切片slice 成员数不确定时用切片 用法和数组一样
	var arr = [...]int{1, 22, 333, 44, 5, 6}
	//引用数组的成员 成员1  到成员3-1  3不包含
	slice := arr[1:3]
	slice = append(slice, 2) //添加成员 ,会把数组的成员也给覆盖了
	slice = append(slice, 2)
	slice = append(slice, 2)
	slice = append(slice, 2)

	//追加切片
	slice = append(slice, slice...)

	fmt.Println("arr=", arr)
	fmt.Println("slice的元素=", slice)
	fmt.Println("slice的元素个数=", len(slice))
	fmt.Println("slice 的容量=", cap(slice)) //容量可以动态变化

	//切片的 make
	//对于切片,必须make才能指定成员数和初始化数据
	slice1 := make([]float64, 5, 10)
	slice1[0] = 50
	slice1[3] = 60
	fmt.Println(slice1)
	fmt.Println("slice1的size=", len(slice1))
	fmt.Println("slice1的cap=", cap(slice1))
	//或者直接赋值
	slice2 := []float64{1, 2, 3, 5, 6537, 87, 87}
	fmt.Println(slice2)
	fmt.Println("slice2的size=", len(slice2))
	fmt.Println("slice2的cap=", cap(slice2))

	//copy
	slice3 := []int{1, 2, 3, 4, 5}
	slice4 := make([]int, 10)
	//将slice4 拷贝上 slice3的数据 就是值拷贝,之后两个切片数据互不影响
	copy(slice4, slice3)
	slice4[0] = 199 //之后两个切片数据互不影响
	fmt.Println("slice3", slice3)
	fmt.Println("slice4", slice4)

}
