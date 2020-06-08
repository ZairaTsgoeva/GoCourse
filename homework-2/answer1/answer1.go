package main

import "fmt"

func main() {
	var num int
	fmt.Println("Введите число")
	fmt.Scanln(&num)
	value := "нечетное"
	if isEven(num) {
		value = "четное"
	}
	fmt.Println("Число", num, value)
}

func isEven(num int) bool {
	return (num % 2) == 0
}
