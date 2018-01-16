package main

import (
	"fmt"
	//"math"
)


type People struct {
	Name string
	Age int
	Father *People
}

func main()  {
	//fmt.Println("number is ",rand.Intn(10))
	//fmt.Printf("number is ",math.Sqrt(9))

	//在GO中，首字母大写的名称是可以导出的，任何未导出的名字是不能被包外的代码访问的
	//fmt.Println(math.Pi)
	//
	//fmt.Println(add(2,3))
	//fmt.Println(plus(2,4))


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


	//copy1()
	//copy2()

	//fmt.Println(newPeople("Tom",26))
	//fmt.Println(newPeople1("Tom",26))



	//p2 := People{Name:"xiaoming",Age:23}
	//
	//p1 := new(People)
	////等同于 p1 := &People{}
	//p1.Age = 26
	//p1.Name = "Army"
	//
	//fmt.Println(p1)
	//fmt.Println(p2)



	peo := &People{
		Name : "Tom",
		Age : 20,
		Father:&People{
			Name:"Smith",
			Age:56,
		},
	}
	fmt.Println(peo.Father)


}

func newPeople(name string,age int) *People  {
	return &People{
		Name:name,
		Age:age,
	}
}

func newPeople1(name string,age int) People  {
	return People{
		Name:name,
		Age:age,
	}
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
	adam.super()
	fmt.Println(adam.Age) //输出17
}
func grow2(people *People)  {
	people.Age += 1
}

func (people *People) super()  {
	people.Age += 2
}

func add(x int,y int)int  {
	return x + y
}
//连续参数类型若一致，可简写为

func plus(x,y int)int  {
	return x - y
}

