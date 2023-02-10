package main

// 0 1 2 3 5 8
func sumOfEvenFibonacci(n int) int {
	first := 1
	second := 2
	sum := 0
	for {
		if second >= n {
			break
		}

		if second%2 == 0 {
			sum += second
		}

		first, second = second, first+second
	}

	return sum
}
