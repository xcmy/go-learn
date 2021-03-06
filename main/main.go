package main

import (
	"fmt"
	//"os"
	//"go-learn/errors"
	//"go-learn/router"
	"errors"
	//"log"
	"time"
)

type Person struct {
	Name   string
	Height int
}
type User struct {
	phone string
}

//接受者为指针类型的方法
func (person *Person) GetName() {
	fmt.Println("my name is", person.Name)
}

//接受者为值类型的方法
func (person Person) GetNameValue() {
	fmt.Println("my name is", person.Name)
}

//接受者为指针类型的方法
func (user *User) GetPhone() {
	fmt.Println("user phone is", user.phone)
}

//接受者为值类型的方法
func (user User) GetPhoneValue() {
	fmt.Println("user phone is", user.phone)
}

type People struct {
	*Person
	User
	Name   string
	Age    int
	Father *People
}

func (people *People) referral() {
	fmt.Println("people name is", people.Name)
}

type userLog interface {
	Log(message string)
}

//
//type Shape interface {
//	area() float64
//}
//
//type Rectangle struct {
//	width,height float64
//}
//
//func (rectangle Rectangle) area() float64 {
//	return rectangle.width*rectangle.height
//}
//
//type Triangle struct {
//	width,height float64
//}
//
//func (triangle Triangle) area() float64 {
//	return triangle.width*triangle.height/2
//}
//
//func getArea(shape Shape) float64  {
//	return shape.area()
//}

//
//type error interface {
//	Error() string
//}

func success(bool bool) error {
	if bool {
		return errors.New("success")
	}
	return nil
}

func a() int {

	fmt.Println("a start")

	ex1 := func() {
		fmt.Println("fun1")
	}
	ex2 := func() {
		fmt.Println("fun2")
	}
	defer ex1()
	defer ex2()
	defer func() {
		r := recover()
		fmt.Println("error is :", r)
	}()

	panic("panic")

	fmt.Println("a end")

	return 1

}
func yuyu(a interface{}, b interface{}) interface{} {

	return string(a.(int) + b.(int))
}

type Area func(a int, b int) int

func rectangle(area Area) int {
	return area(1, 2)
}

var counter = 0

