package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	helpF := flag.Bool("help", false, "Show help")
	h := flag.Bool("h", false, "Show help (short)")
	linesF := flag.Bool("lines", false, "Show lines")
	l := flag.Bool("l", false, "Show lines (short)")
	wordsF := flag.Bool("words", false, "Show words")
	w := flag.Bool("w", false, "Show words (short)")
	lettersF := flag.Bool("letters", false, "Show letters")
	c := flag.Bool("c", false, "Show letters (short)")
	spacesF := flag.Bool("spaces", false, "Show spaces")
	s := flag.Bool("s", false, "Show spaces (short)")
	versionF := flag.Bool("version", false, "Show version")
	v := flag.Bool("v", false, "Show version (short)")

	flag.Parse()

	showLines := *linesF || *l
	showWords := *wordsF || *w
	showLetters := *lettersF || *c
	showSpaces := *spacesF || *s

	if *versionF || *v {
		fmt.Println("TextStats v1.5.0 (11-04-26 release)")
		return
	}

	if *helpF || *h {
		fmt.Println("TextStats v1.5 - Counts lines, words, spaces, letters and other more")
		fmt.Println()
		fmt.Println("Usage: ./textstats-[your version/arch] [Flag(s)] [File]")
		fmt.Print(`Flags:
        
        -h, --help           Shows this help message and exit
        -v, --version        Shows version and exit
        -l, --lines          Shows only lines count
        -w, --words          Shows only words count
        -s, --spaces         Shows only spaces count
        -c, --letters        Shows only letters count
        `)
		return
	}

	args := flag.Args()

	var reader *bufio.Reader

	if len(args) > 0 {
		file, err := os.Open(args[0])
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()
		reader = bufio.NewReader(file)
	} else {
		reader = bufio.NewReader(os.Stdin)
	}

	lettersCount := 0
	spaceCount := 0
	linesCount := 0
	wordCount := 0

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			line = strings.TrimSpace(line)
			if len(line) > 0 {
			}
			break
		}
		if err != nil {
			fmt.Println("Error reading line:", err)
			break
		}
		linesCount++
		line = strings.TrimSpace(line)

		words := strings.Fields(line)
		wordCount += len(words)

		for _, r := range line {
			if r == ' ' {
				spaceCount++
			} else {
				lettersCount++
			}
		}
	}
	if !showLines && !showWords && !showLetters && !showSpaces {
		fmt.Println("In your string(s)", linesCount, "lines")
		fmt.Println("In your string(s)", wordCount, "words")
		fmt.Println("In your string(s)", spaceCount, "spaces")
		fmt.Println("In your string(s)", lettersCount, "letters")
	}
	if showLines {
		fmt.Println("In your string(s)", linesCount, "lines")
	}
	if showWords {
		fmt.Println("In your string(s)", wordCount, "words")
	}
	if showLetters {
		fmt.Println("In your string(s)", lettersCount, "letters")
	}
	if showSpaces {
		fmt.Println("In your string(s)", spaceCount, "spaces")
	}
}
