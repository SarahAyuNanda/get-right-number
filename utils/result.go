package utils

import (
	"strconv"
	"strings"
)

// GetResult function
func GetResult(operationResult string) (result int) {
	splitResult := strings.Split(operationResult, "")
	joinResult := ""
	for i := 0; i < len(splitResult); i++ {
		if splitResult[i] == "#" {
			splitResult[i] = "0"
		}
		joinResult = strings.Join(splitResult, "")
	}
	result, _ = strconv.Atoi(joinResult)
	return
}
