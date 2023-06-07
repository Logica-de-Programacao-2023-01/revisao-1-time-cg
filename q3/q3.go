package q3

	
import (
	"errors"
)

func FindMinMaxAverage(numbers []int) (int, int, float64, error) {
	if len(numbers) == 0 {
		return 0, 0, 0, errors.New("lista vazia")
	}
	maximo := numbers[0]
	min := numbers[0]
	soma := 0

	for _, numero := range numbers {
		if numero > maximo {
			maximo = numero
		}
		if numero < min {
			min = numero
		}
		soma += numero
	}
	average := float64(soma-(maximo+min)) / float64(len(numbers)-2)

	return maximo, min, average, nil
}

}
