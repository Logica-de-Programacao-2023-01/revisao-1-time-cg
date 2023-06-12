package q1

import "errors"

func CalculateDiscount(currentPurchase float64, purchaseHistory []float64) (float64, error) {
	if currentPurchase <= 0 {
		return 0, errors.New("valor invalido")
	}

	totalPurchase := 0.0
	for _, purchase := range purchaseHistory {
		totalPurchase += purchase
	}

	averagePurchase := totalPurchase / float64(len(purchaseHistory))

	discount := 0.0

	if len(purchaseHistory) == 0 {
		discount = currentPurchase * 0.1
	} else if totalPurchase > 1000 {
		discount = currentPurchase * 0.1
	} else if totalPurchase > 500 {
		discount = currentPurchase * 0.05
	} else if totalPurchase <= 500 {
		discount = currentPurchase * 0.02
	}
	if averagePurchase > 1000 {
		discount = currentPurchase * 0.2
	}
	if discount == 0 {
		return 0, nil
	}
	return discount, nil
}

