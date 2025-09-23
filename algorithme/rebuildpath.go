package algorithme

func reconstructPath(parent map[string]string, startRoomName, endRoomName string) []string {
	path := []string{}
	current := endRoomName

	for current != startRoomName {
		path = append([]string{current}, path...)
		current = parent[current]
	}
	path = append([]string{startRoomName}, path...)

	return path
}
