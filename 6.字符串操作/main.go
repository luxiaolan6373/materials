package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "abcdefg中" //中文占两个字节 字母 数字 符号占一个字符
	//len()字符串长度
	fmt.Println(len(s))

	//字符串遍历,同事处理有中文的情况
	r := []rune(s)
	for i, value := range r {
		//字节 转 字符串
		fmt.Println(i, string(value))
	}

	//字符串转整数
	n, err := strconv.Atoi("4565656")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转换成功", n)
	}

	//整数转字符串
	s2 := strconv.Itoa(123)
	fmt.Println("整数转字符串", s2)

	//字符串转[]byte 字节集 字节切片
	var bytes = []byte("hello go")
	fmt.Println("bytes=", bytes)

	//[]byte 转字符串
	var s3 = string([]byte{97, 98, 99})
	fmt.Println("s3", s3)

	//整数10进制转 2 8 16进制的字符串

	s4 := strconv.FormatInt(123, 2)
	fmt.Println("123对应的2进制字符串", s4)
	s5 := strconv.FormatInt(123, 8)
	fmt.Println("123对应的8进制字符串", s5)

	//查找字符串是否存在指定字符串
	fmt.Println("是否存在", strings.Contains("abcdefg", "a"))

	//计算指定字符串中有几个指定子串
	fmt.Println("有多少个c ", strings.Count("abccccddder", "c"))

	//不区分大小写的字符串比较 一般的==是区分的,而不区分应该这样写
	fmt.Println("不区分大小写 ", strings.EqualFold("abc", "Abc"))

	//返回子串 在字符串中第一次出现的index位置,如果没有返回-1
	fmt.Println("a第一次出现的位置 ", strings.Index("dsgasdgagajlgj", "a"))

	//返回子串 在字符串中最后一次次出现的index位置,如果没有返回-1
	fmt.Println("a最后一次出现的位置 ", strings.LastIndex("dsgasdgagajlgj", "a"))

	//字符串替换
	fmt.Println("替换 ", strings.Replace("dsgasdgagajlgj", "a", "123", -1))

	//字符串分割 根据指定标志分割成一个数组
	s_sz := strings.Split("dsgasdgagajlgj", "a")
	for index, value := range s_sz {
		fmt.Println("数组", index, value)
	}

	//字符串大小写转换
	fmt.Println("全部改成小写", strings.ToLower("ABCDADG"))
	fmt.Println("全部改成大写", strings.ToUpper("abcdsdgd"))

	//去掉字符串左右两边的空格
	fmt.Println("去掉两边空格", strings.TrimSpace(" adgsd gds "))

	//将字符串左右两边指定字符去掉 只会去掉开头和结尾的.:
	fmt.Println("去掉两边指定字符", strings.Trim("$% absd@ fdfd$%", "%"))

	//将字符串最左边指定字符去掉:
	fmt.Println("去掉两边指定字符", strings.TrimLeft("$% absd@ fdfd$%", "$"))

	//将字符串最右边指定字符去掉:
	fmt.Println("去掉两边指定字符", strings.TrimRight("$% absd@ fdfd$%", "%"))

	//判断字符串是否以指定字符串开头
	fmt.Println("字符串是否以指定字符串开头", strings.HasPrefix("$% absd@ fdfd$%", "$"))

	//判断字符串是否以指定字符串结尾
	fmt.Println("字符串是否以指定字符串开头", strings.HasSuffix("$% absd@ fdfd$%", "%"))

}
