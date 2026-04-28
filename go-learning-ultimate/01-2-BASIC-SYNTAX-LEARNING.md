# Issue 1.2: 基础语法练习 - 学习知识点

## 目标

完成 Go 语言基础语法的系统学习，掌握变量、常量、数据类型、流程控制等核心概念。先把基础打牢，后面学习函数和更复杂的内容就会轻松一些。

---

## 知识点清单

### 1. 包与入口函数

**核心概念：**
- `package main`：表示这是一个可执行程序，不是库
- `import "fmt"`：导入标准库 fmt，用于格式化输入输出
- `func main()`：程序的入口函数，从这里开始执行

**示例代码：**
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, Go!")
}
```

**关键要点：**
- 每个 Go 文件都必须在第一行声明所属的包
- main 包必须包含 main 函数
- import 语句必须放在 package 之后，其他代码之前
- Go 程序的执行顺序是从 main 函数开始，从上往下执行

---

### 2. 变量声明

Go 有多种变量声明方式，各有适用场景，按你的习惯选择就行。

**方式 1：var 声明（完整形式）**
```go
var name string = "Go"
var age int = 15
```

**方式 2：var 声明（类型推断）**
```go
var name = "Go"  // 编译器自动推断为 string 类型
var age = 15     // 编译器自动推断为 int 类型
```

**方式 3：短变量声明（最常用）**
```go
name := "Go"     // 只能在函数内部使用
age := 15
```

**方式 4：批量声明**
```go
var (
    name = "Go"
    age = 15
    isAwesome = true
)
```

**关键要点：**
- 短变量声明 `:=` 只能在函数内部使用
- 变量声明后必须使用，否则会编译错误（Go 语言的这个设计可以避免声明了却不用的变量）
- 同一作用域内不能重复声明同名变量
- Go 是静态类型语言，变量类型一旦确定就不能改变

---

### 3. 基本数据类型

**数值类型：**
```go
// 整数
var i int = 42
var i8 int8 = 127
var i16 int16 = 32767
var i32 int32 = 2147483647
var i64 int64 = 9223372036854775807

// 无符号整数
var ui uint = 42
var ui8 uint8 = 255
var ui16 uint16 = 65535
var ui32 uint32 = 4294967295
var ui64 uint64 = 18446744073709551615

// 浮点数
var f32 float32 = 3.14
var f64 float64 = 3.141592653589793

// 复数
var c64 complex64 = 1 + 2i
var c128 complex128 = 1 + 2i
```

**字符串类型：**
```go
var s string = "Hello, Go!"
var multiLine string = `第一行
第二行
第三行` // 反引号表示原始字符串，不转义
```

**布尔类型：**
```go
var b bool = true
var notB bool = false
```

**零值概念：**
Go 语言中，如果声明变量但不赋初值，会自动给一个"零值"，这样可以避免未初始化变量带来的问题：
```go
var i int       // 0
var f float64   // 0.0
var s string    // "" （空字符串）
var b bool      // false
```

**关键要点：**
- `int` 的长度取决于操作系统（32位或64位）
- 字符串是 UTF-8 编码的，中文可以正常处理
- 布尔类型只有 `true` 和 `false` 两个值
- 不同类型之间不能直接运算，需要显式类型转换

---

### 4. 常量

**核心概念：**
常量是在编译时就确定值，运行时不能改变的量。适合用来定义不会变的配置或枚举值。

**声明方式：**
```go
const Pi = 3.14159
const (
    StatusOK = 200
    StatusNotFound = 404
    StatusInternalError = 500
)

