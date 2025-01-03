package main

import (
	"GO/ch01"
	"GO/ch02"
)

// main 함수는 Go 프로그램의 진입점
func main() {

	print(" ===== CH01 ===== \n")
	ch01.Print() // ch01 패키지의 Print 함수 호출
	print("\n\n\n")
	ch01.UseVar()

	print(" ===== CH02 ===== \n")
	ch02.UseString()  // ch02_02.go
	ch02.UseBoolean() // ch02_03.go
	ch02.UseIf()      // ch02_04_UseIf.go

}
