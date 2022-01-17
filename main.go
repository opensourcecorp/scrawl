package main

import (
	"bytes"
	"fmt"

	"github.com/yuin/goldmark"
)

func main() {
	source := []byte("# Main\n\n## Topic 1\n\nHey look, that's pre' neat.\n")

	var buf bytes.Buffer
	if err := goldmark.Convert(source, &buf); err != nil {
		panic(err)
	}

	fmt.Println(buf.String())
}
