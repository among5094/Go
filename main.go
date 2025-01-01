package main     // main 패키지를 선언
import "GO/ch01" // 모듈 이름을 기준으로 ch01 패키지 import

// main 함수는 Go 프로그램의 진입점
func main() {

	ch01.Print() // ch01 패키지의 Print 함수 호출
	print("\n\n\n")
	ch01.UseVar()

}
