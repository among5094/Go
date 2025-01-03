package ch02

import (
	"fmt"
	"reflect"
)

// 1. 정수 사용하기
func UseNumber() {

	// 정수형 변수 선언 및 초기화
	var intNumber int = 42 // 명시적 선언
	shortInt := 100        // 암시적 선언

	// 출력
	fmt.Println("명시적 선언:", intNumber)
	fmt.Println("암시적 선언:", shortInt)

	// 다양한 정수 타입
	var int8Value int8 = 127                   // 8비트 정수 (최대값 127)
	var int64Value int64 = 9223372036854775807 // 64비트 정수
	var uintValue uint = 1000                  // 부호 없는 정수 (양수만 가능)

	// 출력
	fmt.Println("8비트 정수 (int8):", int8Value)
	fmt.Println("64비트 정수 (int64):", int64Value)
	fmt.Println("부호 없는 정수 (uint):", uintValue)

	// 숫자 연산

	var num1 int = 1
	var num2 int = 2

	// 기본 연산: 덧셈, 뺄셈, 곱셈, 나눗셈, 나머지 연산
	sum := num1 + num2
	sub := num1 - num2
	mul := num1 * num2
	div := num1 / num2
	mod := num1 % num2

	// 연산 결과 출력
	fmt.Println("num1 값: ", num1)
	fmt.Println("num2 값: ", num2)
	fmt.Println("덧셈 결과:", sum)
	fmt.Println("뺄셈 결과:", sub)
	fmt.Println("곱셈 결과:", mul)
	fmt.Println("나눗셈 결과:", div)
	fmt.Println("나머지 연산 결과:", mod)

	// 형 변환
	var value1 int8 = 100  // int8 타입
	var value2 int16 = 200 // int16 타입

	// 타입 변환 후 연산
	fmt.Println("value1 변수타입:", reflect.TypeOf(value1))
	fmt.Println("value2 변수타입:", reflect.TypeOf(value2))

	// 출력
	result := int16(value1) + value2 // int8 → int16으로 변환
	fmt.Println("타입 변환 후 결과:", result)

	// 오버플로우 주의!
	var overflow int8 = 127 // int8 범위는 -128~127
	overflow += 1           // 오버플로우 발생 (127 + 1 = -128)
	fmt.Println("오버플로우 예제:", overflow)
}
