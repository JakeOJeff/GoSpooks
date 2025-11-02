package main

import (
	"fmt"
	"strings"
)

func handleMovement(direction string, current *Room) bool {
	current = rooms[player.CurrentRoom]
	clearScreen()
	printDefMessage(current)

	nextName, success := current.Exits[direction]
	if !success {
		return false
	}

	next := rooms[nextName]
	if next.Locked {
		if hasItem(next.Key) {
			fmt.Println("You use the", next.Key, " to unlock the door.")
			next.Locked = false
		} else {
			fmt.Println("The Door is locked. You need", next.Key)
			return true
		}
	}
	player.CurrentRoom = nextName
	current = rooms[player.CurrentRoom]
	fmt.Println("You entered", current.Name)
	fmt.Println(current.Desc)
	if len(current.Items) > 0 {
		fmt.Println("You see a", strings.Join(current.Items, ", "), ".")
	}
	return true
}

func handlePickup(input string, current *Room) bool {
	clearScreen()
	printDefMessage(current)
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
		if strings.EqualFold(v, item) {
			return true
		}
	}
	return false
}

func unlockAllRooms() {
	for _, room := range rooms {
		if room.Locked {
			room.Locked = false
			fmt.Println("Unlocked:", room.Name)
		}
	}
	fmt.Println("Unlocked all rooms!")
}

func unlockInteriorRooms() {
	for _, room := range rooms {
		if room.Locked && room.Name != "Patio" {
			room.Locked = false
			fmt.Println("Unlocked Interior Door:", room.Name)
		}
	}
	fmt.Println("Unlocked interior rooms!")
}
