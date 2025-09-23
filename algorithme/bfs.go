package algorithme

func ResolverPath(startRoomName string, endRoomName string, connexions map[string][]string) map[string][]string {
	// BFS
	listOfRooms := []string{startRoomName}
	used := make(map[string]bool)
	parent := make(map[string]string)

	used[startRoomName] = true

	for len(listOfRooms) > 0 {
		current := listOfRooms[0]
		listOfRooms = listOfRooms[1:]

		if current == endRoomName {
			// Reconstruire le chemin et l'emballer dans une map
			path := reconstructPath(parent, startRoomName, endRoomName)
			result := make(map[string][]string)
			result["path"] = path
			return result
		}

		for _, neighbor := range connexions[current] {
			if !used[neighbor] {
				used[neighbor] = true
				parent[neighbor] = current
				listOfRooms = append(listOfRooms, neighbor)
			}
		}
	}

	return nil // Pas de chemin trouvé
}
