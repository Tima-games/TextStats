package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	length := 0
	space := 0
	lines := 0
	wordc := 0

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		lines++
		line = strings.TrimSpace(line)

		words := strings.Fields(line)
		wordc += len(words)

		for _, r := range line {
			if r == ' ' {
				space++
			} else {
				length++
			}
		}
	}
	fmt.Println("In your string(s)", lines, "lines")
	fmt.Println("In your string(s)", wordc, "words")
	fmt.Println("In your string(s)", space, "spaces")
	fmt.Println("In your string(s)", length, "letters")
}
