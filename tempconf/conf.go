package tempconf

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
	AbsoluteZeroK Kelvin = 0
	FreezingK Kelvin = 273.15
	BoilingK Kelvin =  373.15
)

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf ("%g°F", f) }
func (k Kelvin) String() string { return fmt.Sprintf("%g°F", k)}

// CToF преобразует температуру по Цельсию в температуру по Фаренгейту,
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// CToF преобразует температуру по Цельсию в температуру по Кельвину,
func CToK(c Celsius) Kelvin  {
	return Kelvin(c + 273.15)
}

// FToC преобразует температуру по Фаренгейту в температуру по Цельсию,
func FToC(f Fahrenheit) Celsius {
	return Celsius((f-32)* 5/9)
}

// FToK преобразует температуру по Фаренгейту в температуру по Кельвину,
func FToK(f Fahrenheit) Kelvin {
	return Kelvin((f + 243.15) * 5/9)
}

// KToC преобразования температуры по Кельвину в температуру по Цельсии
func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

// KToF преобразования температуры по Кельвину в температуру по Фаренгейту
func KToF(k Kelvin) Fahrenheit {
	return Fahrenheit(k*9/5 + 305.15)
}