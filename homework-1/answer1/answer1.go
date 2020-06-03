package main

import "fmt"

const rate = 68.49

func main() {
	var rub float64
	fmt.Println("Ввведите сумму в рублях")
	fmt.Scanln(&rub)
	fmt.Printf("%.2f$\n", rub/rate)
}
