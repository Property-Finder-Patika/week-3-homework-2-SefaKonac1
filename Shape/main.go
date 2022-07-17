package main

import (
	"fmt"
	rect "shape/shapes"
)

func main() {

	/*It's recovers to error handling*/
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	/*It creates new rectangular objects and produces outputs rectangular's area and circumstances*/
	r1 := rect.NewRectangular(15.5, 6.2)

	fmt.Println("area : %f", r1.CalculateArea())
	fmt.Println("circumstances : %f", r1.CalculateCircumstances())

	/*Generates an error, and handling it according to 6 strategy*/
	r2 := rect.NewRectangular(-5, 6)

	fmt.Println("area : %f", r2.CalculateArea())
	fmt.Println("circumstances : %f", r2.CalculateCircumstances())

}
