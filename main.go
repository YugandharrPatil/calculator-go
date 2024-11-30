package main

import (
	"fmt"
)

type operation func(a int, b int) int

func main() {
	fmt.Println(`what would you like to do?
	1: add
	2: subtract
	3: multiply
	4: divide
		`)

	var operationID int
	_, err := fmt.Scanf("%d", &operationID)
	if err != nil {
		fmt.Printf("%v", err)
	}
	var num1, num2 int
	fmt.Println("please enter 2 space separated numbers")
	_, err = fmt.Scanf("%d %d", &num1, &num2)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("Result: %d", operate(operationID, num1, num2))
}

func operate(operationID int, a int, b int) int {
	var operation operation
	switch operationID {
	case 1:
		operation = add
	case 2:
		operation = subtract
	case 3:
		operation = multiply
	case 4:
		operation = divide
	}
	return operation(a, b)
}

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func multiply(a int, b int) int {
	return a * b
}

func divide(a int, b int) int {
	return a / b
}
