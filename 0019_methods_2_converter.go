package main

import "fmt"

type celsius float64
type kelvin float64
type fahrenheit float64

func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit(c*9.0/5.0 + 32)
}

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

func (k kelvin) fahrenheit() fahrenheit {
	return fahrenheit((k-273.15)*9.0/5.0 + 32)
}

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32) * 5.0 / 9.0)
}

func (f fahrenheit) kelvin() kelvin {
	return kelvin((f-32)*5.0/9.0 + 273.15)
}

func main() {
	var k kelvin = 294.0
	c := k.celsius()

	fmt.Println(k, " K is ", c, " C")
}
