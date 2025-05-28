package main

import (
	"bufio"    // For reading input line by line
	"errors"   // For creating custom error messages
	"fmt"      // For formatted I/O (printing and reading)
	"os"       // For standard input/output streams (os.Stdin)
	"strconv"  // For string to number conversion
	"strings"  // For string manipulation (trimming whitespace)
)

// main is the entry point of the program.
func main() {
	fmt.Println("--- Simple Command-Line Calculator ---")

	// Read the first number
	num1, err := readNumber("Enter the first number: ")
	if err != nil {
		fmt.Println("Error:", err) // Display user-friendly error
		os.Exit(1)                 // Exit gracefully on error
	}

	// Read the operator
	operator, err := readOperator("Enter the operator (+, -, *, /): ")
	if err != nil {
		fmt.Println("Error:", err) // Display user-friendly error
		os.Exit(1)
	}

	// Read the second number
	num2, err := readNumber("Enter the second number: ")
	if err != nil {
		fmt.Println("Error:", err) // Display user-friendly error
		os.Exit(1)
	}

	// Perform the calculation
	result, err := calculateResult(num1, num2, operator)
	if err != nil {
		fmt.Println("Error:", err) // Display specific calculation error
		os.Exit(1)
	}

	// Display the result
	fmt.Printf("Result: %.2f %s %.2f = %.2f\n", num1, operator, num2, result)
}

// readInput reads a line of text from standard input, trims whitespace,
// and returns it. It also returns an error if reading fails.
func readInput(prompt string) (string, error) {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}
	return strings.TrimSpace(input), nil
}

// readNumber prompts the user for a number, reads their input,
// and attempts to convert it to a float64. It returns the number
// and an error if conversion fails.
func readNumber(prompt string) (float64, error) {
	inputStr, err := readInput(prompt)
	if err != nil {
		return 0, err // Pass through the error from readInput
	}

	num, err := strconv.ParseFloat(inputStr, 64)
	if err != nil {
		// Use errors.New for a clean, user-friendly error message
		return 0, errors.New("Invalid number. Please enter a numeric value.")
	}
	return num, nil
}

// readOperator prompts the user for an operator, reads their input,
// and validates if it's one of the supported arithmetic operators.
// It returns the valid operator or an error.
func readOperator(prompt string) (string, error) {
	op, err := readInput(prompt)
	if err != nil {
		return "", err // Pass through the error from readInput
	}

	switch op {
	case "+", "-", "*", "/":
		return op, nil
	default:
		// Use errors.New for a clean, user-friendly error message
		return "", errors.New("Invalid operator. Please use +, -, *, or /.")
	}
}

// calculateResult performs the arithmetic operation based on the
// provided numbers and operator. It handles division by zero.
// It returns the result and an error if the operation is invalid
// (e.g., division by zero).
func calculateResult(num1, num2 float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			// Specific error for division by zero
			return 0, errors.New("Division by zero is not allowed.")
		}
		return num1 / num2, nil
	default:
		// This case should ideally not be reached if readOperator works correctly,
		// but it acts as a safeguard.
		return 0, fmt.Errorf("unexpected operator: %s", operator)
	}
}