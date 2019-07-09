// General purpose program that will convert numbers passed in as arguments to their temperature in Celsius and Fahrenheit, length in feet and meters, and weight in pounds and kilograms.

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"the_go_programming_language/ch02/ex2.2/pkg"
)

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			num, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				panic(err)
			}
			c := pkg.Celsius(num)
			f := pkg.Fahrenheit(num)
			m := pkg.Meters(num)
			ft := pkg.Feet(num)
			lbs := pkg.Pounds(num)
			kg := pkg.Kilograms(num)

			fmt.Printf("%s = %s\n", c, pkg.CtoF(c))
			fmt.Printf("%s = %s\n", f, pkg.FtoC(f))
			fmt.Printf("%s = %s\n", m, pkg.MeterstoFeet(m))
			fmt.Printf("%s = %s\n", ft, pkg.FeettoMeters(ft))
			fmt.Printf("%s = %s\n", lbs, pkg.PoundstoKilograms(lbs))
			fmt.Printf("%s = %s\n", kg, pkg.KilogramstoPounds(kg))
		}
	} else {
		log.Fatal("Please provide arguments")
	}
}
