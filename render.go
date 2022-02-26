package main

// TODO: this will render the scrawl notes from Markdown to HTML, for whatever
// you want but also to be viewed by `scrawl web`

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer/html"
)

func render() {
	scrawldir := getScrawldir()
	webDir := filepath.Join(scrawldir, "web")
	if err := os.MkdirAll(webDir, 0755); err != nil {
		log.Fatalf("could not create rendered HTML directory (%s)\n%v", webDir, err)
	}

	files, err := os.ReadDir(scrawldir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() || !strings.Contains(file.Name(), ".md") {
			continue
		}

		filename := file.Name()
		notePath := filepath.Join(scrawldir, filename)

		contentsRaw, err := os.ReadFile(notePath)
		if err != nil {
			log.Fatalf("error reading file %s\n%v", file.Name(), err)
		}

		// Build the Markdown header as the note's date from the filename (which is itself a timestamp)
		header := filename
		header = strings.TrimSuffix(header, ".md")
		headerSlice := strings.Split(header, "T") // filenames look similar to RFC3339 time
		date := headerSlice[0]
		hms := strings.ReplaceAll(headerSlice[1], "-", ":")
		header = fmt.Sprintf("# %s %s", date, hms)

		contents := fmt.Sprintf(
			"<style>\n%s\n</style>\n%s\n\n%s\n",
			getCSSContent("notes"),
			header,
			string(contentsRaw),
		)

		// Now, convert the Markdown to HTML, get the remaining filesystem info
		// you need, and write it all out
		md := []byte(contents)

		// By constructing a new goldmark converter (instead of converting
		// directly via a top-level Convert() calls), we can specify parsing &
		// conversion options
		gm := goldmark.New(
			goldmark.WithRendererOptions(
				// WithUnsafe lets scrawl use CSS styling & other raw HTML
				// elements inline
				html.WithUnsafe(),
			),
		)

		var buf bytes.Buffer
		if err := gm.Convert(md, &buf); err != nil {
			log.Fatalf("could not convert Markdown input from %s to HTML\n%v", notePath, err)
		}

		outFilename := strings.TrimSuffix(filename, ".md") + ".html"
		outPath := filepath.Join(webDir, outFilename)

		fileInfo, err := os.Stat(notePath)
		if err != nil {
			log.Fatalf("could not determine mode/perms from input file %s\n%v", notePath, err)
		}
		fileMode := fileInfo.Mode()

		if err := os.WriteFile(outPath, buf.Bytes(), fileMode); err != nil {
			log.Fatalf("could not output rendered HTML file to %s\n%v", outPath, err)
		}
	}
}

func getCSSContent(component string) string {
	cssPath := "web/css/" + component + ".css"
	css, err := os.ReadFile(cssPath)
	if err != nil {
		log.Fatalf("could not read CSS file from %s\n%v", cssPath, err)
	}
	return string(css)
}
