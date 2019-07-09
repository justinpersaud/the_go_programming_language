// Ftoc prints two Farenheit-to-Celsius conversions.package main

package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g degrees F = %g degrees C\n", freezingF, fToC(freezingF)) // "32F = 0C"
	fmt.Printf("%g degrees F = %g degrees C\n", boilingF, fToC(boilingF))   // "212F = 100C"
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
