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
			Desc: `YOU HAVE ESCAPED! Leave this madness and do not be tempted to enter again!.
			|-----()-----|
			|____  	 ____|							
			-----Foyer----`,
			Exits:  map[string]string{"north": "Hallway"},
			Locked: true,
			Key:    "White Key",
		},
		"Hallway": {
			Name: "Hallway",
			Desc: `The door locks behind you. Find a way out! Old portraits stare at you from the walls. There’s a locked door to the east. 
			------()------
			_____  	 _____
			|------------|
			|---------[]-|
			|------------|
			|------------|
			|------------|
			)---Hallway--(
			|------------|
			|____  	 ____|
			------()------`,
			Exits: map[string]string{"south": "Foyer", "east": "Library", "west": "Kitchen", "north": "bedroom", "up": "attic"},
			Items: []string{"Silver Key"},
		},
		"Kitchen": {
			Name: "Kitchen",
			Desc: `Pots and pans stay solemn as the doors waver
			------()------
			_____  	 _____
			|------------|
			|--Kitchen---|
			|------------(
			|____________|`,
			Exits: map[string]string{"east": "Hallway", "north": "Storage"},
			Items: []string{"Magic Driller"},
		},
		"Storage": {
			Name: "Storage",
			Desc: `You are now inside the dusty storage room
			______________
			|------------|
			|------------|
			|---Storage--|
			|------------|
			|------------|
			|____    ____|`,
			Locked: true,
			Key:    "Green Key",
			Exits:  map[string]string{"south": "Kitchen"},
			Items:  []string{"Golden Key"},
		},
		"Library": {
			Name: "Library",
			Desc: `Dusty books cover the shelves as they point as if watching you every move.
			-------------------------------()------------
			-------------------___________    ___________
			-------------------|[--][--][------][--][--]|
			-------------------|[--][--][------][--][--]|
			-------------------|[--][--][------][--][--]|
			-------------------|[--][--][------][--][--]|
			___________________|[--][--][------][--][--]|
			|-------------------------------------------|
			|-------------------------------------------|
			)------------------Library------------------|
			|-------------------------------------------|
			|___________________________________________|
			`,
			Exits:  map[string]string{"west": "Hallway", "north": "Vault"},
			Locked: true,
			Key:    "Silver Key",
		},
		"Vault": {
			Name: "Vault",
			Desc: `This room is so dark and empty, you feel around for doors.
			-----------________-----------
			----------/--------\----------
			---------/----[]----\---------
			--------/------------\--------
			-------/--------------\-------
			-------|--------------|-------
			-------|-----Vault----|-------
			-------|--------------|-------
			-------\--------------/-------
			--------\------------/--------
			---------\----------/---------
			----------\--------/----------
			-----------|_    _|-----------
			--------------()--------------`,
			Exits:  map[string]string{"down": "Basement", "south": "Library"},
			Locked: true,
			Key:    "Magic Driller",
			Items:  []string{"Green Key"},
		},
		"Basement": {
			Name: "Basement",
			Desc: `You slowly descended into the basement similar to your madness.
			________________
			|--------------|
			|------[]------|
			|---Basement---|
			|--------------|
			|--------------|
			|______________|
			---()-----------`,
			Exits:  map[string]string{"up": "Vault", "south": "Cracked"},
			Locked: true,
			Key:    "Golden Key",
			Items:  []string{"Red Key"},
		},
		"Cracked": {
			Name: "Cracked",
			Desc: `You discovered a cracked door which smelled musty and old leading to a long corridor with a planked door.
			--__    __----------------
			--|------|----------------
			--|------|----------------
			--|------|----------------
			--|------|----------------
			--|------|-----Cracked----
			--|------|-----Hallway----
			--|------|----------------
			--|------|----------------
			--|/\//\/|----------------`,
			Exits:  map[string]string{"north": "Basement", "south": "Grain"},
			Locked: true,
			Key:    "Gold Coin",
		},
		"Grain": {
			Name: "Grain",
			Desc: `You have discovered the Grain, get the final key here and escape this madness!
			--/------\---------------
			--|------|---------------
			__|------|_______________
			|-+----------+----+-----|
			|---+------+----------+-|
			|-------+-----+---------|
			|--+---------+-----+----|
			|-----+---Grain------+--|
			|---+-----------+-------|
			|-+-----+----------+----|
			|----+-----+----------+-|
			|_______________________|
			`,
		},
		"Bedroom": {
			Name: "Bedroom",
			Desc: `You enter the bedroom where the moss covered mattresses gives you a headache.
			___________________________
			|-------------------------|
			|-------------------------|
			|--[====]-----Bedroom-----|
			|--[====]-----------------|
			|-------------------------|
			|___________________    __|
			---------------------()----`,
			Exits:  map[string]string{"south": "Hallway"},
			Locked: true,
			Key:    "Red Key",
			Items:  []string{"Gold Coin"},
		},
		"Attic": {
			Name: "Attic",
			Desc: `The only place you would expect be filled with cobwebs look newly made.
			-----------------------------------------___________________
			-----------------------------------------|-----------------|
			-----------------------------------------|-----------------|
			-----------------------------------------|-----------------|
			-----------------------------------------|-----------------|
			-----------------------------------------|-----------------|
			_________________________________________|-----------------|
			|----------------------------------------------------------|
			|----------------------------------------------------------|
			|----------------------------------------------------------|
			|----------------------------------------------------------|
			|----------------------------------------------------------|
			|--------------Attic---------------------------------------|
			|----------------------------------------------------------|
			|----------------------------------------------------------|
			|----------------------------------------------[]----------|
			|----------------------------------------------------------|
			|----------------------------------------------------------|
			|----------------------------------------------------------|
			|__________________________________________________________|
			`,
			Exits: map[string]string{"down": "Hallway"},
			Items: []string{"Electric Chainsaw"},
		},
	}

	player = Player{CurrentRoom: "Hallway"}
	events = []Event{
		{"Someones here", 20},
		{"You hear a whisper, as it creeps near you", 40},
		{"GET OUT! GET OUT! GET OUT!", 40},
		{"You just won't leave the dead alone will you", 50},
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
