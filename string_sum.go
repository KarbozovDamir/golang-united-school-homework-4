package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

func StringSum(input string) (output string, err error) {

	input = strings.ReplaceAll(input, " ", "")
	if input == "" {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}

	nStr := [2]string{}

	switch {
	case strings.Contains(input, "+"):
		sep := strings.Split(input, "+")
		if len(sep) != 2 {
			return "", fmt.Errorf("%w", errorNotTwoOperands)
		}
		nStr[0], nStr[1] = sep[0], sep[1]
	case strings.Contains(input, "-"):
		last := strings.LastIndex(input, "-")
		if len(strings.Split(input, "-")) > 3 {
			return "", fmt.Errorf("%w", errorNotTwoOperands)
		}
		nStr[0], nStr[1] = input[:last], input[last:]
	default:
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}

	firstOperand, erfO := strconv.Atoi(nStr[0])
	if erfO != nil {
		return "", fmt.Errorf("%w", erfO)
	}

	secondOperand, ersO := strconv.Atoi(nStr[1])
	if ersO != nil {
		return "", fmt.Errorf("%w", ersO)
	}

	return strconv.Itoa(firstOperand + secondOperand), nil
}
