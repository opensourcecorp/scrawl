package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"sort"
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

// getScrawldir determines which directory to place scrawl notes, and will
// create said dir if it does not exist. Defaults to $HOME/scrawl
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

func setTestScrawldir() {
	err := os.Setenv("SCRAWLDIR", "./testdata")
	if err != nil {
		panic(err)
	}
}