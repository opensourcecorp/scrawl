package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func searchTags(anyTags, allTags, untagged bool, wantTagsArg string) map[string][]string {
	// Only ONE of these flags can be specified, so flip anyTags off if allTags is spec'd
	// Fun fact: staticcheck rule won't let me do anyTags = !allTags, yeesh
	if allTags {
		anyTags = false
	}

	// If specifying -untagged at all, negate the other criteria
	if untagged {
		anyTags = false
		allTags = false
	}

	scrawldir := getScrawldir()

	files, err := os.ReadDir(scrawldir)
	if err != nil {
		fmt.Println("ERROR: could not access SCRAWLDIR files")
		os.Exit(1)
	}

	var filenames []string
	for _, file := range files {
		filenames = append(filenames, filepath.Join(scrawldir, file.Name()))
	}

	wantTags := strings.Split(wantTagsArg, ",")
	matchingTags := make(map[string][]string)
	for _, filename := range filenames {
		contentsRaw, err := os.ReadFile(filename)
		if err != nil {
			fmt.Printf("ERROR: error while reading file %s\n", filename)
		}

		contents := strings.Split(string(contentsRaw), "\n")

		var tagsInFile []string
		for _, line := range contents {
			tagPattern := regexp.MustCompile("^tags:.*$")
			if tagPattern.MatchString(line) {
				// TODO: what if there's more than one space between tags in the file?
				// Should be an easy fix, but just chasing MVP right now
				tagLine := strings.ReplaceAll(line, " ", ",")
				tagsInFile = strings.Split(tagLine, ",")
				break
			}
		}

		for _, gotTag := range tagsInFile {
			if anyTags {
				if containsStringValue(wantTags, gotTag) {
					matchingTags[filename] = append(matchingTags[filename], gotTag)
				}
			} else if allTags {
				fmt.Println("ERROR: -all tag search is not yet implemented")
				os.Exit(1)
			}
		}
	}

	fmt.Println(matchingTags)
	return matchingTags
}

func containsStringValue(slice []string, value string) bool {
	for _, e := range slice {
		if e == value {
			return true
		}
	}
	return false
}
