package ch01 // ch01 패키지를 선언
import "fmt"

// 2. 변수 사용하기
func UseVar() {

	age := 25       // 암시적 타입 추론
	name := "Alice" // 문자열 타입 추론

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)

}
