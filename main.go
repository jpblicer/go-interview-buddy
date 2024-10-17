package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to Go Interview Buddy!")
	fmt.Println("Please choose an option:")
	fmt.Println("1) Add a Company")
	fmt.Println("2) List all Companies")
	fmt.Println("3) Exit")

	getUserInput()
}

func getUserInput() {
	reader := bufio.NewReader(os.Stdin)

	userInput, _ := reader.ReadString('\n')

	fmt.Print(userInput)

	if userInput == "1\n" {
		fmt.Println("Add a Company")
	} else if userInput == "2\n" {
		fmt.Println("List all Companies")
	} else if userInput == "3\n" {
		fmt.Println("Exit")
	} else {
		main()
	}

}
