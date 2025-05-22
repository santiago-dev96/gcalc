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
		HandleError("you must provide two numbers for the operation")
	}
	num1, err := strconv.ParseFloat(flag.Arg(0), 64)
	if err != nil {
		HandleError(fmt.Sprintf("the value %q cannot be parsed as a number", flag.Arg(0)))
	}
	num2, err := strconv.ParseFloat(flag.Arg(1), 64)
	if err != nil {
		HandleError(fmt.Sprintf("the value %q cannot be parsed as a number", flag.Arg(1)))
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
		HandleError(fmt.Sprintf("the operation %q is not valid", *operation))
	}
	fmt.Println(result)
}

func PrintHelp() {
	fmt.Fprintln(os.Stderr, "SYNTAX: gcalc -op {add|sub|mul|div} num1 num2")
	fmt.Println()
	flag.PrintDefaults()
	fmt.Println()
}

func HandleError(errMsg string) {
	fmt.Println("ERROR:", errMsg)
	fmt.Println()
	PrintHelp()
	os.Exit(1)
}
