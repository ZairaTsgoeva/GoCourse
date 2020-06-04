package main

import "fmt"

func main() {
	var num int
	fmt.Println("Введите число")
	fmt.Scanln(&num)
	if isEven(num) {
		fmt.Println("Число четное")
	} else {
		fmt.Println("Число нечетное")
	}
}

func isEven(num int) bool {
	return (num % 2) == 0
}
