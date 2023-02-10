package main

import "fmt"

func main() {
	// testProblem1()
	testProblem2()
}

func testProblem1() {
	sum := findSumsOfMultiplesOf3Or5(1000)
	fmt.Println(sum)
}

func testProblem2() {
	sum := sumOfEvenFibonacci(4_000_000)
	fmt.Println(sum)
}
