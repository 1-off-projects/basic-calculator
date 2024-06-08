package main

import (
	"fmt"
)

func main() {
	var choice string
	fmt.Println("menu")
	fmt.Println("1. add")
	fmt.Println("2. subtract")
	fmt.Println("3. multiply")
	fmt.Println("4. devide")
	fmt.Println("enter options 1-4")
	_, err := fmt.Scanln(&choice)
	if err != nil {
		fmt.Println("error")
	}
	if choice == "1" {
		var num1, num2 float64
		fmt.Println("enter 1st number")
		_, err = fmt.Scanln(&num1)
		if err != nil {
			fmt.Println("error")
		}
		fmt.Println("enter 2nd number")
		_, err = fmt.Scanln(&num2)
		if err != nil {
			fmt.Println("error")
		}
		fmt.Println("sum is:", (num1 + num2))
	} else if choice == "2" {
		var num1, num2 float64
		fmt.Println("enter 1st number")
		_, err = fmt.Scanln(&num1)
		if err != nil {
			fmt.Println("error")
		}
		fmt.Println("enter 2nd number")
		_, err = fmt.Scanln(&num2)
		if err != nil {
			fmt.Println("error")
		}
		fmt.Println("differance is", (num1 - num2))
	} else if choice == "3" {
		var num1, num2 float64
		fmt.Println("enter 1st number")
		_, err = fmt.Scanln(&num1)
		if err != nil {
			fmt.Println("error")
		}
		fmt.Println("enter 2nd number")
		_, err = fmt.Scanln(&num2)
		if err != nil {
			fmt.Println("error")
		}
		fmt.Println("product is:", (num1 * num2))
	} else if choice == "4" {
		var num1, num2 float64
		fmt.Println("enter 1st number")
		_, err = fmt.Scanln(&num1)
		if err != nil {
			fmt.Println("error")
		}
		fmt.Println("enter 2nd number")
		_, err = fmt.Scanln(&num2)
		if err != nil {
			fmt.Println("error")
		}
		fmt.Println("quotient is:", (num1 / num2))
	}
}
