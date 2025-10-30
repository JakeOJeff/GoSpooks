package main

func setupWorld() {
	rooms = map[string]*Room {
		"Foyer": {
			Name:	"Foyer",
			Desc:	"A dimly lit entrance. A cold breeze pushes past you.",
			Exits:	map[string]string{"north": "Hallway"},
			Items:	[]string{},
		},
	}

	player = Player{CurrentRoom: "Foyer"}
	events = []Event{
		{"Someones here", 20},
	}
}