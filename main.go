// Comment créer une carte compréhensible de la fourmilière pour le programme ?
// Comment la lui faire lire ?
// Comment la lui faire parcourir avec une fourmi ?
// Comment lui faire comprendre quelles salles ou quels tunnels sont occupés ?
// Comment lui faire comprendre qu'il doit prendre en compte ces lieux occupés pour recalculer le chemin le plus court ?

//MARK: BFS ?

//MARK: strings.Field à étudier pour bien refaire.

//MARK: mettre les tunnels reliés à chaque salle dans la struct Rooms.
// Pour cela, je dois affilier chaque tunnel à chaque salle.
// J'ai déjà mis en place les éléments dans la Struct.
// Il faut construire la boucle de liaison pour distribuer les tunnels aux salles.

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

/* type StartRoom struct {
	Name string
	x    int
	y    int
}

type EndRoom struct {
	Name string
	x    int
	y    int
} */

type Rooms struct {
	Start    bool
	End      bool
	Name     string
	x        int
	y        int
	TunnelsX int
	TunnelsY int
}

type Tunnels struct {
	x int
	y int
}

func main() {
	data := readFile()
	stringTable := putFileInStringTable(data)
	ants := tidyAnts(stringTable)
	tidyRoomsInStruct(stringTable)
	tidyTunnelsInStruct(stringTable)
	fmt.Println(ants)
}

// Récupère les trois coordonnées de chaque salle (Nom, x, y)
func tidyRoomsInStruct(stringTable []string) {
	var RoomsList []Rooms
	// var EndRoomList []EndRoom
	// var StartRoomList []StartRoom

	// Boucle pour trouver les salles.
	for i := 0; i < len(stringTable); i++ {

		//Pour trouver la Start Room.
		if strings.Contains(stringTable[i], "##start") {
			isStart := true
			isEnd := false
			i++
			first, rest, err0 := strings.Cut(stringTable[i], " ")
			if !err0 {
				fmt.Println("Error 0 on StartRooms")
			}
			second, third, err1 := strings.Cut(rest, " ")
			if !err1 {
				fmt.Println("Error 1 on StartRooms")
			}
			name := first
			x, err2 := strconv.Atoi(second)
			if err2 != nil {
				fmt.Println("Error 2 on StartRooms")
			}
			y, err3 := strconv.Atoi(third)
			if err3 != nil {
				fmt.Println("Error 3 on StartRooms")
			}
			RoomsList = append(RoomsList, Rooms{Start: isStart, End: isEnd, Name: name, x: x, y: y})
		}

		// Pour trouver la End Room.
		if strings.Contains(stringTable[i], "##end") {
			isStart := false
			isEnd := true
			i++
			first, rest, err0 := strings.Cut(stringTable[i], " ")
			if !err0 {
				fmt.Println("Error 0 on EndRooms")
			}
			second, third, err1 := strings.Cut(rest, " ")
			if !err1 {
				fmt.Println("Error 1 on EndRooms")
			}
			name := first
			x, err2 := strconv.Atoi(second)
			if err2 != nil {
				fmt.Println("Error 2 on EndRooms")
			}
			y, err3 := strconv.Atoi(third)
			if err3 != nil {
				fmt.Println("Error 3 on EndRooms")
			}
			RoomsList = append(RoomsList, Rooms{Start: isStart, End: isEnd, Name: name, x: x, y: y})
		}

		// Pour enregistrer les autres Rooms restantes.
		if strings.Contains(stringTable[i], " ") && !strings.Contains(stringTable[i-1], "##end") && !strings.Contains(stringTable[i-1], "##start") {
			first, rest, err0 := strings.Cut(stringTable[i], " ")
			if !err0 {
				fmt.Println("Error 0 on EndRooms")
			}
			second, third, err1 := strings.Cut(rest, " ")
			if !err1 {
				fmt.Println("Error 1 on EndRooms")
			}
			name := first
			x, err2 := strconv.Atoi(second)
			if err2 != nil {
				fmt.Println("Error 2 on EndRooms")
			}
			y, err3 := strconv.Atoi(third)
			if err3 != nil {
				fmt.Println("Error 3 on EndRooms")
			}
			RoomsList = append(RoomsList, Rooms{Name: name, x: x, y: y})
		}
	}
	// fmt.Println("Coordonnées de Start :", StartRoomList) // test d'affichage pour vérifier si c'est bien pris en compte.
	// fmt.Println("Coordonnées de End :", EndRoomList)
	fmt.Println("Coordonnées des salles :", RoomsList)

}

// Récupère les deux coordonnées du tunnel (x, y)
func tidyTunnelsInStruct(stringTable []string) {
	var tunnelsList []Tunnels

	for i := 0; i < len(stringTable); i++ {

		if strings.Contains(stringTable[i], "-") {

			before, after, err0 := strings.Cut(stringTable[i], "-")

			if err0 == false {
				fmt.Println("Error 0 on tunnels")
			}

			x, err1 := strconv.Atoi(before)
			if err1 != nil {
				fmt.Println("Error 1 on tunnels")
			}

			y, err2 := strconv.Atoi(after)
			if err2 != nil {
				fmt.Println("Error 0 on tunnels")
			}
			tunnelsList = append(tunnelsList, Tunnels{x: x, y: y})
		}
	}
	fmt.Println("Coordonnées des tunnels :", tunnelsList)

}

// Récupère la fourmi.
func tidyAnts(stringTable []string) int {
	ants, err := strconv.Atoi(stringTable[0])
	if err != nil {
		fmt.Println("Erreur sur le nombre de fourmis")
	}
	fmt.Printf("Nombre de fourmis : %d \n", ants)
	return ants
}

// Mise en []string.
func putFileInStringTable(data string) []string {
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

// Lit le fichier texte pour pouvoir travailler dessus.
func readFile() string {
	input, err := os.ReadFile("tests-files/test.txt")
	if err != nil {
		fmt.Printf("File not found.\n")
	}
	data := string(input)
	return data
}
