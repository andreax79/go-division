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

package parse

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

const usage = `Usage:
    division <DIVIDEND> <DIVISOR>

Example:
    division 3279 25
`

// Print error and exit
func error(format string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, "division: "+format+"\n", v...)
	os.Exit(1)
}

// Get an integer argument from the line arguments
func GetIntArg(args []string, n int) int {
	num, err := strconv.Atoi(args[n])
	if err != nil {
		error("invalid number")
	}
	if num < 0 {
		error("the argument must be positive")
	}
	return num
}

// Parse command line arguments
func ParseArgs(args []string, version string) (dividend, divisor int) {
	flag.Usage = func() { fmt.Fprintf(os.Stderr, "%s\n", usage) }

	var versionFlag bool
	flag.BoolVar(&versionFlag, "version", false, "print the version")
	flag.CommandLine.Parse(args[1:])
	if versionFlag {
		if version != "" {
			fmt.Println("Version", version)
		}
		os.Exit(0)
	}

	if len(flag.Args()) != 2 {
		flag.Usage()
		os.Exit(2)
	}

	dividend = GetIntArg(flag.Args(), 0)
	divisor = GetIntArg(flag.Args(), 1)
	if divisor == 0 {
		error("division by zero")
	}
	return dividend, divisor
}
