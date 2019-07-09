package pkg

// Conversion functions

func MeterstoFeet(m Meters) Feet {
	return Feet(m * 3.28084)
}

func FeettoMeters(f Feet) Meters {
	return Meters(f * 0.3048)
}

func PoundstoKilograms(p Pounds) Kilograms {
	return Kilograms(p * 0.453592)
}

func KilogramstoPounds(kg Kilograms) Pounds {
	return Pounds(kg * 2.20462)
}

func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FtoC(f Fahrenheit) Celsius {
	return Celsius(f-32) * 5 / 9
}
