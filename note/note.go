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
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

func (note Note) Display() {
	fmt.Printf("Your note titled : %v, has following content: \n\n%v\n\n", note.Title, note.Content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName)
	json, err := json.Marshal(note)
	if err != nil {
		return err
	}
	//os.WriteFile(fileName, json, 0644) it was not necessary to return it, it's just if error comes while writing file it will be reuturned from here only
	return os.WriteFile(fileName, json, 0644)
}
func New(title string, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input string can not be empty")
	}
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
