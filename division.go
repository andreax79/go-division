package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Get the number of digits of a given number
func NumberOfDigits(n int) int {
	if n < 10 {
		return 1
	} else {
		return NumberOfDigits(n/10) + 1
	}
}

func spaces(n int) string {
	return strings.Repeat(" ", n)
}

func GetIntArg(args []string, n int) int {
	num, err := strconv.Atoi(args[n])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	return num
}

// Parse command line arguments
func ParseArgs(args []string) (dividend, divisor int) {
	if len(args) != 3 {
		fmt.Fprintf(os.Stderr, "usage: division <dividend> <divisor>\n")
		os.Exit(2)
	}
	dividend = GetIntArg(args, 1)
	divisor = GetIntArg(args, 2)
	if divisor == 0 {
		fmt.Fprintf(os.Stderr, "division by zero\n")
		os.Exit(1)
	}
	return dividend, divisor
}

func print(left int, right int, digits int, i int) {
	t := ""
	if right >= 0 {
		t = fmt.Sprint(right)
	}
	fmt.Println(spaces(i-NumberOfDigits(left)+1), left, spaces(digits-i-1), "|", t)
}

type Step struct {
	left  int
	right int
	i     int
}

func main() {
	dividend, divisor := ParseArgs(os.Args)

	digits := NumberOfDigits(dividend)
	result := 0
	num := 0
	steps := make([]Step, 0)

	steps = append(steps, Step{dividend, divisor, digits - 1})

	for i, ch := range fmt.Sprint(dividend) {
		n := int(ch - '0')
		num = num*10 + n

		if num >= divisor {
			steps = append(steps, Step{num, -1, i})
			d := 0
			for num >= divisor {
				num = num - divisor
				d++
			}
			result = result*10 + d
			steps = append(steps, Step{d * divisor, -1, i})
		}
	}
	steps = append(steps, Step{num, -1, digits - 1})
	steps[1].right = result

	for _, step := range steps {
		print(step.left, step.right, digits, step.i)
	}

	fmt.Println()
	fmt.Println("Result:", result)
	fmt.Println("Remainder:", num)

}
