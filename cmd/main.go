package main

import (
	"Array/array"
	"fmt"
)

func main() {
	fmt.Println("Fibonacci number is index: ", array.Fibonacci(7))
	fmt.Println(array.ToBinary(10))
	array.Kvadrat(3)
	sum := array.List(2, 5)
	fmt.Println(sum)
	array.Natural(1, 9)
	arr := []int{1, 2, 3, 4, 5}
	reverse := array.TeskarArray(arr)
	fmt.Println(reverse)
	num := []int{1, 2, 3, 4, 5}
	sam := array.SmallestMultipleOf(num, 2, 4)
	fmt.Println(sam)
	fmt.Println(array.Pifagor(5, 2, 4))
	green := []int{104}
	f1 := array.White(green)
	fmt.Println(array.White(f1))
	year := array.YearsToGrow(100, 5, 1000)
	fmt.Println(year)
}
