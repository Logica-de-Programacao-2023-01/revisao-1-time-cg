package bonus
import (
	"fmt"
)
const (
	dano          = 7
	maxdifference = 20
)

func CalculateDamage(characterLevel, enemyLevel int) (int, error) {
	if characterLevel <= 0 || enemyLevel <= 0 {
		return 0, fmt.Errorf("nivel invalido")
	}
	dano := characterLevel * dano

	diferençanivel := characterLevel - enemyLevel
	if diferençanivel > maxdifference {
		dano *= 5
	} else if diferençanivel < -maxdifference {
		dano *= 2
	}
	if characterLevel > enemyLevel {
		dano *= 10
	} else if characterLevel < enemyLevel {
		dano *= 5
	}
	if characterLevel > 100 {
		dano *= 20
	} else if enemyLevel > 100 {
		dano *= 2
	}
	return dano, nil
}
