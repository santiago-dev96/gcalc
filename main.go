package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var operation = flag.String("op", "", "the operation to perform, possible values are: add, sub, mul, and div")
var help = flag.Bool("h", false, "prints the help text")

type Operation string

const (
	AddOperation      Operation = "add"
	SubtractOperation Operation = "sub"
	MultiplyOperation Operation = "mul"
	DivideOperation   Operation = "div"
)

func main() {
	flag.Parse()
	if *help {
		PrintHelp(os.Stderr)
		os.Exit(0)
	}
	operation, err := ParseOperation(*operation)
	if err != nil {
		PrintError(os.Stderr, err)
		PrintHelp(os.Stderr)
		os.Exit(1)
	}
	n1, n2, err := ParseNumers(flag.Arg(0), flag.Arg(1))
	if err != nil {
		PrintError(os.Stderr, err)
		PrintHelp(os.Stderr)
		os.Exit(1)
	}
	result, err := Operate(operation, n1, n2)
	if err != nil {
		PrintError(os.Stderr, err)
		PrintHelp(os.Stderr)
		os.Exit(1)
	}
	fmt.Println(result)
}

const Syntax = "SYNTAX: gcalc -op {add|sub|mul|div} num1 num2"

func PrintHelp(dst io.Writer) {
	fmt.Fprintln(dst, Syntax)
	fmt.Fprintln(dst)
	flag.CommandLine.SetOutput(dst)
	flag.PrintDefaults()
}

func PrintError(dst io.Writer, err error) {
	fmt.Fprintln(dst, "ERROR:", err)
	fmt.Fprintln(dst)
}

var ValidOperations = [4]string{"add", "mul", "div", "sub"}

func ParseOperation(operation string) (Operation, error) {
	operation = strings.TrimSpace(operation)
	for _, op := range ValidOperations {
		if op == operation {
			return Operation(operation), nil
		}
	}
	return "", fmt.Errorf("invalid operation %q", operation)
}

func Operate(operation Operation, num1, num2 float64) (float64, error) {
	switch operation {
	case AddOperation:
		return num1 + num2, nil
	case SubtractOperation:
		return num1 - num2, nil
	case MultiplyOperation:
		return num1 * num2, nil
	case DivideOperation:
		return num1 / num2, nil
	}
	// this case never happens
	return 0, fmt.Errorf("invalid operation %q", operation)
}

func ParseNumers(num1, num2 string) (float64, float64, error) {
	n1, err := strconv.ParseFloat(strings.TrimSpace(num1), 64)
	if err != nil {
		return 0, 0, err
	}
	n2, err := strconv.ParseFloat(strings.TrimSpace(num2), 64)
	if err != nil {
		return 0, 0, err
	}
	return n1, n2, nil
}
