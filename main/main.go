package main

import (
	"fmt"
	"math"
)


type People struct {
	Name string
	Age int
}

func main()  {
	//fmt.Println("number is ",rand.Intn(10))
	//fmt.Printf("number is ",math.Sqrt(9))

	//在GO中，首字母大写的名称是可以导出的，任何未导出的名字是不能被包外的代码访问的
	fmt.Println(math.Pi)

	fmt.Println(add(2,3))
	fmt.Println(plus(2,4))


	//Bob := People{}
	//tom := People{Name:"Tom",Age:16}
	//carl := People{"Carl",23}
	//adam := People{
	//	Name:"Adam",
	//	Age:45,
	//}
	//adam.Age= 34
	//tom.Age = 23
	//carl.Age = 22
	//Bob.Age = 56

	//fmt.Println(adam.Age)


	copy1()
	copy2()

}

func copy1()  {
	adam := People{
		Name:"Adam",
		Age:16,
	}
	grow1(adam)
	fmt.Println(adam.Age) //输出16
}
func grow1(people People)  {
	people.Age += 1
}



func copy2()  {
	adam := &People{
		Name:"Adam",
		Age:16,
	}
	grow2(adam)
	fmt.Println(adam.Age) //输出17
}
func grow2(people *People)  {
	people.Age += 1
}


func add(x int,y int)int  {
	return x + y
}
//连续参数类型若一致，可简写为

func plus(x,y int)int  {
	return x - y
}

