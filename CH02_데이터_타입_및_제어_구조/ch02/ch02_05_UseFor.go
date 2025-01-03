package ch02

import "fmt"

// 5. for 반복문 사용하기
func UseFor() {
	for i := 0; i < 5; i++ {
		fmt.Println("Iteration:", i)
	}

	numbers := []int{1, 2, 3, 4, 5}
	for _, number := range numbers {
		fmt.Println("Number:", number)
	}

	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("Sum from 1 to 10:", sum)
}
