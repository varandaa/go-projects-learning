package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var args []string = os.Args[1:]
	result, err := calculate(args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("The result is: %v\n", result)
}

func calculate(args []string) (int, error) {
	var err error
	if len(args) != 3 {
		err = errors.New("Maximum args = 3")
		return -1, err
	}
	a, err := strconv.Atoi(args[0])
	if err != nil {
		return -1, err
	}
	b, err := strconv.Atoi(args[2])
	if err != nil {
		return -1, err
	}
	var sign string = args[1]
	switch sign {
	case "+":
		return a + b, err
	case "-":
		return a - b, err
	case "calc.go":
		return a * b, err
	case "/":
		return a / b, err
	}
	err = errors.New("Sign is not: + | - | * | / ")
	return -1, err
}
