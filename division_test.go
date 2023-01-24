package main

import "testing"

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
