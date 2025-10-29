package main

import (
    "fmt"
	"time"
	"strings"
	// "math/rand"
)

var (
	rooms	map[string]*Room
	player	Player
	events	[]Event
)

func main() {
	// sec := time.Second
    printMessages(beginnerMessage)

    var input string

    for {
		// current := rooms[player.CurrentRoom]


        fmt.Print("> ")
        fmt.Scanln(&input) // waits for user input
		input = strings.ToLower(input)

        switch input {
        case "n","no":
            fmt.Println("You step out into the fog... never to return.")
            return
		case "y", "yes":
			printMessages(houseEnter)
        case "look":
            fmt.Println("The walls seem to breathe... dust fills the air.")
        default:
            fmt.Println("You whisper:", input)
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
