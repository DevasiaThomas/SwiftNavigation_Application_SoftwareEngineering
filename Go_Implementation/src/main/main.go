package main

import (
	"fmt"

	"../fib"
	"../prime"
)

func main() {
	var a int
	for {
		fmt.Println("Enter number of terms required for Fibonacci Series")
		_, err := fmt.Scanf("%d", &a)
		if err == nil {
			if a > 0 {
				fmt.Printf("Taking integer %d found in input\n", a)
				break
			}
		}
	}
	series := fib.Fib(a)
	for i := 0; i < a; i++ {
		if i != 0 {
			fmt.Print(", ")
		}
		if i < 3 {
			fmt.Printf("%d", series[i])
			continue
		}
		if i > 2 && i < 6 {
			fmt.Print("BuzzFizz")
			continue
		}
		str := ""
		if series[i]%5 == 0 {
			str = str + "Fizz"
		}
		if series[i]%3 == 0 {
			str = str + "Buzz"
		}
		if prime.Prime(series[i]) {
			str = "BuzzFizz"
		}
		if str != "" {
			fmt.Print(str)
		} else {
			fmt.Printf("%d", series[i])
		}
	}
}
