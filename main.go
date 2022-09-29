// Write a console application in Go that takes an integer argument
// and reports to console all prime numbers between zero and the
// given argument.
// Make sure your code is clean, your solution scales and is optimized on efficiency.

package main

import (
	"fmt"
	"os"
	"strconv"
)

// Write a console application in Go that takes an integer argument
// and reports to console all prime numbers between zero and the
// given argument.

func main() { // main is the entry point for the program
	if len(os.Args) != 2 { // if the number of command line arguments is not 2
		fmt.Println("Please provide a number") // print message to stdout
		return                                 // exit
	}
	n, err := strconv.Atoi(os.Args[1]) // convert the command line argument to an integer
	if err != nil {                    // if there is an error
		fmt.Println("Please provide a valid number") // print message to stdout
		return                                       // exit the program
	}
	if n < 0 { // if the number is negative
		fmt.Println("Please provide a positive number") // print message to stdout
		return                                          // exit the program
	}
	for i := 2; i <= n; i++ { // for each number from 2 to n
		if isPrime(i) { // if the number is prime (call the isPrime function)
			fmt.Println(i) // print the number to stdout (it is prime)
		}
	}
}

func isPrime(n int) bool { // isPrime returns true if n is prime
	for i := 2; i < n; i++ { // for each number from 2 to n
		if n%i == 0 { // if the number is divisible by i
			return false // return false (it is not prime)
		} // end if
	}
	return true // return true (it is prime)
}
