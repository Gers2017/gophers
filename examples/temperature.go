package examples

import "fmt"

type Celsius float32
type Fahrenheit float32


func Temperature() {
	v := ToFahrenheit(100) // should be 212
	fmt.Printf("%f should be 212, %v", v, v == 212)
}

func ToFahrenheit(celsius Celsius) Fahrenheit {
	return Fahrenheit((celsius * 9 / 5) + 32)
}