package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/jpblicer/go-interview-buddy/company"
)

func main() {
	displayMenu()
}

func displayMenu() {
	fmt.Println("Welcome to Go Interview Buddy!")
	fmt.Println("Please choose an option:")
	fmt.Println("1) Add a Company")
	fmt.Println("2) List all Companies")
	fmt.Println("3) Exit")
	getUserMenuInput()
}

func getUserMenuInput() {
	reader := bufio.NewReader(os.Stdin)

	userInput, _ := reader.ReadString('\n')

	if userInput == "1\n" {
		AddCompany()
	} else if userInput == "2\n" {
		ListCompanies()
	} else if userInput == "3\n" {
		fmt.Println("Exit")
	} else {
		displayMenu()
	}

}

func AddCompany() {
	fmt.Println("Please enter the name of the new Company:")
	reader := bufio.NewReader(os.Stdin)

	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)

	company := company.Add(userInput)

	fmt.Printf("Added company: %s", company.Name)
	fmt.Println()
	fmt.Println()

	displayMenu()
}

func ListCompanies() {
	companies := company.List()

	clearScreen()

	fmt.Println("List of all Companies:")

	for index, company := range companies {
		fmt.Printf("%d) %s \n", index+1, company.Name)
	}
	fmt.Println()
	fmt.Println()

	displayMenu()
}

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls") // Windows clear screen
	} else {
		cmd = exec.Command("clear") // Unix clear screen
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
