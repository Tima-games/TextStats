package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	var reader *bufio.Reader

	if len(os.Args) > 1 {
		file, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()
		reader = bufio.NewReader(file)
	} else {
		reader = bufio.NewReader(os.Stdin)
	}

	length := 0
	space := 0
	lines := 0
	wordCount := 0

	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Println("Error reading line:", err)
			break
		}
		lines++
		line = strings.TrimSpace(line)

		words := strings.Fields(line)
		wordCount += len(words)

		for _, r := range line {
			if r == ' ' {
				space++
			} else {
				length++
			}
		}
		if err == io.EOF {
			break
		}
	}
	fmt.Println("In your string(s)", lines, "lines")
	fmt.Println("In your string(s)", wordCount, "words")
	fmt.Println("In your string(s)", space, "spaces")
	fmt.Println("In your string(s)", length, "letters")
}
