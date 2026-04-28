package main

import "fmt"

// 基础函数
func add(a, b int) int {
	return a + b
}

// 多返回值
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// 命名返回值
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // 裸返回，返回已命名的返回值
}

// 变参函数
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// 函数作为参数
func applyOperation(a, b int, op func(int, int) int) int {
	return op(a, b)
}

// 闭包
func makeAdder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

// 闭包 - 计数器
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	// 基础函数调用
	fmt.Printf("3 + 5 = %d\n", add(3, 5))

	// 多返回值
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("错误:", err)
	}

	// 命名返回值
	x, y := split(17)
	fmt.Printf("17 拆分为 %d 和 %d\n", x, y)

	// 变参函数
	fmt.Printf("sum(1, 2, 3) = %d\n", sum(1, 2, 3))
	fmt.Printf("sum(1, 2, 3, 4, 5) = %d\n", sum(1, 2, 3, 4, 5))

	// 使用切片作为变参
	nums := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("sum(nums...) = %d\n", sum(nums...))

	// 函数作为参数
	mul := func(a, b int) int { return a * b }
	fmt.Printf("applyOperation(3, 5, add) = %d\n", applyOperation(3, 5, add))
	fmt.Printf("applyOperation(3, 5, mul) = %d\n", applyOperation(3, 5, mul))

	// 闭包
	add5 := makeAdder(5)
	fmt.Printf("add5(3) = %d\n", add5(3))
	fmt.Printf("add5(10) = %d\n", add5(10))

	// 闭包 - 计数器
	counter1 := makeCounter()
	fmt.Printf("counter1(): %d\n", counter1())
	fmt.Printf("counter1(): %d\n", counter1())
	fmt.Printf("counter1(): %d\n", counter1())

	counter2 := makeCounter()
	fmt.Printf("counter2(): %d\n", counter2())
	fmt.Printf("counter1(): %d\n", counter1())
}
