package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notepad/note"
)

func main() {
	title, content := getNoteData()

	noteStruct, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	noteStruct.Print()
	err = noteStruct.Save()

	if err != nil {
		fmt.Println("saved failed")
		return
	}

	fmt.Println("saved successfully")

	// fmt.Printf("Title: %s\nContent: %s\n", title, content)

}

func getNoteData() (string, string) {
	title := getUserInput("Title: ")

	content := getUserInput("Content: ")

	return title, content

}

func getUserInput(promptText string) string {
	fmt.Print(promptText)
	// var text string
	// fmt.Scanln(&text)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
