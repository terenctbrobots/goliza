package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/terenctbrobots/goliza/eliza"
)

func main() {
	greetings := eliza.Greetings()

	fmt.Println("Eliza: " + greetings)

	for {
		statement := getInput()

		fmt.Println("Eliza: " + eliza.ReplyTo(statement))
	}
}

func getInput() string {
	fmt.Print("Awak: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input
}
