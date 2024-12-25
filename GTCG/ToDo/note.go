package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type saver interface

type Todo struct {
	Text string `json:"text"`
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{Text: ""}, errors.New("content cannot be empty")
	}

	return Todo{Text: content}, nil
}

func (todo Todo) Display() {
	fmt.Println(todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	json, error := json.Marshal(todo)

	if error != nil {
		fmt.Println("error")
	}

	return os.WriteFile(fileName, json, 0644)
}

func main() {

}
