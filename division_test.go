package main

import (
	"strings"
	"testing"
)

func TestNumberOfDigits(t *testing.T) {
	if NumberOfDigits(100) != 3 {
		t.Fail()
	}
	if NumberOfDigits(99) != 2 {
		t.Fail()
	}
	if NumberOfDigits(0) != 1 {
		t.Fail()
	}
}

func TestGetIntArg(t *testing.T) {
	args := []string{"test", "10", "20"}
	GetIntArg(args, 1)
	GetIntArg(args, 2)
}

func TestParseArgs(t *testing.T) {
	args := []string{"test", "9876", "543"}
	dividend, divisor := ParseArgs(args)
	if dividend != 9876 {
		t.Fatalf("dividend != 9876")
	}
	if divisor != 543 {
		t.Fatalf("divisor != 543")
	}
}

func TestFormatStep(t *testing.T) {
	checkDivision := func(dividend int, divisor int) {
		division := NewDivision(dividend, divisor)
		division.Calculate()
		i := strings.Index(division.Steps[0].FormatStep(division), "|")
		for _, step := range division.Steps {
			if i != strings.Index(step.FormatStep(division), "|") {
				t.Fatalf("wrong | position")
			}
		}
	}
	checkDivision(3279, 1)
	checkDivision(3279, 2)
	checkDivision(3279, 25)
	checkDivision(3279, 250)
	checkDivision(3279, 2500)
	checkDivision(3279, 25000)
}

func TestDivision(t *testing.T) {
	checkDivision := func(dividend int, divisor int) {
		division := NewDivision(dividend, divisor)
		division.Calculate()
		if dividend/divisor != division.Result {
			t.Fatalf("wrong remainder dividend: %d divisor: %d result: %d expected: %d", dividend, divisor, dividend/divisor, division.Result)
		}
		if dividend%divisor != division.Remainder {
			t.Fatalf("wrong remainder dividend: %d divisor: %d remainder: %d expected: %d", dividend, divisor, dividend%divisor, division.Remainder)
		}
	}
	checkDivision(3279, 1)
	checkDivision(3279, 2)
	checkDivision(3279, 25)
	checkDivision(3279, 250)
	checkDivision(3279, 2500)
	checkDivision(3279, 25000)
}
