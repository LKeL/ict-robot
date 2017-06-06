package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func b() int {
	return 1
}

func main() {
	a := func() int {
		return 1
	}
	pos, neg := adder(), adder()
	fmt.Println(a())
	fmt.Println(b())
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
