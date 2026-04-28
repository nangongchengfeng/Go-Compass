package main

import "fmt"

func main() {
	// if/else
	score := 85
	if score >= 90 {
		fmt.Println("优秀")
	} else if score >= 80 {
		fmt.Println("良好")
	} else if score >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}

	// if 带短语句
	if num := 10; num%2 == 0 {
		fmt.Printf("%d 是偶数\n", num)
	} else {
		fmt.Printf("%d 是奇数\n", num)
	}

	// for 循环 - 基础
	fmt.Println("\n基础 for 循环:")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// for 循环 - 类似 while
	fmt.Println("\n类似 while 的循环:")
	j := 0
	for j < 5 {
		fmt.Printf("%d ", j)
		j++
	}
	fmt.Println()

	// for 循环 - 无限循环 + break
	fmt.Println("\n无限循环 + break:")
	k := 0
	for {
		if k >= 5 {
			break
		}
		fmt.Printf("%d ", k)
		k++
	}
	fmt.Println()

	// for 循环 - continue
	fmt.Println("\ncontinue 跳过偶数:")
	for m := 0; m < 10; m++ {
		if m%2 == 0 {
			continue
		}
		fmt.Printf("%d ", m)
	}
	fmt.Println()

	// switch
	fmt.Println("\nSwitch:")
	day := "周一"
	switch day {
	case "周一", "周二", "周三", "周四", "周五":
		fmt.Println("工作日")
	case "周六", "周日":
		fmt.Println("周末")
	default:
		fmt.Println("无效的日期")
	}

	// switch 带表达式
	fmt.Println("\nSwitch 带表达式:")
	hour := 14
	switch {
	case hour < 12:
		fmt.Println("上午")
	case hour < 18:
		fmt.Println("下午")
	default:
		fmt.Println("晚上")
	}

	// switch 带短语句
	fmt.Println("\nSwitch 带短语句:")
	switch n := 15; {
	case n%3 == 0 && n%5 == 0:
		fmt.Println("FizzBuzz")
	case n%3 == 0:
		fmt.Println("Fizz")
	case n%5 == 0:
		fmt.Println("Buzz")
	default:
		fmt.Println(n)
	}
}
