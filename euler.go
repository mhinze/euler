package main

import (
	"flag"
	"fmt"
)

var problem = flag.Int("p", 0, "the problem to execute")

func main() {
	flag.Parse()

	fmt.Println("Project Euler in Go")
	fmt.Println("Invoking Problem #", *problem)

	switch *problem {
		case 1: one()
		case 2 : two()
		default: fmt.Println("Use -p to specify a problem")
	}
}

type fibResult struct {
	prev,  curr int
}

func (this fibResult) Next() fibResult {
	return fibResult {this.curr, this.prev + this.curr}
}

func two() {
	fmt.Println("By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms")

	inc := 0
	var number fibResult
	number = fibResult {1,2}
	for; number.curr < 4000000; number = number.Next() {
		if number.curr % 2 == 0 {
			inc += number.curr
		}
	}
	fmt.Println(inc)
}

func one() {
	fmt.Println("Find the sum of all the multiples of 3 or 5 below 1000.")
	var x int = 0
	for i := 1000 - 1; i > 0; i-- {
		if i % 3 == 0 || i % 5 == 0 {
			x = x + i
		}
	}
	fmt.Println("Answer:", x)
}

