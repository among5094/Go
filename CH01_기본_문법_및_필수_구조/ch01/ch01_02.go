package ch01 // ch01 패키지를 선언
import "fmt" // 표준 입출력을 위한 fmt 패키지 임포트

// UseVar 함수는 Go의 다양한 데이터 타입과 그 사용 예를 보여줍니다.
func UseVar() {

	// === 기본 데이터 타입 ===

	// 정수형 데이터 타입 (Integers)
	var intVal int = 10                         // 기본 정수형 (플랫폼에 따라 크기 다름)
	var int8Val int8 = -128                     // 8비트 정수형 (-128 ~ 127)
	var uintVal uint = 100                      // 부호 없는 정수 (0 이상)
	var uint64Val uint64 = 18446744073709551615 // 64비트 부호 없는 정수

	// 실수형 데이터 타입 (Floating-point)
	var float32Val float32 = 3.14        // 32비트 부동소수점
	var float64Val float64 = 2.718281828 // 64비트 부동소수점 (더 높은 정밀도)

	// 복소수형 데이터 타입 (Complex numbers)
	var complex64Val complex64 = 1 + 2i         // 32비트 복소수
	var complex128Val complex128 = 3.14 + 1.59i // 64비트 복소수

	// 불리언 데이터 타입 (Boolean)
	var isGoAwesome bool = true // 참/거짓을 나타냄

	// 문자 데이터 타입 (Rune, UTF-8)
	var runeVal rune = 'G' // 'G' 문자의 UTF-8 값 (정수형)

	// 바이트 데이터 타입 (Byte, uint8의 별칭)
	var byteVal byte = 255 // 8비트 부호 없는 정수

	// 문자열 데이터 타입 (String)
	var stringVal string = "Hello, Go!" // 문자열

	// === 복합 데이터 타입 ===

	// 포인터 데이터 타입 (Pointer)
	var intPointer *int = &intVal // 변수 intVal의 메모리 주소를 저장

	// 배열 데이터 타입 (Array)
	var myArray [3]int = [3]int{1, 2, 3} // 정수 3개를 담는 고정 길이 배열

	// 슬라이스 데이터 타입 (Slice)
	var mySlice []int        // 빈 슬라이스 선언
	mySlice = []int{1, 2, 3} // 초기화하면서 선언

	// 맵 데이터 타입 (Map)
	var mapVal map[string]int = map[string]int{"Alice": 25, "Bob": 30} // 키-값 쌍으로 구성된 자료구조

	// 채널 데이터 타입 (Channel)
	channelVal := make(chan int) // 동시성 프로그래밍에서 사용되는 채널 생성

	// 인터페이스 데이터 타입 (Interface)
	var anyVal interface{} = "I can hold any value" // 임의의 데이터 타입을 저장 가능

	// === 출력 ===

	// 기본 데이터 타입 출력
	fmt.Println("=== 기본 데이터 타입 ===")
	fmt.Println("Integer (int):", intVal)
	fmt.Println("Integer (int8):", int8Val)
	fmt.Println("Unsigned Integer (uint):", uintVal)
	fmt.Println("Unsigned Integer (uint64):", uint64Val)
	fmt.Println("Floating-point (float32):", float32Val)
	fmt.Println("Floating-point (float64):", float64Val)
	fmt.Println("Complex (complex64):", complex64Val)
	fmt.Println("Complex (complex128):", complex128Val)
	fmt.Println("Boolean:", isGoAwesome)
	fmt.Println("Rune (UTF-8):", runeVal, string(runeVal)) // UTF-8 값과 문자 출력
	fmt.Println("Byte (uint8):", byteVal)
	fmt.Println("String:", stringVal)

	// 복합 데이터 타입 출력
	fmt.Println("\n=== 복합 데이터 타입 ===")
	fmt.Println("Pointer:", intPointer, *intPointer) // 포인터 주소와 해당 값 출력
	fmt.Println("Array:", myArray)                   // 배열 출력
	fmt.Println("Slice:", mySlice)                   // 슬라이스 출력
	fmt.Println("Map:", mapVal)                      // 맵 출력

	// 채널 사용: 값을 쓰고 읽기
	fmt.Println("Channel (write): writing 42 to channel")
	go func() { channelVal <- 42 }()             // 별도 고루틴에서 채널에 값 쓰기
	fmt.Println("Channel (read):", <-channelVal) // 채널에서 값 읽기

	// 인터페이스 출력
	fmt.Println("Interface:", anyVal) // 인터페이스에 저장된 값 출력

	// 여러 변수를 동시에 선언하고 초기화 해보기
	x, y, z := 1, 2, 3
	fmt.Println(x, y, z)

	a, b, c := "Go", true, 3.14
	fmt.Println(a, b, c)

	// 여러 변수를 묶음으로 선언
	var (
		name  = "Alice"
		age   = 25
		alive = true
	)
	fmt.Println(name, age, alive)

}
