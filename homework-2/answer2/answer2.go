package main

import "fmt"

func main() {
	var num int
	fmt.Println("Введите число")
	fmt.Scanln(&num)
	value := "не делится на 3"
	if check(num) {
		value = "делится на 3"
	}
	fmt.Println("Число", num, value)
}

func check(num int) bool {
	return (num % 3) == 0
}
