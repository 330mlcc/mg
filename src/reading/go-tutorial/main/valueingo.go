package main

import (
	"fmt"
	"time"
)



func main() {
	fmt.Println("------------变量相关的操作方法示例------------------")
	fmt.Println("go" + "lang")


	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)


	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	fmt.Println("------------for循环相关的操作方法示例------------------")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}


	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	fmt.Println("------------if/else相关的操作方法示例------------------")
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	fmt.Println("------------分支结构相关的操作方法示例------------------")
	ints := 2
	fmt.Print("write ", ints, " as ")
	switch ints {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
	fmt.Println("it's the weekend")
	default:
	fmt.Println("it's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
	fmt.Println("it's before noon")
	default:
	fmt.Println("it's after noon")
	}

	fmt.Println("------------数组相关的操作方法示例------------------，")
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		//fmt.Println("twoD[",i,"] = ",twoD[i])
		for j := 0; j < 3; j++ {
			//fmt.Println("twoD[",i,"] = ",twoD[i])
			fmt.Println("twoD[",i,"][",j,"] = ",twoD[i][j])
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	fmt.Println("------------切片相关的操作方法示例------------------")
	s := make([]string,3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s)
	s = append(s, "d")
	s = append(s,"e")
	s = append(s,"f")
	fmt.Println(s)

	c := make([]string,len(s))
	copy(c, s)
	l := c[2:5]
	fmt.Println(l)

	arrays := make([][]int ,3)
	for i:=0;i<3;i++{
		innerlen := i+1
		arrays[i] = make([]int,innerlen)
		for j:=0;j<innerlen;j++{
			arrays[i][j] = i+j
		}
	}
	fmt.Println("2d: ",arrays)

	fmt.Println("------------关联数组相关的操作方法示例------------------")
	fmt.Println("map是go内置的关联数据类型，其它语言种成为哈希或字典")
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 11

	fmt.Println("map : ",m)
}

