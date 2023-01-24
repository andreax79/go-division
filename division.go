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

// Get an integer argument from the line arguments
func GetIntArg(args []string, n int) int {
	num, err := strconv.Atoi(args[n])
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid number\n")
		os.Exit(1)
	}
	if num < 0 {
		fmt.Fprintf(os.Stderr, "the argument must be positive\n")
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

type Step struct {
	left  int
	right int
	i     int
}

func spaces(n int) string {
	return strings.Repeat(" ", n)
}

// Print a step
func (step *Step) PrintStep(digits int) {
	right := ""
	if step.right >= 0 {
		right = fmt.Sprint(step.right)
	}
	fmt.Println(spaces(step.i-NumberOfDigits(step.left)+1), step.left, spaces(digits-step.i-1), "| ", right)
}

type Steps struct {
	steps  []Step
	digits int
}

// Add a step
func (steps *Steps) AddStep(left int, right int, i int) {
	steps.steps = append(steps.steps, Step{left, right, i})
}

// Set result
func (steps *Steps) SetResult(result int) {
	steps.steps[1].right = result
}

// Print steps
func (steps *Steps) Print() {
	for _, step := range steps.steps {
		step.PrintStep(steps.digits)
	}
}

func main() {
	dividend, divisor := ParseArgs(os.Args)

	digits := NumberOfDigits(dividend)
	result := 0
	num := 0
	steps := Steps{make([]Step, 0), digits}

	steps.AddStep(dividend, divisor, digits-1)

	for i, ch := range fmt.Sprint(dividend) {
		n := int(ch - '0')
		num = num*10 + n

		if num >= divisor {
			steps.AddStep(num, -1, i)
			d := 0
			for num >= divisor {
				num = num - divisor
				d++
			}
			result = result*10 + d
			steps.AddStep(d*divisor, -1, i)
		}
	}
	steps.AddStep(num, -1, digits-1)
	steps.SetResult(result)
	steps.Print()

	fmt.Println()
	fmt.Println("Result:", result)
	fmt.Println("Remainder:", num)
}
