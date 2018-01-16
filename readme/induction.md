
## 目录


- [简介、编译、运行](#简介)

- [变量](#变量)

- [函数](#函数)

- [结构体](#结构体)





### 简介


> go是一种编译型(编译速度快)、具有静态类型(静态类型意味着变量必须指定一个类型)和类c风格语法的语言，并具备垃圾回收机制。

> 它不是面向对象的语言，所以没有对象和继承的概念（使用组合来实现继承）。也没有多态和重载的特性



- 对于简单的例子，我们不需要深入到go的工作目录中

- 在go语言中，程序的入口是main包中的main函数

- go程序运行命令

```go

go run main.go

```

- 编译

go run会先编译然后再运行你的代码，它会在一个临时的目录下编译这段代码，然后执行，最后自动清除生成的临时文件。
通过运行以下命令你可以看见这个零临时文件的位置：

```go
go run --work main.go

```

- 打包，可执行文件

这会生成一个可执行文件main，你可以直接运行它。在linux/osx中，不要忘记在可执行文件前面加上点和反斜杠，所以你需要输入./main运行程序。

```go
go build main.go
```

- go有很多内置的函数，例如println，不需要引用即可使用。
- go在函数名前加了包名作为前缀，例如，fmt.Println。
- go在导入包的时候是比较严格的，如果导入的包没有被使用，那么程序不能被编译。(这就是为啥编译快的原因)


[go的标准库](https://golang.org/pkg/)



### 变量

1. `var NAME TYPE`声明一个变量，并且变量的初始值为它相应类型的零值，
1. 使用`NAME := VALUE`声明一个变量并赋值，
1. 使用`NAME = VALUE`去给已经声明过的变量赋值


- 一般赋值

```
//声明变量，默认整型为0，布尔型赋false，字符串型赋""等
var power int
//赋值
power = 9000

```

- 声明并赋值

```
var power int = 9000
```
- 简易写法
```
power := 9000
```
- 函数赋值
```
func main() {
    power := getPower()
}
func getPower() int {
    return 9001
}
```

> 在一个代码范围内变量只能被声明一次

- 错误的声明

```
func main() {
   power := 9000
   power := 9001
}

```



- 多变量赋值


```go
//go支持多个变量同时赋值（使用=或者:=）
name, power := "Goku", 9000

```
- 如果多个变量既存在新变量，也存在已声明的变量，也可以使用:=进行赋值

```
 power := 1000
 name, power := "Goku", 9000

```




> go程序中不能存在未使用的变量，




### 函数

- 声明函数

```go

//无返回值函数
func log(name string) {
}

- 一个返回值函数
func add(a int, b int) int {
}

- 多返回值函数
func power(name string) (string, bool) {
}
```

- 函数调用

```go

//普通调用
name, result := power("------")

//只接收某个值
_, result := power("------")

```

>`_`是一个空白标识符。（用于接收不想接收的值）



### 结构体


- 定义一个结构体

```
//结构体名称大写，成员名称大写（这样外部可访问）

type People struct {
	Name string
	Age int
}

```

- 声明和初始化

```go

tom := People{Name:"Tom",Age:16}

//简写
carl := People{"Carl",23}

//结尾的逗号,是不能省的。（不然报编译错误）
adam := People{
    Name:"Adam",
    Age:45,
}

//结构体的字段也会有一个0值。
Bob := People{}

```



> 在go语言中，函数的参数传递都是值传递(值引用)，即传递的是一个拷贝。

- 值引用

```
func copy1()  {
	adam := People{
		Name:"Adam",
		Age:16,
	}
	grow1(adam)
	fmt.Println(adam.Age) //输出16
}
// 值引用 或者说 值copy ，对原结构体的值没有影响
func grow1(people People)  {
	people.Age += 1
}

```

- 指针引用

```go
func copy2()  {
	adam := &People{
		Name:"Adam",
		Age:16,
	}
	grow2(adam)
	fmt.Println(adam.Age) //输出17
}
//也是指针的copy值，但是指向的都是同一个资源。
func grow2(people *People)  {
	people.Age += 1
}




```
> &叫取地址符

> People 是一个值类型

> *People 地址类型 指针类型

> 复制一个指针变量的开销比复制一个复杂的结构体小。在一个64的系统上，指针的大小只有64位。如果我们的结构体有很多字段，
创建一个结构体的拷贝会有很大的性能开销。指针的真正意义就是通过指针可以共享值,所以上面的指针引用很能节省开销。


#### 结构体函数

或者说实例方法，（结构体实例调用的方法）

- 类型*People是super方法的接受者

```go
func (people *People) super()  {
	people.Age += 2
}

```

调用

```go

adam := &People{
		Name:"Adam",
		Age:16,
	}
adam.super()
fmt.Println(adam.Age) //输出18

```


#### 构造函数

结构体没有构造函数，不过可以通过函数来返回一个结构体实例

```go

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

```

```go

fmt.Println(newPeople("Tom",26))  //输出 &{Tom 26}

fmt.Println(newPeople1("Tom",26)) //输出 {Tom 26}

```

尽管没有构造函数，go有一个内置的函数new，可以用来分配一个类型需要的内存。

```go

p1 := new(People)
//等同于 p1 := &People{}
p1.Age = 26
p1.Name = "Army"

fmt.Println(p1)  //输出 &{Army 26}

```

值类型和指针类型类型

```go

//p2 是值类型
p2 := People{Name:"xiaoming",Age:23}

//p1 是指针类型
p1 := new(People)

fmt.Println(p1)  //输出 &{Army 26}
fmt.Println(p2) //输出  {xiaoming 23}

```

#### 结构体成员

> 结构体成员可以是任意的类型，比如整型、字符串类型、布尔类型，以及数组、映射、接口和函数类型

比如我们给People扩展一下


```go

type People struct {
	Name string
	Age int
	Father *People
}

```
初始化

```go

peo := &People{
		Name : "Tom",
		Age : 20,
		Father:&People{
			Name:"Smith",
			Age:56,
		},
	}
fmt.Println(peo.Father)  //输出 &{Smith 56 <nil>}
```


#### 嵌入类型

> 结构体类型可以包含匿名或者嵌入字段。也叫做嵌入一个类型。当我们嵌入一个类型到结构体中时，该类型的名字充当了嵌入字段的字段名。

定义一个Person的结构体
```go
type Person struct {
	Name string
	Height int
}

//接受者为指针类型的方法
func (person *Person) GetName()  {
	fmt.Println("my name is",person.Name)
}
//接受者为值类型的方法
func (person Person) GetNameValue() {
	fmt.Println("my name is",person.Name)
}

```
定义一个User的结构体

```go

type User struct {
	phone string
}

//接受者为指针类型的方法
func (user *User) GetPhone()  {
	fmt.Println("user phone is",user.phone)
}
//接受者为值类型的方法
func (user User) GetPhoneValue() {
	fmt.Println("user phone is",user.phone)
}
```
扩展People

```go

type People struct {
	*Person  //嵌入指针类型
	User       //嵌入结构体类型
	Name string
	Age int
	Father *People
}

```

初始画一个值类型和指针类型的People
```go
//指针类型
p1 := &People{
    Person:&Person{"Tom",67},
    User:User{phone:"17789876543"},
    Name:"Army",
    Age:23,
    Father:&People{
        Name:"Smith",
        Age:56,
    },

}
//值类型
p2 := People{
    Person:&Person{"Alen",78},
    User:User{phone:"16698765432"},
    Name:"Adam",
    Age:26,
    Father:&People{
        Name:"Adan",
        Age:59,
    },
}

```
结果分析

```go

//我们可以通过访问访问嵌入类型间接访问嵌入类型的方法和成员
p1.Person.GetName()
p1.Person.GetNameValue()
p1.User.GetPhone()
p1.User.GetPhoneValue()
p2.Person.GetName()
p2.Person.GetNameValue()
p2.User.GetPhone()
p2.User.GetPhoneValue()
fmt.Println(p1.User.phone)
fmt.Println(p1.Person.Name)
fmt.Println(p2.User.phone)
fmt.Println(p2.Person.Name)

//go语言嵌入类型的方法提升
//指针类型和值类型的people可以直接访问嵌入类型的接受者为值类型或者指针类的方法和成员
p1.GetName()
p1.GetNameValue()
p1.GetPhone()
p1.GetPhoneValue()
p2.GetName()
p2.GetNameValue()
p2.GetPhone()
p2.GetPhoneValue()
fmt.Println(p1.phone)
fmt.Println(p2.phone)

//当people的成员字段名和嵌入类型的字段名冲突的时候，people的成员优先级更高
fmt.Println(p2.Name) //输出p2的name，而不是p2的User的name
fmt.Println(p1.Name)


```




以上可看出 假设结构体类型P(*People)和嵌入类型B(*Person)

- P可以直接访问B的

