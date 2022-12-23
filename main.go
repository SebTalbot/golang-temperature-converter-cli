package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	originUnit  string
	originValue float64
)

var shouldConvertAgain string

var err error

var (
	errInvalidArguments = errors.New("Invalid arguments")
	errReadingInput     = errors.New("Error reading input")
)

func main() {
	if len(os.Args) != 2 {
		printError(errInvalidArguments)
	}

	originUnit := strings.ToUpper(os.Args[1])

	if originUnit != "C" && originUnit != "F" {
		printError(errInvalidArguments)
	}

	for {
		fmt.Print("What is the current temperature in " + originUnit + " ? ")
		_, err = fmt.Scanln(&originValue)
		if err != nil {
			printError(errReadingInput)
		}

		if originUnit == "C" {
			convertToFahrenheit(originValue)
		}

		if originUnit == "F" {
			convertToCelsius(originValue)
		}

		fmt.Print("Would you like to convert another temperature ? (y/n) ")
		_, err = fmt.Scanln(&shouldConvertAgain)
		if err != nil {
			printError(errReadingInput)
		}

		if strings.ToUpper(shouldConvertAgain) != "Y" {
			fmt.Println("Good bye!")
			break
		}
	}
}

func printError(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}

func convertToCelsius(value float64) {
	convertedValue := (value - 32) * 5 / 9
	fmt.Printf("%v F = %.0f C\n", value, convertedValue)
}

func convertToFahrenheit(value float64) {
	convertedValue := (value * 9 / 5) + 32
	fmt.Printf("%v C = %.0f F\n", value, convertedValue)
}
