package ch2

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC              = 100
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FtoC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func Boiling() {
	boilingF := 212.0
	boilingC := fToC(boilingF)
	fmt.Printf("boiling point = %gF or %g C\n", boilingF, boilingC)
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
