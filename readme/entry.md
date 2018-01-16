### 入门

###### 包的导入

```go
import "math"
import "fmt"

```

或者

```go
import (
	"fmt"
	"math"
)
```


###### 在GO中，首字母大写的名称是可以导出的，任何未导出的名字是不能被包外的代码访问的

```go
fmt.Println(math.Pi)

//错误的写法
//fmt.Println(math.pi)

```



###### 参数类型简写

```

func add(x int,y int)int  {
	return x + y
}
//连续参数类型若一致，可简写为

func plus(x,y int)int  {
	return x - y
}

```



