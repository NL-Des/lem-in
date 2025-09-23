package tunnels

import (
	"fmt"
	"lemin/tidydata/rooms"
	"strings"
)

// Récupère les deux coordonnées du tunnel (x, y)
func TidyTunnelsInStruct(stringTable []string, roomsList []rooms.Rooms) []rooms.Rooms {

	for i := 0; i < len(stringTable); i++ {

		if strings.Contains(stringTable[i], "-") {

			room1, room2, err0 := strings.Cut(stringTable[i], "-")

			if !err0 {
				fmt.Println("Error 0 on tunnels")
			}
			for i := range roomsList {
				if roomsList[i].Name == room1 {
					roomsList[i].Room1Name = append(roomsList[i].Room1Name, room1)
					roomsList[i].Room2Name = append(roomsList[i].Room2Name, room2)
				}
				if roomsList[i].Name == room2 {
					roomsList[i].Room1Name = append(roomsList[i].Room1Name, room2)
					roomsList[i].Room2Name = append(roomsList[i].Room2Name, room1)
				}
			}
		}
	}
	return roomsList
}
