package main

import "fmt"

const years int = 5

func main() {
	var (
		value   float64
		percent float64
		result  float64
	)
	fmt.Println("Введите сумму ")
	fmt.Scanln(&value)
	fmt.Println("Введите процент")
	fmt.Scanln(&percent)
	result = value
	for n := 0; n < years; n++ {
		result += result * (percent / 100)
	}
	fmt.Printf("%.2f\n", result)

}
