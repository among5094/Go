package ch01 // ch01 패키지를 선언
import "fmt"

// 2. 변수 사용하기
func UseVar() {

	// 정수 (Integers)
	var intVal int = 10
	var int8Val int8 = -128
	var uintVal uint = 100
	var uint64Val uint64 = 18446744073709551615

	// 실수 (Floating-point)
	var float32Val float32 = 3.14
	var float64Val float64 = 2.718281828

	// 복소수 (Complex numbers)
	var complex64Val complex64 = 1 + 2i
	var complex128Val complex128 = 3.14 + 1.59i

	// 불 (Boolean)
	var isGoAwesome bool = true

	// 문자 (Rune, UTF-8)
	var runeVal rune = 'G' // 'G'의 UTF-8 값

	// 바이트 (Byte, uint8 별칭)
	var byteVal byte = 255

	// 문자열 (String)
	var stringVal string = "Hello, Go!"

	// 포인터 (Pointer)
	var intPointer *int = &intVal

	// 배열 (Array)
	var arrayVal [3]int = [3]int{1, 2, 3}

	// 슬라이스 (Slice)
	var sliceVal []string = []string{"Go", "Python", "Java"}

	// 맵 (Map)
	var mapVal map[string]int = map[string]int{"Alice": 25, "Bob": 30}

	// 채널 (Channel)
	channelVal := make(chan int)

	// 인터페이스 (Interface)
	var anyVal interface{} = "I can hold any value"

	// 출력
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
	fmt.Println("Rune (UTF-8):", runeVal, string(runeVal))
	fmt.Println("Byte (uint8):", byteVal)
	fmt.Println("String:", stringVal)

	fmt.Println("\n=== 복합 데이터 타입 ===")
	fmt.Println("Pointer:", intPointer, *intPointer)
	fmt.Println("Array:", arrayVal)
	fmt.Println("Slice:", sliceVal)
	fmt.Println("Map:", mapVal)
	fmt.Println("Channel (write): writing 42 to channel")
	go func() { channelVal <- 42 }() // 채널에 값 쓰기
	fmt.Println("Channel (read):", <-channelVal)
	fmt.Println("Interface:", anyVal)

}
