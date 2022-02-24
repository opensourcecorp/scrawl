package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type tagMap map[string][]string

func searchTags(anyTags, allTags, untagged bool, wantTagsArg string) tagMap {
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
	matchingTags := make(tagMap)
	for _, filename := range filenames {
		contentsRaw, err := os.ReadFile(filename)
		if err != nil {
			fmt.Printf("ERROR: error while reading file %s\n", filename)
		}

		contents := strings.Split(string(contentsRaw), "\n")
		tagsInFile := getTagsFromFileContents(contents)

		// If searching for untagged notes, record the filename and skip to the next one
		if untagged {
			if len(tagsInFile) == 0 {
				matchingTags[filename] = []string{}
				continue
			}
		}

		for _, gotTag := range tagsInFile {
			if anyTags {
				if containsStringValue(wantTags, gotTag) {
					// matchingTags[filename] = append(matchingTags[filename], gotTag) // this will only store the tag it found, not all tags in the file
					matchingTags[filename] = tagsInFile
					break
				}
			} else if allTags {
				fmt.Println("ERROR: -all tag search is not yet implemented")
				os.Exit(1)
			}
		}
	}

	return matchingTags
}

func getTagsFromFileContents(contents []string) []string {
	var tags []string
	for _, line := range contents {
		tagPattern := regexp.MustCompile("^tags:.*$")
		if tagPattern.MatchString(line) {
			// TODO: what if there's more than one space between tags in the file?
			// Should be an easy fix, but just chasing MVP right now
			tagLine := line
			tagLine = strings.ReplaceAll(tagLine, "tags: ", "")
			tagLine = strings.ReplaceAll(tagLine, " ", "")
			tags = strings.Split(tagLine, ",")
			break
		}
	}
	return tags
}
