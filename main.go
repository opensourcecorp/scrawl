// Command-line parsing for subcommands inspired from here:
// https://gobyexample.com/command-line-subcommands
// I really wanted to get the following approach working, but it's not kicking off the external command to start the editor (but, it's still commented out in here to maybe try again later):
// https://www.digitalocean.com/community/tutorials/how-to-use-the-flag-package-in-go
package main

import (
	"flag"
	"fmt"
	"os"
)

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
			gotTagMap := searchTags(*searchAnyTags, *searchAllTags, *searchUntagged, wantTags)
			printSortedTagMap(gotTagMap)

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
