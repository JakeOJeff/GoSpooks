package main

import (
    "fmt";
	"time";
)

func main() {
	// sec := time.Second
	delay := 500 * time.Millisecond


    fmt.Println("Welcome to your GoSpooks adventure!")
	time.Sleep(delay)
    fmt.Println("There are some things you must know beforehand")


    var input string

    for {
        fmt.Print("> ")
        fmt.Scanln(&input) // waits for user input

        switch input {
        case "quit":
            fmt.Println("You step out into the fog... never to return.")
            return
        case "look":
            fmt.Println("The walls seem to breathe... dust fills the air.")
        default:
            fmt.Println("You whisper:", input)
        }
    }
}