// iota 枚举
const (
    Monday = iota  // 0
    Tuesday        // 1
    Wednesday      // 2
    Thursday       // 3
    Friday         // 4
    Saturday       // 5
    Sunday         // 6
)
```

**关键要点：**
- 常量不需要 `:=` 声明
- `iota` 是 Go 语言的常量计数器，从 0 开始，每次自动加 1
- 常量可以是字符、字符串、布尔值或数值
- 常量可以不指定类型，由上下文推断

---

### 5. 流程控制 - if/else

**基本语法：**
```go
if age >= 18 {
    fmt.Println("成年人")
} else {
    fmt.Println("未成年人")
}
```

**进阶用法 - if with short statement：**
```go
if age := 18; age >= 18 {
    fmt.Println("成年人")
} else {
    fmt.Println("未成年人")
}
// age 变量只在 if/else 块内有效
```

**关键要点：**
- 条件不需要括号
- 左大括号 `{` 必须和 if/else 在同一行
- 可以在 if 中声明变量，作用域仅限于 if/else 块

---

### 6. 流程控制 - switch

Go 的 switch 比其他语言更灵活，不需要每个 case 都写 break。

**基本用法：**
```go
day := "Monday"
switch day {
case "Monday":
    fmt.Println("周一")
case "Tuesday":
    fmt.Println("周二")
default:
    fmt.Println("其他")
}
```

**高级用法 - switch without condition：**
```go
age := 18
switch {
case age < 12:
    fmt.Println("儿童")
case age < 18:
    fmt.Println("青少年")
case age < 60:
    fmt.Println("成年人")
default:
    fmt.Println("老年人")
}
```

**关键要点：**
- Go 的 switch 自动 break，不需要显式写 break
- 可以用 `fallthrough` 继续执行下一个 case（但不推荐，容易让人困惑）
- case 可以是任意表达式，不一定是常量
- switch 可以不带条件，等价于 switch true

---

### 7. 流程控制 - for 循环

Go 只有 for 循环，没有 while 或 do-while，但通过不同的写法可以实现同样的效果。

**基本循环：**
```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

**while 风格：**
```go
i := 0
for i < 10 {
    fmt.Println(i)
    i++
}
```

**无限循环：**
```go
for {
    fmt.Println("无限循环")
    break // 记得 break，不然会一直跑
}
```

**range 循环：**
```go
nums := []int{1, 2, 3, 4, 5}
for index, value := range nums {
    fmt.Printf("索引: %d, 值: %d\n", index, value)
}

// 如果只想用值，不想用索引，可以用 _ 忽略
for _, value := range nums {
    fmt.Println(value)
}
```

**关键要点：**
- Go 只有 for 循环，没有 while 或 do-while
- `_` 用于忽略不需要的变量
- range 可以遍历数组、切片、字符串、map、channel

---

### 8. 类型转换

**显式类型转换：**
```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)

var s string = string(65) // "A"

// 数值和字符串之间的转换需要用 strconv 包
import "strconv"

num := 42
str := strconv.Itoa(num)         // "42"
parsedNum, _ := strconv.Atoi(str) // 42
```

**关键要点：**
- Go 不支持隐式类型转换，必须显式转换
- 类型转换语法：`T(v)`，T 是目标类型，v 是要转换的值
- 不同类型之间的转换可能会有精度损失
- 字符串和数值的转换需要使用 `strconv` 包

---

## 练习任务

### 练习 1：变量与数据类型
1. 声明各种类型的变量并打印
2. 尝试不同的变量声明方式
3. 验证零值概念
4. 写一个简单程序，演示不同类型的零值

### 练习 2：流程控制
1. 编写一个程序，判断一个数是奇数还是偶数
2. 编写一个程序，打印 1-100 的斐波那契数列
3. 编写一个程序，判断一个年份是不是闰年
4. 用 switch 语句改写上面的判断逻辑（可选）

### 练习 3：综合练习
编写一个简易计算器，支持加减乘除运算。可以先从简单的两个数运算开始，慢慢完善。

---

## 学习检查清单

- [ ] 理解包和 main 函数的作用
- [ ] 掌握 4 种变量声明方式
- [ ] 了解 Go 的基本数据类型
- [ ] 理解零值概念
- [ ] 会使用常量和 iota
- [ ] 掌握 if/else 条件判断
- [ ] 掌握 switch 语句
- [ ] 掌握 for 循环的 4 种写法
- [ ] 会进行基本的类型转换
- [ ] 完成所有练习任务

---

## 下一步

完成基础语法练习后，进入：
- **函数学习**（见 Issue 1.3）

先这样，这些知识点慢慢来，今天能往前推进一点就算 ok 了。
