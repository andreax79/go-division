package division

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
	checkDivision(32798, 250)
	checkDivision(327989, 11)
	checkDivision(3279890, 7654)
	checkDivision(32798901, 76543)
	checkDivision(327989012, 765432)
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
	checkDivision(32798, 250)
	checkDivision(327989, 11)
	checkDivision(3279890, 7654)
	checkDivision(32798901, 76543)
	checkDivision(327989012, 765432)
}
