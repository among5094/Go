package ch02

import "fmt"

// 4. if 조건문 사용하기
func UseIfTest() {

	// 나이를 기준으로 성인 여부 판단
	age := 20
	if age >= 18 {
		fmt.Println("성인")
	} else {
		fmt.Println("성인이 아닙니다.")
	}

	// 점수를 기준으로 등급을 매기기
	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else {
		fmt.Println("Grade: C")
	}
}
