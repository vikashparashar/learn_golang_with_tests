package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	var firstnumber, secondnumber float64
	var choice int
	fmt.Printf("\nEnter First Number       :  ")
	fmt.Scanf("%f\n", &firstnumber)
	fmt.Printf("Enter Second Number      :  ")
	fmt.Scanf("%f\n", &secondnumber)
	fmt.Println("Enter Your Choice From Options Mentioned Below: \n i.e: for Addition Type 1 And Then Press Enter\n\n\n ")
	fmt.Printf("1. Addition  \n2. Subtraction  \n3. Multiplication  \n4. Division  \n5. Remainder  \n")
	fmt.Scanf("%d", &choice)

	switch choice {
	case 1:
		res := Addition(firstnumber, secondnumber)
		fmt.Printf("Addition Of %f And %f Is : %f\n", firstnumber, secondnumber, res)
	case 2:
		res := Subtraction(firstnumber, secondnumber)
		fmt.Printf("Subtraction Of %f And %f Is : %f\n", firstnumber, secondnumber, res)
	case 3:
		res := Multiplication(firstnumber, secondnumber)
		fmt.Printf("Multiplication Of %f And %f Is : %f\n", firstnumber, secondnumber, res)
	case 4:
		res, err := Division(firstnumber, secondnumber)
		if err != nil {
			return
		}
		fmt.Printf("Division Of %f And %f Is : %f\n", firstnumber, secondnumber, res)

	default:
		os.Exit(1)

	}
}

func Addition(x, y float64) float64 {
	result := x + y
	return result

}

func Subtraction(x, y float64) float64 {
	result := x - y
	return result
}

func Multiplication(x, y float64) float64 {
	result := x * y
	return result
}

func Division(x, y float64) (float64, error) {
	result := x / y
	if y == 0 {
		return result, errors.New("can not devide by zero")
	}
	return result, nil
}
