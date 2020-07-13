package service

import "fmt"

// ShowService function
func ShowService(firstNumber, secondNumber, result int) {
	finalResult := 0
	for i := firstNumber; i <= (firstNumber + 90); i += 10 {
		for j := result; j <= (result + 9000); j += 1000 {
			finalResult = secondNumber * i
			if finalResult == j {
				fmt.Printf("%d * %d = %d\n", i, secondNumber, j)
				break
			} else {
				continue
			}
		}
	}
}
