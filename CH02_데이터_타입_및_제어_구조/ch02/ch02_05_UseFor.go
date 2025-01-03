package ch02

import "fmt"

// 5. for 반복문 사용하기
func UseForTest() {

	//  기본적인 for 반복문
	for i := 0; i < 5; i++ {
		fmt.Println("반복:", i) // 반복 횟수 출력
	}

	// 슬라이스를 순회하는 for 반복문

	// numbers 슬라이스를 선언
	numbers := []int{1, 2, 3, 4, 5}

	// range를 사용한 for 반복문
	for _, number := range numbers {
		fmt.Println("Number:", number)
	}

	// 인덱스를 사용한 for 반복문
	for i := 0; i < len(numbers); i++ {
		fmt.Println("Number:", numbers[i]) // 인덱스로 요소 접근
	}

	// 1부터 10까지의 합 계산
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("1부터 10까지의 합: ", sum)
}
