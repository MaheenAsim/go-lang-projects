package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func prob_2(s string) bool {
	// for simple words:
	/*for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true*/

	// For sentences: Remove spaces and convert to lower case
	processedString := strings.ReplaceAll(s, " ", "")
	processedString = strings.ToLower(processedString) //ask: should we add this line to convert to lower of no? since Aa would then be true

	// Check for palindrome
	for i := 0; i < len(processedString)/2; i++ {
		if processedString[i] != processedString[len(processedString)-1-i] {
			return false
		}
	}
	return true
}


func main() {
  fmt.Println("\nEnter a string to check if it's a palindrome:")

	// Use bufio.Scanner to read the whole line
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	if prob_2(input) {
		fmt.Println("\nThe sentence is a palindrome.\n")
	} else {
		fmt.Println("\nThe sentence is not a palindrome.\n")
	}


  
}
