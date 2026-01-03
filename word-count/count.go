package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	var total_lines, total_words, total_chars uint32

	file, err := os.Open("hist.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')

		if err != nil && err != io.EOF {
			fmt.Println(err.Error())
			return
		}

		words := strings.Fields(line)
		for _, w := range words {
			total_chars += uint32(utf8.RuneCountInString(w))
		}

		total_words += uint32(len(words))
		total_lines++

		if err == io.EOF {
			break
		}
	}
	file.Close()
	fmt.Printf("Total Lines: %v\n", total_lines)
	fmt.Printf("Total Words: %v\n", total_words)
	fmt.Printf("Total Chars (Runes): %v\n", total_chars)
}
