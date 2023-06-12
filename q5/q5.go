package q5

import(
	"errors"
	"fmt"
)

func ConvertTemperature(temp float64, fromScale string, toScale string) (float64, error) {
	if fromScale != "C" && fromScale != "F" && fromScale != "K" {
		return 0, fmt.Errorf("escala invalida")
	}
	if toScale != "C" && toScale != "F" && toScale != "K" {
		return 0, fmt.Errorf("escala invalida")
	}
	tempconver := 0.0

	switch fromScale {
	case "C":
		switch toScale {
		case "F":
			tempconver = (temp * 9 / 5) + 32
		case "K":
			tempconver = temp + 273.15
		}
	case "F":
		switch toScale {
		case "C":
			tempconver = (temp - 32) * 5 / 9
		case "K":
			tempconver = (temp + 459.67) * 5 / 9

		}
	case "K":
		switch toScale {
		case "C":
			tempconver = temp - 273.15
		case "F":
			tempconver = (temp * 9 / 5) - 459.67
		}
	default:
		return 0, errors.New("escala invalida")
	}
	return tempconver, nil
}

