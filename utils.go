package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

func showHelp(exitCode int) {
	helpText := `
Usage: scrawl [help] [new | search [--any (default) | --all] <tags> | --untagged]

scrawl is a simple note-taking application that supports tagging. By running:

	scrawl new

you will enter your editor (as defined by $EDITOR) to take notes. scrawl will
fall back to the nano editor if the $EDITOR env var is not defined.

scrawl note files expect your note's tags to appear as a comma-separated list of
tag values, specified as in this example:

	tags: dev, infra, ideas
	# rest of your note(s) below
	# ...

The tag spec can appear anywhere in the file, but it's probably best to put them
at or near the top. You can then search for notes with certain tags via:

	scrawl search [--any (default) | --all] <tags>

By default, all notes end up as filepaths named '$HOME/scrawl/[timestamp].md'.
You can override the root directory by setting the 'SCRAWLDIR' environment
variable.
`
	fmt.Println(strings.TrimSpace(helpText))
	os.Exit(exitCode)
}

// getScrawldir determines which directory to place scrawl notes, and will
// create said dir if it does not exist. Defaults to $HOME/scrawl
func getScrawldir() string {
	var scrawldir string

	if scrawldir = os.Getenv("SCRAWLDIR"); scrawldir == "" {
		user, err := user.Current()
		if err != nil {
			log.Fatalf("could not determine home directory for current user, and SCRAWLDIR var not found, so cannot set a directory for scrawl to use\n%v", err)
		}
		homedir := user.HomeDir
		scrawldir = filepath.Join(homedir, "scrawl")
	}

	if err := os.MkdirAll(scrawldir, 0755); err != nil {
		log.Fatalf("could not create SCRAWLDIR (%s)\n%v", scrawldir, err)
	}

	return scrawldir
}

// containsStringValue will do a regex match on each value of the input slice
// against the value provided
func containsStringValue(slice []string, value string) bool {
	for _, e := range slice {
		// if e == value {
		re := regexp.MustCompile(e)
		if re.MatchString(value) {
			return true
		}
	}
	return false
}

// printSortedTagMap ensures that printing the map of matching tags is sorted by
// filename when output (since filenames are timestamps)
func printSortedTagMap(tm tagMap) {
	keys := make([]string, 0, len(tm))
	for key := range tm {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(k, tm[k])
	}
}
