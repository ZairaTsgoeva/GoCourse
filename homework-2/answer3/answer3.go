package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(fib2(uint64(i)))
	}
}

func fib2(n uint64) uint64 {
	var res uint64 = 1
	if n > 1 {
		var a uint64 = 1
		var b uint64 = 1
		var i uint64 = 2
		var c uint64
		for ; i <= n; i++ {
			c = a + b
			a = b
			b = c
		}

		res = b
	}

	return res
}
