package parse

import "testing"

func TestGetIntArg(t *testing.T) {
	args := []string{"test", "10", "20"}
	GetIntArg(args, 1)
	GetIntArg(args, 2)
}

func TestParseArgs(t *testing.T) {
	args := []string{"test", "9876", "543"}
	dividend, divisor := ParseArgs(args, "")
	if dividend != 9876 {
		t.Fatalf("dividend != 9876")
	}
	if divisor != 543 {
		t.Fatalf("divisor != 543")
	}
}
