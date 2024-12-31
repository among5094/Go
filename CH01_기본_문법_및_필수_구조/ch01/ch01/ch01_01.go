package ch01 // ch01 패키지를 선언
import "fmt"

// 함수 이름의 첫 글자를 대문자로!

// 1. 출력해보기
func Print() {

	num := 42 // num 변수를 선언하고 초기화

	// 출력해보기
	fmt.Print("줄 바꿈 없이 출력")
	fmt.Println("출력 후 줄 바꿈을 추가")
	fmt.Printf("서식 지정 출력 -> %d", num)

}

// 2. 변수 사용하기
func Var() {

	age := 25       // 암시적 타입 추론
	name := "Alice" // 문자열 타입 추론

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)

}
