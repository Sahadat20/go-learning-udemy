package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

// type displayer interface{
// 	Display()
// }

type outputable interface {
	saver
	Display()
}

// type outputable interface{
// 	Save() error
// 	Display()
// }

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text")

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = outputData(todo)

	// todo.Display()
	// err = saveData(todo)
	if err != nil {
		return
	}
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = outputData(userNote)
	// userNote.Display()
	// err = saveData(userNote)
	if err != nil {
		return
	}

}

func outputData(data outputable) error {
	data.Display()
	return saveData(data)

}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving failed")
		return err
	}
	fmt.Println("Save success")
	return nil

}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")

	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(promt string) string {
	fmt.Println(promt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}
