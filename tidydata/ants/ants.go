package ants

import (
	"fmt"
	"strconv"
)

// Récupère le nombre de fourmis.
func TidyAnts(stringTable []string) int {
	if len(stringTable[0]) == 0 {
		fmt.Println("Absence de fourmis")
		return 0
	}

	ant, err := strconv.Atoi(stringTable[0])

	if err != nil {
		fmt.Println("Erreur sur le nombre de fourmis")
		return 0
	}

	if ant <= 0 {
		fmt.Println("Nombre de fourmis inférieur à 0")
		return 0
	}

	return ant
}
