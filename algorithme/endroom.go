package algorithme

import (
	"fmt"
	"lemin/tidydata/rooms"
)

func GetEndRoom(roomsList []rooms.Rooms) string {
	var endRoomName string
	for _, room := range roomsList {
		if room.Start {
			endRoomName = room.Name
		}
	}
	if len(endRoomName) == 0 {
		fmt.Println("Error, no information in endRoomName")
	}
	return endRoomName
}
