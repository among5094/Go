package ch02 // ch01 패키지를 선언
import "fmt"

// 2. 문자열 사용하기
func UseString() {
	var name string = "GoLang"
	shortName := "Go"
	fmt.Println("Full Name:", name)
	fmt.Println("Short Name:", shortName)

	multiLine := `This is a
multi-line string in Go.`
	fmt.Println("Multi-line string:", multiLine)
}

// 3. 불 사용하기
func UseBoolean() {
	var isLearning bool = true
	var isExpert bool = false
	fmt.Println("Is Learning:", isLearning)
	fmt.Println("Is Expert:", isExpert)

	result := isLearning && !isExpert
	fmt.Println("Result of logical operations:", result)
}

// 4. if 조건문 사용하기
func UseIf() {
	age := 20
	if age >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Not an adult")
	}

	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else {
		fmt.Println("Grade: C")
	}
}

// 5. for 반복문 사용하기
func UseFor() {
	for i := 0; i < 5; i++ {
		fmt.Println("Iteration:", i)
	}

	numbers := []int{1, 2, 3, 4, 5}
	for _, number := range numbers {
		fmt.Println("Number:", number)
	}

	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("Sum from 1 to 10:", sum)
}

// 6. switch, break 사용하기
func UseSwitch() {
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
