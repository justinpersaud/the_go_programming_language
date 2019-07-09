package tempconv

// CToF converts a Celsius temperature to Fahrenheit.

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius(f-32) * 5 / 9
}

// fmt.Println("Brrr! %v\n", tempconv.AbsoluteZeroC) // "Brrr! -273.15 degrees Celsius"
