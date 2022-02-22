// Command-line parsing for subcommands inspired from here:
// https://gobyexample.com/command-line-subcommands
// I really wanted to get the following approach working, but it's not kicking off the external command to start the editor (but, it's still commented out in here to maybe try again later):
// https://www.digitalocean.com/community/tutorials/how-to-use-the-flag-package-in-go
package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func showHelp(exitCode int) {
	helpText := fmt.Sprintln(`
Usage: scrawl [-h | --help] [new | search [--any (default) | --all] <tags> | search --untagged]

scrawl is a simple note-taking application. By running:

	scrawl new

you will enter your editor ($EDITOR) to take notes.

scrawl note files expect your note's tags to appear as a comma-separated list of
tag values, specified as in this example:

	tags: dev, infra, ideas
	# rest of your note(s) below
	# ...

The tag spec can appear anywhere in the file, but it's probably best to put them
up top. You can then search for notes with certain tags via:

	scrawl search [--any (default) | --all] <tags>

By default, all notes end up as filepaths named '$HOME/scrawl/[timestamp].md'.
You can override the root directory by setting the 'SCRAWLDIR' environment
variable.
	`)
	fmt.Println(strings.TrimSpace(helpText))
	os.Exit(exitCode)
}

// getScrawldir determines which directory to place scrawl notes. Defaults to $HOME/scrawl
func getScrawldir() string {
	var scrawldir string

	if scrawldir = os.Getenv("SCRAWLDIR"); scrawldir == "" {
		user, err := user.Current()
		if err != nil {
			fmt.Println("ERROR: could not determine home directory for current user, and SCRAWLDIR var not found, so cannot set a directory for scrawl to use")
			os.Exit(1)
		}
		homedir := user.HomeDir
		scrawldir = filepath.Join(homedir, "scrawl")
	}

	err := os.MkdirAll(scrawldir, 0755)
	if err != nil {
		fmt.Printf("ERROR: Could not create SCRAWLDIR (%s)\n", scrawldir)
		os.Exit(1)
	}

	return scrawldir
}

func main() {
	newCmd := flag.NewFlagSet("new", flag.ExitOnError)

	searchCmd := flag.NewFlagSet("search", flag.ExitOnError)
	searchAnyTags := searchCmd.Bool("any", true, "Search notes for ANY of the provided tags")
	searchAllTags := searchCmd.Bool("all", false, "Search notes for ALL of the provided tags")
	searchUntagged := searchCmd.Bool("untagged", false, "Search for notes that are missing tags")

	if len(os.Args) < 2 {
		fmt.Printf(">>> ERROR: no command specified\n\n")
		showHelp(1)
	} else {
		switch os.Args[1] {

		case "new":
			newCmd.Parse(os.Args[2:])
			newNote()

		case "search":
			searchCmd.Parse(os.Args[2:])
			// TODO: hopefully tags are the last arg passed
			wantTags := os.Args[len(os.Args)-1]
			searchTags(*searchAnyTags, *searchAllTags, *searchUntagged, wantTags)

		case "render":
			fmt.Println(">>> ERROR: not yet implemented")
			os.Exit(1)

		case "web":
			fmt.Println(">>> ERROR: not yet implemented")
			os.Exit(1)

		case "sync":
			fmt.Println(">>> ERROR: not yet implemented")
			os.Exit(1)

		default:
			fmt.Printf(">>> ERROR: must specify a valid command\n\n")
			showHelp(1)
		}
	}
}
