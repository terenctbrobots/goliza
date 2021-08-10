package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/terenctbrobots/goliza/eliza"
)

func ReplyEliza(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		message := r.FormValue("message")
		fmt.Fprintln(w, "Eliza: "+eliza.ReplyTo(message))
	}
}

func main() {
	args := os.Args[1:]

	if len(args) > 0 {
		if strings.ToLower(args[0]) == "--console" {
			console()
		}
	}

	http.HandleFunc("/reply", ReplyEliza)

	http.ListenAndServe(":8000", nil)
}

func console() {

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
