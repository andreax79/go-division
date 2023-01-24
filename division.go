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

func print(num int, digits int, i int) {
	fmt.Println(spaces(i-NumberOfDigits(num)+1), num, spaces(digits-i-1), "|")
}

func main() {
	dividend, divisor := ParseArgs(os.Args)
	fmt.Println("", dividend, " |", divisor)

	digits := NumberOfDigits(dividend)
	result := 0
	num := 0
	for i, ch := range fmt.Sprint(dividend) {
		n := int(ch - '0')
		num = num*10 + n

		if num >= divisor {
			print(num, digits, i)
			d := 0
			for num >= divisor {
				num = num - divisor
				d++
			}
			result = result*10 + d
			print(d*divisor, digits, i)
		}
	}
	print(num, digits, digits-1)
	fmt.Println(spaces(NumberOfDigits(dividend)-NumberOfDigits(num)), num, " |")

	fmt.Println()
	fmt.Println("Result:", result)
	fmt.Println("Remainder:", num)

}
