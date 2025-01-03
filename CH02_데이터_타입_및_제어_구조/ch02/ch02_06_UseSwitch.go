package ch02

import "fmt"

// 6. switch, break 사용하기
func UseSwitchTest() {

	// 숫자를 기반으로 switch 문 사용
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	default:
		fmt.Println("Weekend")
	}

	// 문자열을 기반으로 switch 문 사용
	grade := "B"
	switch grade {
	case "A":
		fmt.Println("Excellent")
	case "B":
		fmt.Println("Good")
		fallthrough
	case "C":
		fmt.Println("Average")
	default:
		fmt.Println("Poor")
	}
}
