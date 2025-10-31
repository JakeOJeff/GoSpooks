package main

import (
	"bufio"
    "fmt"
	"time"
	"strings"
	"os"
	"math/rand"
)

var (
	rooms	map[string]*Room
	player	Player
	events	[]Event
)

func main() {
	// sec := time.Second
	var insideHouse bool = false
	var current *Room

    printMessages(beginnerMessage)
	scanner := bufio.NewScanner(os.Stdin)

    var input string
	rand.Seed(time.Now().UnixNano())
	setupWorld()

    for {
		if insideHouse {
			current = rooms[player.CurrentRoom]
			fmt.Printf("\nYou are in: %s\n%s\n", current.Name, current.Desc)
			if len(current.Items) > 0 {
				fmt.Println("You see ", strings.Join(current.Items, ", "))
			}

			triggerRandomEvent()

			fmt.Println(" Exits: ", getExitList(current.Exits))
		}
		

        fmt.Print("> ")
        scanner.Scan()
    	input = strings.TrimSpace(strings.ToLower(scanner.Text()))

		if !insideHouse {
			switch input {
			case "n","no", "leave":
					fmt.Println("You step out into the fog... never to return.")

				return
			case "y", "yes":
				printMessages(houseEnter)
				insideHouse = true

			case "look":
				fmt.Println("The walls seem to breathe... dust fills the air.")
			default:
				fmt.Println("You whisper:", input)
			}
		} else if insideHouse {
			switch input {
				
			case "look":
				fmt.Println(current.Desc)		
			}
		}
    }
}

func printMessages(messages []string){
	delay := 500 * time.Millisecond
	for _, msg := range messages {
		fmt.Println(msg)
		time.Sleep(delay)
	}
}
