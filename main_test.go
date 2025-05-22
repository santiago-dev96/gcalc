package main

import (
	"bytes"
	"errors"
	"flag"
	"io"
	"strings"
	"testing"
)

func TestPrintHelp(t *testing.T) {
	buf := new(bytes.Buffer)
	flag.CommandLine.SetOutput(buf)
	flag.PrintDefaults()
	data, _ := io.ReadAll(buf)
	expected := strings.Join([]string{
		Syntax,
		"",
		string(data),
	}, "\n")
	buf = new(bytes.Buffer)
	PrintHelp(buf)
	data, _ = io.ReadAll(buf)
	got := string(data)
	if got != expected {
		t.Error("help strings doesn't match")
	}
}

func TestParseOperation(t *testing.T) {
	type output struct {
		operation Operation
		anError   bool
	}
	type testCase struct {
		input    string
		expected output
	}
	cases := []testCase{
		{"asd", output{"", true}},
		{"add", output{AddOperation, false}},
		{"   sub  ", output{SubtractOperation, false}},
		{"  div", output{DivideOperation, false}},
		{" other  ", output{"", true}},
	}
	for _, tc := range cases {
		op, err := ParseOperation(tc.input)
		if tc.expected.anError && err == nil {
			t.Errorf("expected an error for input %q", tc.input)
		} else if !tc.expected.anError && err != nil {
			t.Errorf("expected no error for input %q", tc.input)
		} else if op != tc.expected.operation {
			t.Errorf("expected %q, got %q", tc.expected.operation, op)
		}
	}
}

func TestPrintError(t *testing.T) {
	type testCase struct {
		input    error
		expected string
	}
	cases := []testCase{
		{errors.New("something bad happened"), "ERROR: something bad happened\n\n"},
		{errors.New("Oops! something went wrong"), "ERROR: Oops! something went wrong\n\n"},
	}
	for _, tc := range cases {
		buf := new(bytes.Buffer)
		PrintError(buf, tc.input)
		data, _ := io.ReadAll(buf)
		output := string(data)
		if tc.expected != output {
			t.Errorf("expected %q, got %q", tc.expected, output)
		}
	}
}

func TestOperate(t *testing.T) {
	type input struct {
		operation  Operation
		num1, num2 float64
	}
	type output struct {
		num           float64
		errorExpected bool
	}
	type testCase struct {
		input    input
		expected output
	}
	cases := []testCase{
		{input{AddOperation, 13, 3}, output{16, false}},
		{input{"   other ", 5.0, -10}, output{0, true}},
		{input{MultiplyOperation, 67, -10}, output{-670, false}},
		{input{DivideOperation, 27, 3}, output{9, false}},
	}
	for _, tc := range cases {
		output, err := Operate(tc.input.operation, tc.input.num1, tc.input.num2)
		if tc.expected.errorExpected && err == nil {
			t.Errorf("expected an error")
		} else if !tc.expected.errorExpected && err != nil {
			t.Errorf("did not expect an error")
		} else if tc.expected.num != output {
			t.Errorf("expected %f, got %f", tc.expected.num, output)
		}
	}
}

func TestParseNumbers(t *testing.T) {
	type input struct {
		num1, num2 string
	}
	type output struct {
		num1, num2    float64
		errorExpected bool
	}
	type testCase struct {
		input  input
		output output
	}
	cases := []testCase{
		{input{"abc", "11"}, output{0, 0, true}},
		{input{"1", "2"}, output{1, 2, false}},
		{input{"-8", "-11.97"}, output{-8, -11.97, false}},
		{input{"", "3"}, output{0, 0, true}},
		{input{"   13   ", "   -7.45"}, output{13, -7.45, false}},
	}
	for _, tc := range cases {
		num1, num2, err := ParseNumers(tc.input.num1, tc.input.num2)
		if err != nil && !tc.output.errorExpected {
			t.Error("expected no error")
		} else if err == nil && tc.output.errorExpected {
			t.Error("expected an error")
		} else if tc.output.num1 != num1 {
			t.Errorf("expected %f, got %f", tc.output.num1, num1)
		} else if tc.output.num2 != num2 {
			t.Errorf("expected %f, got %f", tc.output.num2, num2)
		}
	}
}