func incr() {
	counter++
	fmt.Println(counter)
}

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}
func main() {

	done := make(chan bool, 1)
	go worker(done)
	<-done
	fmt.Println("contiune")

	//message1 := make(chan string,2)
	////message2 := make(chan string)
	//
	//go func() {
	//	message1 <- "hello"
	//	message1 <- "jack"
	//}()
	//
	//go func() {
	//	fmt.Println(<-message1)
	//	fmt.Println(<-message1)
	//}()


	//go func() {
	//	msg := <-message1
	//	fmt.Println("2", msg)
	//}()

	//go func() {
	//	msg := <-message2
	//	fmt.Println("3",msg)
	//}()
	//
	//go func() {
	//	msg := <-message2
	//	fmt.Println("4",msg)
	//}()

	time.Sleep(time.Millisecond * 10)



	//i := 1
	//for i <= 3 {
	//	fmt.Println(i)
	//	i = i + 1
	//}

	//for j := 1; j <= 10; j+=2 {
	//	fmt.Println(j)  //输出1、3、5、7、9
	//}

	//for {
	//	fmt.Println("loop")
	//	return
	//}

	//for n := 0; n <= 5; n++ {
	//	if n%2 == 0 {
	//		continue
	//	}
	//	fmt.Println(n)
	//}

	//mytype := func(i interface{}) {
	//	switch i.(type) {
	//	case int,string:
	//		fmt.Println("string or int")
	//	case float64:
	//		fmt.Println("float")
	//	case bool:
	//		fmt.Println("bool")
	//	default:
	//		fmt.Println("other")
	//	}
	//}
	//mytype("me")    //输出





	//i :=



	//for i := 0; i < 2; i++ {
	//	go incr()
	//}
	//time.Sleep(time.Millisecond * 10)

	//go fmt.Println("hello")

	//time.Sleep(time.Millisecond * 10) // this is bad, don't do this!
	//fmt.Println("done")

	//area := rectangle(func(a int, b int) int{
	//	return a * b
	//})
	//
	//fmt.Println(area)

	//fmt.Println(yuyu(2,2))

	//count := 41
	//if x := 10; count > x {
	//	fmt.Println("2222")
	//}

	//if err := process(); err != nil {
	//	fmt.Println("222")
	//}

	//file, err := os.Open("a_file_to_read")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer file.Close()

	// 读这个文件

	//fmt.Println(a())
	//fmt.Println("continue")

	//fmt.Println(success(true))

	//rectangle := Rectangle{width:10,height:10}
	//triangle := Triangle{width:10,height:10}
	//
	//fmt.Println("长方形面积为：",getArea(rectangle))
	//fmt.Println("三角形面积为：",getArea(triangle))
	//
	//fmt.Println(rectangle.area())

	//rou := router.Router{Name:"polly"}
	//fmt.Println(rou)
	//rou.GetName()

	//
	//dic := make(map[string]string,10)
	//dic["name"] = "xiaoming"
	//fmt.Println(dic)
	//
	////第一个参数代表对应的值，第二个代表是否存在
	//value,exists := dic["name"]
	//fmt.Println(value,exists) //  输出  xiaoming true
	//
	//
	//fmt.Println(len(dic))

	//删除键为name的映射
	//delete(dic,"name")

	//for key,value := range dic{
	//	fmt.Println(key,value)
	//}

	//mmp := map[int]string{
	//	2:"a",
	//	3:"b",
	//}
	//fmt.Println(mmp)

	//fmt.Println(len(dic))

	//str := "welcome to beijing"
	//str1 := strings.Index(str[5:], "o")
	//fmt.Println(str)
	//fmt.Println(str[5:])
	//fmt.Println(str1)

	//user := []int{1,2,3,4,5}
	//peo := make([]int,5)
	//copy(peo[2:3],user[0:3])
	//fmt.Println(peo) // [0 0 1 0 0]
	//
	//copy(peo[2:4],user[0:3])
	//fmt.Println(peo) // [0 0 1 2 0]

	//
	//arrays := []string{"a","b"}
	//
	//arrays[1] = "c"
	//fmt.Println(arrays)
	//
	//定义一个长度和容量均为5的切片
	//pos := make([]int,5)
	//pos[1] = 3
	//pos[4] = 8
	//fmt.Println(pos)

	//定义一个长度为0容量为10的切片
	//users := make([]int,0,10)
	//fmt.Println(users)

	//users = users[0:6]
	//users[5] = 9
	//fmt.Println(users)

	//users = append(users,5)
	//fmt.Println(users)
	//
	//fmt.Println(cap(users))
	//uc := cap(users)
	//for i:=0;i<25 ;i++  {
	//	users = append(users,5)
	//	if uc != cap(users) {
	//		fmt.Println("索引:",i,"; user容量:",cap(users))
	//		uc = cap(users)
	//	}
	//}

	//users := []int{1,2,3,4,5}
	//slice := users[2:5]
	//fmt.Println(slice)
	//slice[1] = 10
	//fmt.Println(users)
	//
	//users[2] = 10
	//fmt.Println(users)

	//var array [4]int
	//array[0] = 1
	//fmt.Println(array) //输出[1 0 0 0]
	//
	//array2 := [3]string{"a","b","c"}
	//fmt.Println(array2) //输出[a b c]
	//
	//for index,value := range array2 {
	//	fmt.Println(index,"--",value)
	//}

	//指针类型
	//p1 := &People{
	//	Person:&Person{"Tom",67},
	//	User:User{phone:"17789876543"},
	//	Name:"Army",
	//	Age:23,
	//	Father:&People{
	//		Name:"Smith",
	//		Age:56,
	//	},
	//
	//}
	////值类型
	//p2 := People{
	//	Person:&Person{"Alen",78},
	//	User:User{phone:"16698765432"},
	//	Name:"Adam",
	//	Age:26,
	//	Father:&People{
	//		Name:"Adan",
	//		Age:59,
	//	},
	//}
	//
	//
	//p1.GetName()
	//p1.GetNameValue()
	//p1.GetPhone()
	//p1.GetPhoneValue()
	//p2.GetName()
	//p2.GetNameValue()
	//p2.GetPhone()
	//p2.GetPhoneValue()
	//fmt.Println(p1.phone)
	//fmt.Println(p2.phone)
	//
	//p1.Person.GetName()
	//p1.Person.GetNameValue()
	//p1.User.GetPhone()
	//p1.User.GetPhoneValue()
	//p2.Person.GetName()
	//p2.Person.GetNameValue()
	//p2.User.GetPhone()
	//p2.User.GetPhoneValue()
	//fmt.Println(p1.User.phone)
	//fmt.Println(p1.Person.Name)
	//fmt.Println(p2.User.phone)
	//fmt.Println(p2.Person.Name)
	////
	////
	//fmt.Println(p2.Name)
	//fmt.Println(p1.Name)

	//指针类型可调用嵌入类型 接受者为值类型的方法
	//peo.intro()
	////值类型可调用嵌入类型 接受者为指针类型的方法
	//p1.intro()
	//
	//peo.Person.intro()
	//fmt.Println(peo.Name)
	//fmt.Println(peo.Person.Name)

	//peo.referral()

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

	//peo := &People{
	//	Name : "Tom",
	//	Age : 20,
	//	Father:&People{
	//		Name:"Smith",
	//		Age:56,
	//	},
	//}
	//fmt.Println(peo.Father)

}

func newPeople(name string, age int) *People {
	return &People{
		Name: name,
		Age:  age,
	}
}

func newPeople1(name string, age int) People {
	return People{
		Name: name,
		Age:  age,
	}
}

func copy1() {
	adam := People{
		Name: "Adam",
		Age:  16,
	}
	grow1(adam)
	fmt.Println(adam.Age) //输出16
}
func grow1(people People) {
	people.Age += 1
}

func copy2() {
	adam := &People{
		Name: "Adam",
		Age:  16,
	}
	grow2(adam)
	adam.super()
	fmt.Println(adam.Age) //输出17
}
func grow2(people *People) {
	people.Age += 1
}

func (people *People) super() {
	people.Age += 2
}

func add(x int, y int) int {
	return x + y
}

//连续参数类型若一致，可简写为

func plus(x, y int) int {
	return x - y
}
