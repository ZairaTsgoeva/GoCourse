package main

import (
	"GoCourse/homework-4/answer3/calculator"
	"fmt"
)

func main() {
	input := ""
	for {
		fmt.Print("> ")
		if _, err := fmt.Scanln(&input); err != nil {
			fmt.Println(err)
			continue
		}

		if input == "exit" {
			break
		}

		if input == "help" {
			println("Введите выражение без пробелов, пример 2+4 или 5*8")
			println("Можно использовать константы, пример 2+pi")
			println("Доступные константы:")
			for key, value := range calculator.GetConstants() {
				println(key, value)
			}
			println("Можно использовать функции, пример sin(2)")
			println("Доступные функции:")
			for key, _ := range calculator.GetFuncMap() {
				fmt.Printf("%s(x)\n", key)
			}
			continue;
		}

		if res, err := calculator.Calculate(input); err == nil {
			fmt.Printf("Результат: %v\n", res)
		} else {
			fmt.Println("Не удалось произвести вычисление")
		}
	}
}
