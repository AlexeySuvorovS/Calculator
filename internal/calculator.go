package calculator

import (
	"errors"
	"strconv"
	"strings"
)

type Calculator struct {
	operationsMap map[string]func(int, int) (int, error)
}

func NewCalculator() *Calculator {
	calc := Calculator{}

	calc.operationsMap = make(map[string]func(int, int) (int, error), 4)

	calc.operationsMap = map[string]func(int, int) (int, error){
		"+": func(a int, b int) (int, error) { return a + b, nil },
		"-": func(a int, b int) (int, error) { return a - b, nil },
		"*": func(a int, b int) (int, error) { return a * b, nil },
		"%": func(a int, b int) (int, error) { return a % b, nil },
		"/": func(a int, b int) (int, error) {
			if b == 0 {
				err := errors.New("second number is zero")
				return 0, err
			}
			return a / b, nil
		},
	}

	return &calc
}

func (c *Calculator) Calculate(s string) (int, error) {
	var err error
	bs := []byte(s)
	var sIdx int
	for i, l := range bs {
		if l == '+' ||
			l == '-' ||
			l == '*' ||
			l == '%' ||
			l == '/' {
			sIdx = i
			break
		}
	}
	if sIdx == 0 {
		return 0, errors.New("math symbol not found")
	}

	symbol := string(bs[sIdx])

	first := bs[0:sIdx]
	second := bs[sIdx+1:]

	var a, b int
	a, err = strconv.Atoi(strings.ReplaceAll(string(first), " ", ""))
	if err != nil {
		return 0, err
	}

	b, err = strconv.Atoi(strings.ReplaceAll(string(second), " ", ""))
	if err != nil {
		return 0, err
	}

	res, err := c.operationsMap[symbol](a, b)
	return res, err
}
