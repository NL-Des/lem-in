package extraction

import (
	"fmt"
	"os"
)

// Lit le fichier de test.
func ReadFile() string {
	input, err := os.ReadFile("/home/zone01student/Exercices/lemin/filestests/test.txt")
	if err != nil {
		fmt.Printf("File not found.\n")
	}
	data := string(input)
	return data
}

/* func ReadFile() string {
	if len(os.Args) != 2 {
		return "Please enter the path of the file to resolve"
	}
	path := os.Args[1]
	input, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("File not found.\n")
	}
	data := string(input)
	return data
} */

// Mise en []string.
func PutFileInStringTable(data string) []string {
	var runeTable []rune
	var stringTable []string

	for _, caracter := range data {
		if caracter == '\n' {
			stringTable = append(stringTable, string(runeTable))
			runeTable = nil
		} else {
			runeTable = append(runeTable, caracter)
		}

	}
	// Affichage du fichier dans la console.
	for i := 0; i < len(stringTable); i++ {
		//fmt.Println(stringTable[i])
	}
	return stringTable
}
