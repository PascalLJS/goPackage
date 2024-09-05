package atelier3

import (
	"math/rand"
)

func arrayGenerate(nbElem int) []int {
	var array []int

	for i := 0; i < nbElem; i++ {
		rand := rand.Intn(10000)
		array = append(array, rand)
	}
	return array
}
