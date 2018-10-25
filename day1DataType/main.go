package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)
	// Declare second integer, double, and String variables.
	var resultInt uint64
	var resultDouble float64
	var resultString string
	// Read and save an integer, double, and String to your variables.
	fmt.Scan(&resultInt)
	fmt.Scan(&resultDouble)
	// Print the sum of both integer variables on a new line.
	fmt.Printf("%d\n", resultInt+i)
	fmt.Printf("%.1f\n", resultDouble+d)
	// Print the sum of the double variables on a new line.
	for scanner.Scan() {
		resultString = scanner.Text()
		fmt.Printf("%s%s\n", s, resultString)
	}
	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.
}
