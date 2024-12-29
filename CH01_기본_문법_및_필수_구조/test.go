package main // main 패키지를 선언

import "fmt" // fmt 패키지(입/출력함수 제공)

// main 함수는 Go 프로그램의 진입점
func main() {

	num := 42

	fmt.Print("줄 바꿈 없이 출력")
	fmt.Println("출력 후 줄 바꿈을 추가")
	fmt.Printf("서식 지정 출력 -> %d", num)

}
