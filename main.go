package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func clearconsole() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	var Events []string
	var UserSelection string
	var UserInput string

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Welcome to ToDo app")
		fmt.Println("1 - all events")
		fmt.Println("2 - add event")
		fmt.Println("3 - exit")

		fmt.Print("your select: ")
		UserSelectionRaw, _ := reader.ReadString('\n')
		UserSelection = strings.TrimRight(UserSelectionRaw, "\r\n")

		switch UserSelection {
		case "1":
			if len(Events) <= 0 {
				fmt.Println("you have 0 events to do")
			} else {
				clearconsole()
				fmt.Println("All events:")
				for _, event := range Events {
					fmt.Println(event)
				}
				fmt.Println("")
			}
		case "2":
			fmt.Print("Enter event: ")
			UserInputRaw, _ := reader.ReadString('\n')
			UserInput = strings.TrimRight(UserInputRaw, "\r\n")
			Events = append(Events, UserInput)
			clearconsole()
			fmt.Println("Event added successfully.")
		case "3":
			fmt.Println("Exiting the ToDo app. Goodbye!")
			return
		default:
			clearconsole()
			fmt.Println("Invalid selection. Please choose again.")
		}
	}
}
