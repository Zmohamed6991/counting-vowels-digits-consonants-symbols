package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*
	Write a Go program that takes an input string and
	counts the number of vowels, consonants, digits, and special characters in the string.

*/

	func symbols(userInput string) int {

		count := 0 

		for i := 0; i < len(userInput); i++ {
			symbols := rune(userInput[i])
		 		if !unicode.IsDigit(symbols) && !unicode.IsLetter(symbols){
					count++
				}
			}

		return count

	}

	func digits(userInput string) int {

		count := 0

		for i := 0; i < len(userInput); i++ {
			if unicode.IsDigit(rune(userInput[i])) {
				count++
			}
		}

		return count
	}

	func consonants(userInput string) int {

		count := 0
		consonants := "bcdfghjklmnpqrstvwxyzBCDFGHJKLMNPQRSTVWXYZ"

		for i := 0; i < len(userInput); i++ {
			if strings.ContainsRune(consonants, rune(userInput[i])){
				count++
			}
		}
		 
		return count
	}

	func vowels(userInput string) int {
	
		var vowels = "AEOIUaeoiu"
		count := 0

		for i := 0; i < len(userInput); i++ {
			if strings.ContainsRune(vowels, rune(userInput[i])) {
				count++
			}
		} 
		return count
				
	}

	
	func main() {

		for {
			
			var userInput string 
			fmt.Println("Enter a word with capital letters and lower letters, with numbers and a symbol: ")
			fmt.Scan(&userInput)

			vowels := vowels(userInput)
			consonants := consonants(userInput)
			digits:= digits(userInput)
			symbols := symbols(userInput)

			fmt.Println("The number of vowels used is: ", vowels)
			fmt.Println("The number of consonants is: ", consonants)
			fmt.Println("The number of digits used is: ", digits)
			fmt.Println("The number of symobols used is:", symbols)
		}
		
	}