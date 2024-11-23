package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note-app/note"
)

func main() {
	title, content := getNoteData()
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	userNote.Display()
	err = userNote.Save()
	if err != nil {
		fmt.Print("Error saving the note")
		return
	}
	fmt.Println("Saving the Note Succeeded")
}

func getNoteData() (string, string) {
	title := getUserInput("Note Title: ")
	content := getUserInput("Note Content: ")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	//var value string
	//fmt.Scanln(&value)   for longer text reads this is not realiable it won't work
	return text
}
