// Copyright 2023 Andrea Bonomi - andrea.bonomi@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package division

import (
	"fmt"
	"strings"
)

const empty = -1
const separator = -2

// Returns the largest of a or b
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func spaces(n int) string {
	if n > 0 {
		return strings.Repeat(" ", n)
	} else {
		return ""
	}
}

// Get the number of digits of a given number
func NumberOfDigits(n int) int {
	if n < 10 {
		return 1
	} else {
		return NumberOfDigits(n/10) + 1
	}
}

type Step struct {
	left  int
	right int
	i     int
	step  int
}

// Format a step for printing
func (step *Step) FormatStep(division *Division) string {
	// Left part
	left := ""
	switch step.left {
	case empty:
		left = spaces(division.dividendDigits + 2)
	case separator:
		length := NumberOfDigits(step.i)
		preSpaces := spaces(step.i - length + 2)
		left = preSpaces + strings.Repeat("-", division.dividendDigits-step.i+1)
	default:
		length := NumberOfDigits(step.left)
		preSpaces := spaces(step.i - length + 2)
		postSpaces := spaces(division.dividendDigits - step.i)
		left = fmt.Sprintf("%s%d%s", preSpaces, step.left, postSpaces)
	}
	// Right part
	right := ""
	switch step.right {
	case empty:
		right = ""
	case separator:
		length := max(division.dividendDigits, NumberOfDigits(division.Result))
		right = strings.Repeat("-", length)
	default:
		right = fmt.Sprintf(" %d", step.right)
	}
	return left + "|" + right
}

type Division struct {
	Steps          []Step
	Dividend       int
	Divisor        int
	Result         int
	Remainder      int
	dividendDigits int
	divisorDigits  int
}

// Create a new Division instance
func NewDivision(dividend, divisor int) *Division {
	division := new(Division)
	division.Steps = make([]Step, 0)
	division.Dividend = dividend
	division.Divisor = divisor
	division.dividendDigits = NumberOfDigits(dividend)
	division.divisorDigits = NumberOfDigits(divisor)
	return division
}

// Calculate the division
func (division *Division) Calculate() {
	result := 0
	remainder := 0
	step := 0
	division.addStep(division.Dividend, division.Divisor, division.dividendDigits-1, 0)
	for i, ch := range fmt.Sprint(division.Dividend) {
		n := int(ch - '0')
		remainder = remainder*10 + n

		d := 0
		if remainder >= division.Divisor {
			if step > 0 {
				division.addStep(remainder, empty, i, step)
			}
			for remainder >= division.Divisor {
				remainder = remainder - division.Divisor
				d++
			}
		}

		result = result*10 + d
		division.addStep(d*division.Divisor, empty, i, step)
		division.addStep(separator, empty, i, step)
		step++
	}
	division.addStep(remainder, empty, division.dividendDigits-1, step)
	division.setResult(result, remainder)
}

// Add a step
func (division *Division) addStep(left int, right int, i int, step int) {
	division.Steps = append(division.Steps, Step{left, right, i, step})
}

// Set result and reminder
func (division *Division) setResult(result, remainder int) {
	division.Result = result
	division.Remainder = remainder
	division.Steps[1].right = separator
	if len(division.Steps) < 3 {
		division.addStep(empty, result, division.dividendDigits-1, 1)
	} else {
		division.Steps[2].right = result
	}
}

// Format the division
func (division *Division) String() string {
	var sb strings.Builder
	for _, step := range division.Steps {
		sb.WriteString(step.FormatStep(division))
		sb.WriteString("\n")
	}
	sb.WriteString(fmt.Sprintf("\n%d : %d = %d (%d)\n", division.Dividend, division.Divisor, division.Result, division.Remainder))
	return sb.String()
}

// Print the division
func (division *Division) Print() {
	fmt.Println(division)
}
