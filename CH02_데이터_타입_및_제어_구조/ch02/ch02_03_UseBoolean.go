package ch02

import "fmt"

// 3. 불 사용하기
func UseBooleanTest() {

	// boolean 타입 변수 선언
	var Student bool = true // 학생인지
	var Expert bool = false // 전문가인지

	// bool 값 출력
	fmt.Println("학생인가요? ", Student)
	fmt.Println("전문가인가요? ", Expert)

	// 출력
	result := Student && !Expert
	fmt.Println("논리 연산 결과: ", result)
}
