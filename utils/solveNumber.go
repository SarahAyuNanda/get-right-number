package utils

import (
	"strconv"
	"strings"

	"github.com/golang/sarahayunanda/goday6/complete_number/service"
)

// Solving function
func Solving(expression string) {
	splitExpression := strings.Split(expression, " ")
	firstOperand := splitExpression[0]
	secondOperand := splitExpression[2]
	operationResult := splitExpression[4]

	firstNumber := GetFirstNumber(firstOperand)
	secondNumber, _ := strconv.Atoi(secondOperand)
	result := GetResult(operationResult)

	service.ShowService(firstNumber, secondNumber, result)
}
