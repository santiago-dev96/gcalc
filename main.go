package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var operation = flag.String("op", "", "the operation to perform, possible values are: add, sub, mul, and div")

func main() {
	flag.Parse()
	if len(flag.Args()) != 2 {
		fmt.Println("you must provide two numbers for the operation")
		os.Exit(1)
	}
	num1, err := strconv.ParseFloat(flag.Arg(0), 64)
	if err != nil {
		fmt.Printf("the value %q cannot be parsed as a number\n", flag.Arg(0))
		os.Exit(1)
	}
	num2, err := strconv.ParseFloat(flag.Arg(1), 64)
	if err != nil {
		fmt.Printf("the value %q cannot be parsed as a number\n", flag.Arg(1))
		os.Exit(1)
	}
	*operation = strings.TrimSpace(*operation)
	var result float64
	switch *operation {
	case "add":
		result = num1 + num2
	case "sub":
		result = num1 - num2
	case "mul":
		result = num1 * num2
	case "div":
		result = num1 / num2
	default:
		fmt.Printf("the operation %q is not valid\n", *operation)
		os.Exit(1)
	}
	fmt.Println(result)
}
