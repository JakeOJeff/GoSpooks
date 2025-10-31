package main

func setupWorld() {
	rooms = map[string]*Room {
		"Foyer": {
			Name:	"Foyer",
			Desc:	"A dimly lit entrance. A cold breeze pushes past you.",
			Exits:	map[string]string{"north": "Hallway"},
			Items:	[]string{},
		},		
		"Hallway": {
			Name:        "Hallway",
			Description: "Old portraits stare at you from the walls. There’s a locked door to the east.",
			Exits:       map[string]string{"south": "Foyer", "east": "Library", "west": "Kitchen"},
			Items:       []string{"silver key"},
		},
	}

	player = Player{CurrentRoom: "Foyer"}
	events = []Event{
		{"Someones here", 20},
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
			fmt.Println("*DIED*", e.Message)
			return
		}
	}
}