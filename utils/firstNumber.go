package utils

import (
	"strconv"
	"strings"
)

// GetFirstNumber function
func GetFirstNumber(firstOperand string) (number int) {
	splitFirstOperand := strings.Split(firstOperand, "")
	joinNumbers := ""
	for i := 0; i < len(splitFirstOperand); i++ {
		if splitFirstOperand[i] == "#" {
			splitFirstOperand[i] = "0"
		}
		joinNumbers = strings.Join(splitFirstOperand, "")
	}
	number, _ = strconv.Atoi(joinNumbers)
	return
}
