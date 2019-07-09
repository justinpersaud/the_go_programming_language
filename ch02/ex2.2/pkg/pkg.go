package pkg

import "fmt"

type Meters float64
type Feet float64
type Pounds float64
type Kilograms float64
type Celsius float64
type Fahrenheit float64

func (c Celsius) String() string {
	return fmt.Sprintf("%gC", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%gF", f)
}

func (m Meters) String() string {
	return fmt.Sprintf("%gm", m)
}

func (ft Feet) String() string {
	return fmt.Sprintf("%gf", ft)
}

func (p Pounds) String() string {
	return fmt.Sprintf("%glbs", p)
}

func (kg Kilograms) String() string {
	return fmt.Sprintf("%gkg", kg)
}
