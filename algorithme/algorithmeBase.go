package algorithme

import (
	"lemin/tidydata/rooms"
)

type Connexions map[string][]string // Map des connections des connections entre les salles.

// Construction de la map
func AlgorithmeTest(roomsList []rooms.Rooms) Connexions {
	connexions := make(Connexions)

	// Construction des salles.
	for i := 0; i < len(roomsList); i++ { // Parcours l'index de la structure.

		for y := 0; y < len(roomsList[i].Room1Name); y++ { // Parcours les tunnels enregistrés sous chaque nom de salle.
			room1 := roomsList[i].Room1Name[y]
			room2 := roomsList[i].Room2Name[y]

			// Ajout des tunnels de liaisons entre les salles.
			connexions[room1] = append(connexions[room1], room2)
			connexions[room2] = append(connexions[room2], room1)
		}
	}
	return connexions
}

