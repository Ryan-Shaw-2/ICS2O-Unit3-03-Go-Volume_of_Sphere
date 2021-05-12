// Created by: Ryan-Shaw-2
// Created on: May 2021
//
// This program calculates the volume of a sphere

package main

import (
	"fmt"
	"math"
)

func main() {
	// This function calculates the volume of a sphere
	var radius float64

	// input
	fmt.Println("This program calculates the volume of a sphere")
	fmt.Println()
	fmt.Print("Enter in the radius (cm): ")
	fmt.Scanln(&radius)
	fmt.Println()

	// process
	var volume = float64(4.0/3.0) * math.Pi * math.Pow(radius, 3)

	// output
	fmt.Println("The volume is:", math.Round(volume*1000)/1000, "cm³")
}
