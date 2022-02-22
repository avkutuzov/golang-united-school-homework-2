package main

import (
	"fmt"
	square "square/Square"
)

func main() {
	fmt.Println(square.CalcSquare(10.0, "SidesTriangle"))
	fmt.Println(square.CalcSquare(10.0, "SidesSquare"))
	fmt.Println(square.CalcSquare(10.0, "SidesCircle"))
	fmt.Println(square.CalcSquare(10.0, "Sides"))
}