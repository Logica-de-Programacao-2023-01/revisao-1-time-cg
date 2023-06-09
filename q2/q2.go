package q2
import ("errors" 
	"strings"
)
func AverageLettersPerWord(text string) (float64, error) {
	simbolos := []string{",", ".", "?", "!", ";", ":"}
	list := strings.Fields(text)
	soma := 0
	if text == "" {
		return 0, errors.New("textp vazio")
	}
	for i := 0; i < len(simbolos); i++ {
		text = strings.ReplaceAll(text, simbolos[i], "")
	}

	for i := 0; i < len(list); i++ {
		soma += len(list[i])
	}
	return float64(soma) / float64(len(list)), nil
}
