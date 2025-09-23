package ants

import (
	"fmt"
	"strconv"
)

// Récupère le nombre de fourmis.
func TidyAnts(stringTable []string) int {
	if len(stringTable[0]) == 0 {
		fmt.Println("Error, No ants")
		return 0
	}

	ant, err := strconv.Atoi(stringTable[0])

	if err != nil {
		fmt.Println("Error, false number of ants")
		return 0
	}

	if ant <= 0 {
		fmt.Println("Error, Number of ants bellow zero")
		return 0
	}

	return ant
}
