package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string
	Content   string
	CreatedAt time.Time
}

func New(title string, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("title and content cannot be empty")
	}
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (n Note) Save() error {
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, error := json.Marshal(n)

	if error != nil {
		return error
	}

	return os.WriteFile(fileName, json, 0644)
}

func (n Note) Print() {
	fmt.Printf("Title: %s\nContent: %s\nCreated At: %s\n", n.Title, n.Content, n.CreatedAt)
}
