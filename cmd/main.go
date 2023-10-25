package main

import (
	calculator "cmd/main.go/internal"
	"flag"
	"fmt"
)

func main() {
	flag.Parse()

	strIn := flag.Arg(0)
	fmt.Println(strIn)

	if strIn == "" {
		strIn = "3-1"
		fmt.Println("run with args. example: \"", strIn, "\"")
	}
	calc := calculator.NewCalculator()
	res, err := calc.Calculate(strIn)
	if err != nil {
		fmt.Println(err.Error(), "for string: ", strIn)
	} else {
		fmt.Println(strIn, "=", res)
	}
}
