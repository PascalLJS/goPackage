package atelier3

import (
	"math/rand"
)

// salut
func ArrayGenerate(nbElem int) []int {
	var array []int

	for i := 0; i < nbElem; i++ {
		rand := rand.Intn(10000)
		array = append(array, rand)
	}
	return array
}
