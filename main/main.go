package main

import (
	"fmt"
	"math"
)


func main()  {
	//fmt.Println("number is ",rand.Intn(10))
	//fmt.Printf("number is ",math.Sqrt(9))

	//在GO中，首字母大写的名称是可以导出的，任何未导出的名字是不能被包外的代码访问的
	fmt.Println(math.Pi)

	fmt.Println(add(2,3))
	fmt.Println(plus(2,4))


}


func add(x int,y int)int  {
	return x + y
}
//连续参数类型若一致，可简写为

func plus(x,y int)int  {
	return x - y
}

func plus0(name string)  {

}

func plus1(a,b int) int  {

}

func plus2(name string)  {

}

