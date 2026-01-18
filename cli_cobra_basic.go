package main

import (
	"errors"
	"log"
	"os"

	"github.com/plh97/task-cli/cmd"
)

func main() {
	file, err := os.OpenFile("data.json", os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)

	if !errors.Is(err, os.ErrExist) {
		// Create the file
		_, err := os.Create("data.json")
		if err != nil {
			log.Fatal(err) // Handle potential errors (e.g., invalid path, permission issues)
		}
		return
	}

	defer file.Close() // Ensure the file is closed when the function exits
	cmd.Execute()
}
