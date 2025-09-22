package algorithme

import (
	"fmt"
	"lemin/tidydata/rooms"
)

func GetStartRoom(roomsList []rooms.Rooms) string {
	var startRoomName string
	for _, room := range roomsList {
		if room.Start {
			startRoomName = room.Name
		}
	}
	if len(startRoomName) == 0 {
		fmt.Println("Error, no information in startRoomName")
	}
	return startRoomName
}
