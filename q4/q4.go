package q4
import ( "errors")
func CalculateFinalPrice(basePrice float64, state string, taxCode int) (float64, error) {
	if basePrice <= 0 {
		return 0, errors.New("preço invalido")
	}
	var taxa float64
	switch taxCode {
	case 1:
		taxa = 0.05
	case 2:
		taxa = 0.10
	case 3:
		taxa = 0.15
	default:
		return 0, errors.New("imposto não encontrado")
	}
	var frete float64

	if state == "SP" {
		frete = 0.1
	} else if state == "RJ" {
		frete = 0.15
	} else if state == "MG" {
		frete = 0.2
	} else if state == "ES" {
		frete = 0.25
	} else {
		frete = 0.3
	}

	precototal := basePrice + (basePrice * taxa) + (basePrice * frete)
	return precototal, nil

}
