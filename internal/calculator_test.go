package calculator

import (
	"strconv"
	"testing"
)

func TestCalculator(t *testing.T) {
	tests := []struct {
		str     string
		res     int
		success bool
	}{
		{str: "0", res: 0, success: false},
		{str: "+", res: 0, success: false},
		{str: "qew", res: 0, success: false},
		{str: "123 + b", res: 0, success: false},
		{str: "a12+34", res: 0, success: false},
		{str: "12v34-23", res: 0, success: false},
		{str: "0-0", res: 0, success: true},
		{str: "5/2", res: 2, success: true},
		{str: "5/0", res: 0, success: false},
		{str: "5%2", res: 1, success: true},
		{str: "10000+1000", res: 11000, success: true},
		{str: "1+2", res: 3, success: true},
		{str: "3-1", res: 2, success: true},
		{str: "+", res: 0, success: false},
	}

	for _, test := range tests {
		calc := NewCalculator()
		res, err := calc.Calculate(test.str)
		if err == nil {
			if res != test.res {
				t.Error("Expected: ", test.res, " got: ", res, " value ", test.str)
			}
		} else {
			if test.success == true {
				t.Error("Error expected: ", test.success, " got: ", err, " value ", test.str)
			}
		}

	}
}

func FuzzCalcAdd(f *testing.F) {
	f.Fuzz(func(t *testing.T, i int, j int) {
		strIn := strconv.Itoa(i) + "+" + strconv.Itoa(j)

		calc := NewCalculator()
		res, err := calc.Calculate(strIn)

		if err == nil {
			if i+j != res {
				t.Error("Error add numbers: ", strIn, " expected: ", i+j, "got: ", res)
			}
		} else {
			t.Error("Error wasn't expected: ", strIn, " got: ", err.Error())
		}
	})
}

func FuzzCalcMul(f *testing.F) {
	f.Fuzz(func(t *testing.T, i int, j int) {
		strIn := strconv.Itoa(i) + "*" + strconv.Itoa(j)

		calc := NewCalculator()
		res, err := calc.Calculate(strIn)

		if err == nil {
			if i*j != res {
				t.Error("Error add numbers: ", strIn, " expected: ", i*j, "got: ", res)
			}
		} else {
			t.Error("Error wasn't expected: ", strIn, " got: ", err.Error())
		}
	})
}

func FuzzCalcSub(f *testing.F) {
	f.Fuzz(func(t *testing.T, i int, j int) {
		strIn := strconv.Itoa(i) + "-" + strconv.Itoa(j)

		calc := NewCalculator()
		res, err := calc.Calculate(strIn)

		if err == nil {
			if i-j != res {
				t.Error("Error add numbers: ", strIn, " expected: ", i-j, "got: ", res)
			}
		} else {
			t.Error("Error wasn't expected: ", strIn, " got: ", err.Error())
		}
	})
}
