package main

import (
    "fmt"
	"strings"

)
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

func handlePickup(input string) bool {
	parts := strings.Split(input, " ")
	if len(parts) < 2 || parts[0] != "pickup" {
		
		return false
	}

	item := strings.Join(parts[1:], " ")
	room := rooms[player.CurrentRoom]

	for i, v := range room.Items {
		if strings.ToLower(v) == item {
			player.Inventory = append(player.Inventory, v)
			fmt.Println(" You picked up : ", v)

			room.Items = append(room.Items[:i], room.Items[i+1:]...)
			return true
		}
	}

	fmt.Println("That item doesn't exist here")
	return true
}

func hasItem(item string) bool {
	for _, v := range player.Inventory {
		if strings.ToLower(v) == strings.ToLower(item) {
			return true
		}
	}
	return false
}