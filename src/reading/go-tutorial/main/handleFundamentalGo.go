package main

import (
	"fmt"
	"time"
	"math"
)

func myplus(aa int,bb int)int{
	return aa+bb
}

func returnManyVlue()(int,int){
	return 3,9
}

// 使用任意数量的参数调用
func sums(nums ...int) int{
	fmt.Println(nums,"")
	totals := 0
	for _, num :=range nums{
		totals += num
	}
	//fmt.Println(totals)
	return totals
}

//匿名函数的使用
/*
Go支持通过闭包来使用匿名函数。匿名函数在你想定义一个不需要命名的内联函数时是很实用的。
这个intSeq函数返回另一个在intSeq函数体内定义的匿名函数。这个返回的函数使用闭包的方式 隐藏变量i。
 */
func intSeq() func() int{
	i := 0
	return func() int{
		i += 1
		return i
	}
}

// 函数的递归调用
func fact(n int) int  {
	if n == 0{
		return 1
	}
	return n * fact(n-1)
}

// 指针的使用
func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

/*
Go 支持在结构体类型中定义方法
 */
//area 方法有一个接收器类型 rect
type rects struct {
	width,height int
}
//指针类型接收器的例子
func (r *rects) area() int {
	return r.width * r.height
}
//值类型接收器的例子
func (r rects) perim() int{
	return 2*r.width + 2*r.height
}

/*
接口 是方法特征的命名集合
 */
 //几何体的基本接口
type geometry interface {
	area2() float64
	perim2() float64
}
//rect 和 circle 实现这个接口
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}
//Go中实现一个接口，只需要实现接口中的所有方法。rect实现了geometry接口。
func (r rect) area2() float64 {
	return r.width * r.height
}
func (r rect) perim2() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area2() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim2() float64 {
	return 2 * math.Pi * c.radius
}
// 如果一个变量的是接口类型，可以调用这个被命名的接口中的方法。measure函数利用这个特性，可以用在任何geometry上。
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area2())
	fmt.Println(g.perim2())
}

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
	n := map[string]int{"foo":1,"bar":2}
	m["k1"] = 7
	m["k2"] = 11

	fmt.Println("map[m] : ",m)
	fmt.Println("len(map) : ",len(m))
	fmt.Println("map[n] : ",n)

	fmt.Println("------------range相关的操作方法示例------------------")
	arrays2 := []int{1,2,3,4,5}
	sum := 0
	for _,num := range arrays2{
		sum += num
	}
	fmt.Println("sum : ",sum)

	for i,num := range arrays2{
		if num == 3{
			fmt.Println("index : ",i)
		}
	}

	kvs := map[string]string{"key1":"apple","key2":"banana"}
	for k,v := range kvs{
		fmt.Printf("%s -> %s \n",k,v)
	}

	fmt.Println("------------函数相关的操作方法示例------------------")
	res := myplus(3,5)
	fmt.Println("3 + 5 = ",res)
	a3,b3 := returnManyVlue()
	fmt.Println(a3)
	fmt.Println(b3)
	c3, _ := returnManyVlue()
	fmt.Println(c3)

	sums(1,2,3,4,5,6,7,8,9,10)

	fmt.Println("------------函数闭包相关的操作方法示例------------------")
	/*
	我们调用intSeq函数，将返回值也是一个函数）赋给nextInt。这个函数的值包含了自己的值i，这样在每次调用nextInt 时都会更新i的值。
	 */

	 nextInt := intSeq()
	 // 通过多次调用 nextInt 来看看闭包的效果。
	 fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	 //为了确认这个状态对于这个特定的函数是唯一的，我们重新创建并测试一下。
	 newInts := intSeq()
	 fmt.Println(newInts())

	fmt.Println("------------函数递归相关的操作方法示例------------------")
	 fmt.Println(fact(3))


	fmt.Println("------------指针相关的操作方法示例------------------")
	 i1 := 1
	 print("Initial : ",i1)

	 zeroval(i1)
	 fmt.Println("zeroval : ",i1)

	 zeroptr(&i1)
	 fmt.Println("zeroptr : ",i1)

	 fmt.Println("pointer : ",&i1)

	fmt.Println("------------结构体相关的操作方法示例------------------")
	 type persons struct {
	 	name string
	 	age int
	 }

	 fmt.Println(persons{"Bob",13})

	fmt.Println("------------goZ中，结构体中方法相关的操作方法示例------------------")
	 //调用上面为结构体定义的两个方法
	r := rects{width:10, height:5}
	fmt.Println("ares : ",r.area())
	fmt.Println("perim : ",r.perim())
	//Go 自动处理方法调用时的值和指针之间的转化。使用指针来调用方法来避免在方法调用时产生一个拷贝，或者让方法能够改变接受的数据。
	rp := &r
	fmt.Println("areas : ",rp.area())
	fmt.Println("perim",rp.perim())

	fmt.Println("------------go中，接口相关的操作方法示例------------------")
	//结构体类型circle和rect都实现了geometry接口，可以使用它们的实例作为measure的参数。
	int_r := rect{width: 3, height: 4}
	int_c := circle{radius: 5}
	measure(int_r)
	measure(int_c)
}

