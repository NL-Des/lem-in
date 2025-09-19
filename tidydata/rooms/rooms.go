//MARK: strings.Field à étudier pour bien refaire.

package rooms

import (
	"fmt"
	"strconv"
	"strings"
)

type Rooms struct {
	Start     bool     // Si c'est l'entrée, true.
	End       bool     // Si c'est la sortie, true.
	Name      string   // Nom de la salle.
	x         int      // Coordonnée x de la salle.
	y         int      // Coordonnée y de la salle.
	Room1Name []string // Coordonnées X des tunnels reliant les salles. (Nom de la salle)
	Room2Name []string // Coordonnées y des tunnels reliant les salles. (Nom de la salle)
}

var roomsList []Rooms

// Récupère les trois coordonnées de chaque salle (Nom, x, y)
func TidyRoomsInStruct(stringTable []string) []Rooms {

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
			roomsList = append(roomsList, Rooms{Start: isStart, End: isEnd, Name: name, x: x, y: y})
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
			roomsList = append(roomsList, Rooms{Start: isStart, End: isEnd, Name: name, x: x, y: y})
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
			roomsList = append(roomsList, Rooms{Name: name, x: x, y: y})
		}
	}
	return roomsList
}
