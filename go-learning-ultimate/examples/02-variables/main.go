package main

import (
	"fmt"
)

func main() {
	// 变量声明
	var name string = "Gopher"
	var age int = 7
	var isHappy bool = true

	// 短变量声明
	city := "北京"
	temperature := 25.5

	// 打印变量
	fmt.Printf("姓名: %s\n", name)
	fmt.Printf("年龄: %d\n", age)
	fmt.Printf("开心: %t\n", isHappy)
	fmt.Printf("城市: %s\n", city)
	fmt.Printf("温度: %.1f°C\n", temperature)

	// 常量声明
	const Pi = 3.1415926
	const (
		StatusOK   = 200
		StatusNotFound = 404
	)

	fmt.Printf("Pi: %.4f\n", Pi)
	fmt.Printf("状态码 OK: %d\n", StatusOK)

	// 基本数据类型
	var (
		intVar     int     = 42
		int8Var    int8    = 127
		int16Var   int16   = 32767
		int32Var   int32   = 2147483647
		int64Var   int64   = 9223372036854775807
		uintVar    uint    = 42
		float32Var float32 = 3.14
		float64Var float64 = 3.1415926535
	)

	fmt.Printf("int: %d\n", intVar)
	fmt.Printf("int8: %d\n", int8Var)
	fmt.Printf("int16: %d\n", int16Var)
	fmt.Printf("int32: %d\n", int32Var)
	fmt.Printf("int64: %d\n", int64Var)
	fmt.Printf("uint: %d\n", uintVar)
	fmt.Printf("float32: %.2f\n", float32Var)
	fmt.Printf("float64: %.6f\n", float64Var)
}
