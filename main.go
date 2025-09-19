// Comment créer une carte compréhensible de la fourmilière pour le programme ?
// Comment la lui faire lire ?
// Comment la lui faire parcourir avec une fourmi ?
// Comment lui faire comprendre quelles salles ou quels tunnels sont occupés ?
// Comment lui faire comprendre qu'il doit prendre en compte ces lieux occupés pour recalculer le chemin le plus court ?

//MARK: BFS ?

//MARK: mettre les tunnels reliés à chaque salle dans la struct Rooms.
// Pour cela, je dois affilier chaque tunnel à chaque salle.
// J'ai déjà mis en place les éléments dans la Struct.
// Il faut construire la boucle de liaison pour distribuer les tunnels aux salles.

package main

import (
	"fmt"
	"lemin/extraction"
	"lemin/tidydata/ants"
	"lemin/tidydata/rooms"
	"lemin/tidydata/tunnels"
)

func main() {

	data := extraction.ReadFile()
	stringTable := extraction.PutFileInStringTable(data)
	ant := ants.TidyAnts(stringTable)

	roomsList := rooms.TidyRoomsInStruct(stringTable)

	roomsList = tunnels.TidyTunnelsInStruct(stringTable, roomsList)

	fmt.Printf("Nombre de fourmis : %d \n", ant)
	fmt.Println(roomsList)
}
