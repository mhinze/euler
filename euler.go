package main

import (
	"flag"
	"fmt"
)

var problem = flag.Int("p", 0, "the problem to execute")

func main() {
	flag.Parse()

	fmt.Println("Project Euler in Go")

	if *problem != 0 {
		fmt.Println("Invoking Problem #", *problem)
	}

	switch *problem {
	case 1:
		one()
	case 2:
		two()
	case 3:
		three()
	case 4:
		four()
	case 5:
		five()
	default:
		fmt.Println("Use -p to specify a problem")
	}
}

func five() {
	fmt.Println("What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?")

	isDiv := func(x int) bool {
		for i := 1; i <= 20; i = i + 1 {
			if x%i != 0 {
				return false
			}
		}
		return true
	}

	answer := -1
	i := 1
	for answer < 0 {
		if isDiv(i) {
			answer = i
		}
		i = i + 1
	}

	fmt.Println("Answer:", answer)
}

func four() {
	fmt.Println("Find the largest palindrome made from the product of two 3-digit numbers.")

	largest := int64(0)
	for i := 999; i > 100; i-- {
		for j := 999; j > 100; j-- {
			value := int64(i * j)
			if value > largest && isPalindromeInt(value) {
				largest = value
			}
		}
	}

	fmt.Println("Answer:", largest)
}

func three() {
	fmt.Println("What is the largest prime factor of the number 600851475143 ?")

	target := 600851475143
	div := 2

	for {
		if target/div == 1 && target%div == 0 {
			break
		}
		if target%div == 0 {
			target /= div
		}
		div += 1
	}

	fmt.Println("Answer:", div)

}

func two() {
	fmt.Println("By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms")

	inc := 0
	for number := fib(); number.curr < 4000000; number.next() {
		if number.curr%2 == 0 {
			inc += number.curr
		}
	}
	fmt.Println("Answer:", inc)
}

func one() {
	fmt.Println("Find the sum of all the multiples of 3 or 5 below 1000.")
	var x int = 0
	for i := 1000 - 1; i > 0; i-- {
		if i%3 == 0 || i%5 == 0 {
			x = x + i
		}
	}
	fmt.Println("Answer:", x)
}
