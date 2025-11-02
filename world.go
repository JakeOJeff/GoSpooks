package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func setupWorld() {
	rooms = map[string]*Room{
		"Foyer": {
			Name: "Foyer",
			Desc: `A dimly lit entrance. A cold breeze pushes past you.
			|-----()-----|
			|____  	 ____|							
			-----Foyer----`,
			Exits: map[string]string{"north": "Hallway"},
			Items: []string{"blue key"},
		},
		"Hallway": {
			Name: "Hallway",
			Desc: `Old portraits stare at you from the walls. There’s a locked door to the east. 
			______________
			|------------|
			|------------|
			)---Hallway--(
			|------------|
			|____  	 ____|
			------()------`,
			Exits: map[string]string{"south": "Foyer", "east": "Library", "west": "Kitchen"},
			Items: []string{"silver key"},
		},
		"Kitchen": {
			Name: "Kitchen",
			Desc: `Pots and pans stay solemn as the doors waver
			`,
			Exits: map[string]string{"east": "Hallway"},
			Items: []string{"Sledge Hammer"},
		},
		"Library": {
			Name:   "Library",
			Desc:   "Dusty books cover the shelves as they point as if watching you every move.",
			Exits:  map[string]string{"west": "Hallway", "north": "Vault"},
			Locked: true,
			Key:    "silver key",
		},
		"Vault": {
			Name:   "Vault",
			Desc:   "This room is so dark and empty, you feel around for doors.",
			Exits:  map[string]string{"down": "Dark Basement", "south": "Library"},
			Locked: true,
			Key:    "Sledge Hammer",
		},
	}

	player = Player{CurrentRoom: "Foyer"}
	events = []Event{
		{"Someones here", 20},
		{"You hear a whisper, as it creeps near you", 40},
	}
}

func getExitList(exits map[string]string) string {
	keys := make([]string, 0, len(exits))
	for k := range exits {
		keys = append(keys, k)
	}
	return strings.Join(keys, ", ")
}

func triggerRandomEvent() {
	for _, e := range events {
		if rand.Intn(100) < e.Chance {
			fmt.Println("*OBSERVE*", e.Message)
			return
		}
	}
}
