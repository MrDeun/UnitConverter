package convert

import (
	"errors"
	"strconv"
)

type Tempeture string

const (
	kelvin     Tempeture = "kelvin"
	celcius    Tempeture = "celcius"
	fahrenheit Tempeture = "fahrenheit"
)

func ConvertTemp(args []string) (float64, error) {
	switch args[1] {
	case "kelvin", "celcius", "fahrenheit":
		quantity, err := strconv.ParseFloat(args[2], 64)
		if err != nil {
			return 0.0, errors.New("Error: failed to parse a quantity from string to number")
		} else {
			return chooseConversion(args[0], args[1], quantity), nil
		}
	default:
		return 0.0, errors.New("Error: tried unrelated type to the tempeture!")
	}
}

func chooseConversion(from, to string, quantity float64) float64 {
	switch from {
	case "kelvin":
		switch to {
		case "fahrenheit":
			return kelvinToFahrenheit(quantity)
		case "celcius":
			return kelvinToCelcius(quantity)
		case "kelvin":
			return quantity
		}
	case "fahrenheit":
		switch to {
		case "fahrenheit":
			return quantity
		case "celcius":
			return fahrenheitToCelcius(quantity)
		case "kelvin":
			return fahrenheitToKelvin(quantity)
		}
	case "celcius":
		switch to {
		case "fahrenheit":
			return celciusToFahrenheit(quantity)
		case "celcius":
			return quantity
		case "kelvin":
			return celciusToKelvin(quantity)
		}
	default:
		panic("That's not supposed to happen...")
	}
}

func celciusToKelvin(cel float64) float64 {
	panic("Lack of implementation!")
}

func celciusToFahrenheit(cel float64) float64 {
	panic("Lack of implementation!")
}

func fahrenheitToCelcius(fahr float64) float64 {
	panic("Lack of implementation!")
}

func fahrenheitToKelvin(fahr float64) float64 {
	panic("Lack of implementation!")
}

func kelvinToFahrenheit(kel float64) float64 {
	panic("Lack of implementation!")
}

func kelvinToCelcius(kel float64) float64 {
	panic("Lack of implementation!")

}
