package main

import "fmt"

func main() {
	var num int
	fmt.Println("Введите число")
	fmt.Scanln(&num)
	if check(num) {
		fmt.Println("Число делится на 3")
	} else {
		fmt.Println("Число не делится")
	}
}

func check(num int) bool {
	return (num % 3) == 0
}
