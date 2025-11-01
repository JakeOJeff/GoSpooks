package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

var (
	rooms  map[string]*Room
	player Player
	events []Event
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
		}

		fmt.Print("> ")
		scanner.Scan()
		input = strings.TrimSpace(strings.ToLower(scanner.Text()))

		if !insideHouse {
			switch input {
			case "n", "no", "leave":
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
				clearScreen()
				printDefMessage(current)
				fmt.Println(current.Desc)
				if len(current.Items) > 0 {
					fmt.Println("You see a", strings.Join(current.Items, ", "), ". You can pick them up for further use.")
				}
				fmt.Println("============================")
			case "inventory":
				clearScreen()
				printDefMessage(current)
				if len(player.Inventory) == 0 {
					fmt.Println("You have nothing on you.")

				} else {
					fmt.Println("Inventory: ", strings.Join(player.Inventory, ", "))
				}
				fmt.Println("============================")

			case "exits":
				clearScreen()
				printDefMessage(current)
				fmt.Println(" > Exits: ", getExitList(current.Exits))
				fmt.Println("============================")

			case "help":
				clearScreen()
				printDefMessage(current)
				printMessages(helpMsgs)
				fmt.Println("============================")

			case "clear":
				clearScreen()
			default:
				if handleMovement(input) {
					current = rooms[player.CurrentRoom]
					clearScreen()
					printDefMessage(current)
					fmt.Println("You entered", current.Name)
					fmt.Println(current.Desc)
					if len(current.Items) > 0 {
						fmt.Println("You see a", strings.Join(current.Items, ", "), ".")
					}
					fmt.Println("============================")
					continue
				} else if handlePickup(input) {
					continue
				} else {
					fmt.Println("Such a command cannot be used?")
				}
			}
		}
	}
}

func printMessages(messages []string) {
	delay := 500 * time.Millisecond
	for _, msg := range messages {
		fmt.Println(msg)
		time.Sleep(delay)
	}
}

func clearScreen() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear") // future proofing incase linux
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func printDefMessage(current *Room) {
	fmt.Println("============================")
	// fmt.Printf("\nYou are in: %s\n%s\n", current.Name, current.Desc)
	// if len(current.Items) > 0 {
	// 	fmt.Println("You see a", strings.Join(current.Items, ", "), ". You can pick them up for further use.")
	// }
	triggerRandomEvent()

}
