
## 目录


- [简介、编译、运行](#简介)

- [变量](#变量)

- [函数](#函数)

- [结构体](#结构体)

- [嵌入类型](#嵌入类型)




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


#### 结构体方法

go里面的method和func的区别就是，方法有一个接收者

> A method is a function with a receiver

- 类型*People是super方法的接收者

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


### 嵌入类型


go语言没有继承的概念，但是提倡使用代码复用的方式，即组合。嵌入类型就是组合的一种方式。

>嵌入类型方便我们扩展或者修改已有的类型

> 结构体类型可以包含匿名或者嵌入字段。也叫做嵌入一个类型。当我们嵌入一个类型到结构体中时，该类型的名字充当了嵌入字段的字段名。

定义一个Person的结构体
```go
type Person struct {
	Name string
	Height int
}

//接收者为指针类型的方法
func (person *Person) GetName()  {
	fmt.Println("my name is",person.Name)
}
//接收者为值类型的方法
func (person Person) GetNameValue() {
	fmt.Println("my name is",person.Name)
}

```
定义一个User的结构体

```go

type User struct {
	phone string
}

//接收者为指针类型的方法
func (user *User) GetPhone()  {
	fmt.Println("user phone is",user.phone)
}
//接收者为值类型的方法
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

初始化一个值类型和指针类型的People
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
//指针类型和值类型的people可以直接访问嵌入类型的接收者为值类型或者指针类的方法和成员
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

//当people的成员字段名和嵌入类型的字段名冲突的时候，嵌入类型的字段名被覆盖。
fmt.Println(p2.Name) //输出p2的name，而不是p2的User的name
fmt.Println(p1.Name)


```


以上可看出 假设结构体类型P(*People 或者说外部类型)和嵌入类型B(*Person)

> P和&P可以直接访问B的接收者为值类型或者指针类型的方法和成员（类似继承）

> P和&P也可以通过访问嵌入类型来间接访问嵌入类型的方法和成员

> 外部类型可以设置和嵌入类型一样的成员字段名和方法名来覆盖嵌入类型 （类似重载）



### 数组


在go中，数组是固定大小的。则不可改变。不能访问超过数组边界的元素。

```go
//声明再赋值
var array [4]int
array[0] = 1
fmt.Println(array) //输出[1 0 0 0]
//直接初始化
array2 := [3]string{"a","b","c"}
fmt.Println(array2) //输出[a b c]
//遍历
for index,value := range array2 {
    fmt.Println(index,"--",value)
}

```


### 切片

切片是Go数组的一个抽象，或者说是Go数组的一个封装。切片比数组更加灵活好用。

#### 切片声明和创建

> 切片同数组一样不能访问超出切片范围的元素

1. 第一种，简单切片使用这种方式

```go
arrays := []string{"a","b"}

arrays[1] = "c"
fmt.Println(arrays)  //  [a c]
```
2. 第二种,在切片特定位置写入值很有用

```go
//定义一个长度和容量均为5的切片
pos := make([]int,5)
pos[1] = 3
pos[4] = 8
fmt.Println(pos)

```

3. 第三种，定义一个空切片，对于切片数量未知的时候使用
```go
var names []string
fmt.Println(names)

```

4. 第四种，可以指定切片的长度和容量。

```go
//定义一个长度为0容量为10的切片
users := make([]int,0,10)
fmt.Println(users)

```

####  添加元素、切片扩容

- `cap()`是一个内置函数，能计算切片的容量

```go
fmt.Println(cap(users)) //输出10
```

> 使用`append()`给切片添加元素，如果切片容量不够会自动扩容。go扩展数组使用的是2倍算法（2x algorithm）


```go
fmt.Println(users)  //输出 []
fmt.Println(cap(users)) //输出10

uc := cap(users)
for i:=0;i<25 ;i++  {
    users = append(users,5)
    if uc != cap(users) {
        fmt.Println("索引:",i,"; users容量:",cap(users))
        uc = cap(users)
    }
}

```


输出结果

```shell
[]
10
索引: 10 ; users容量: 20
索引: 20 ; users容量: 40
```
可以看出，users初始容量为10；
当使用`append()`给users添加第十一个元素的时候，users的容量变为20；
给users添加第21个元素的时候，users的容量变为40；


#### 子切片

在其他语言中如ruby、js中使用`slice`获取数组的切片。事实上是操作一个copy的全新的数组，不会对原切片造成影响。
而go语言不同的是操作使用`slice`获取的切片等同于操作原切片


```go
users := []int{1,2,3,4,5}
slice := users[2:5]
fmt.Println(slice)  //[3 4 5]
slice[1] = 10
fmt.Println(users)  //[1 2 3 10 5]

//等同于
users[2] = 10

```

#### 使用切片

- [a:b]获取切片索引a到索引b的切片
- [:b]获取切片开始到索引b的切片
- [a:]获取索引b到切片结束的切片

> `strings.Index()`获取字符串某个字符所在索引

> `len()`获取切片长度

如下例

```
str := "welcome to beijing"
str1 := strings.Index(str[5:], "o")
fmt.Println(str) // welcome to beijing
fmt.Println(str[5:]) // me to beijing
fmt.Println(str1) // 4

user := []int{1,2,3,4,5}
fmt.Println(user[:len(user)-1]) //  [1 2 3 4]
```


