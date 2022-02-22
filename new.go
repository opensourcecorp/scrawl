package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func newNote() {
	var editor string

	// Defaults to nano
	if editor = os.Getenv("EDITOR"); editor == "" {
		editor = "nano"
	}

	filepath := makeNoteFilepath()
	cmd := exec.Command(editor, filepath)

	// Need to set these so that the user has interactive access to the `EDITOR`
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("ERROR: error running editor program")
		os.Exit(1)
	}
}

func makeNoteFilepath() string {
	var timestamp string
	timestamp = time.Now().Format(time.RFC3339)
	timestamp = strings.ReplaceAll(timestamp, ":", "-")
	// Something's off between the time package in the Playgound & locally, and
	// how they're representing the same time format -- so, only keep the first
	// 20 characters of the timestamp and drop all the garbage at the end, in
	// case it creates it
	timestamp = timestamp[:19]
	filepath := fmt.Sprintf("%s/%s.md", getScrawldir(), timestamp)
	return filepath
}
