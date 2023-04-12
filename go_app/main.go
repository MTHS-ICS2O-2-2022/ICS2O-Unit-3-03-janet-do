package main

import "fmt"
import "math"

func main() {
	// This function finds the volume of a sphere
	var radius float64
	var volume float64

	// input
	fmt.Println("This program finds the volume of a sphere.")
	fmt.Println()
	fmt.Print("Enter the radius (in cm): ")
	fmt.Scanln(&radius)
	fmt.Println()

	// process
	volume = (4.0 / 3.0) * math.Pi * math.Pow(radius, 3)

	// output
	fmt.Printf("The volume of the sphere is: %.2f cm³\n", volume)
	fmt.Println()
	fmt.Println("Done.")
}
