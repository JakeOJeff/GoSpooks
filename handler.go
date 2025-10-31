package main

func handleMovement(direction string) bool {
	current := rooms[player.CurrentRoom]
	nextName, success := current.Exits[direction]
	if !success {
		return false
	}

	next := rooms[nextName]
	if next.Locked {
		if hasItem(next.Key) {
			fmt.Println("You use the ", next.Key, " to unlock the door.")
			next.Locked = false
		} else {
			fmt.Println("The Door is locked. You need ", next.Key)
			return true
		}
	}

	player.CurrentRoom = nextName
	return true
}