package ch02

import "fmt"

// 2. 문자열 사용하기
func UseString() {

	// 문자열 변수 선언 및 초기화
	var name string = "GoLang" // 명시적 변수 선언
	shortName := "Go"          // 암시적 변수 선언

	// 출력
	fmt.Println("Full Name:", name)
	fmt.Println("Short Name:", shortName)

	// 여러 줄 문자열 선언
	poem :=
		`		뿌리내리는 것		 
						노혜민

		너는 기다리지 않는다.
		그저 네 자리에서
		뿌리를 내려,
		땅을 어루만질 뿐.

		나는 너의 고요를 배운다.
		뿌리내리는 것은 기다림이 아니라,
		살아가는 방식임을.

		오늘도 나는 땅 위에 선다.
		흔들리는 잎사귀로
		하늘에 닿으려 애쓰는 너처럼.
		
		`

	// 출력
	fmt.Println(poem)
}
