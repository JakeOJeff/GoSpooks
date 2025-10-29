package main

import (
	"bufio"
    "fmt"
	"time"
	"strings"
	"os"
	// "math/rand"
)

var (
	rooms	map[string]*Room
	player	Player
	events	[]Event
)

func main() {
	// sec := time.Second
	var insideHouse bool = false

    printMessages(beginnerMessage)
	scanner := bufio.NewScanner(os.Stdin)

    var input string

    for {
		// current := rooms[player.CurrentRoom]


        fmt.Print("> ")
        scanner.Scan()
    	input = strings.TrimSpace(strings.ToLower(scanner.Text()))


        switch input {
        case "n","no", "leave":
			if !insideHouse {
				fmt.Println("You step out into the fog... never to return.")
			} else {
				fmt.Println("You leave the house in a hurry... never to return")
			}
            return
		case "y", "yes":
			printMessages(houseEnter)
			insideHouse = true
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
