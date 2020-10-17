package main

import (
	"bufio"
	"io"
	"strings"
)

var punctuation = `.,"'`

// count occurrences of word in r
func countOccurrences(r io.Reader, word string) (int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		if strings.Trim(scanner.Text(), punctuation) == word {
			count++
		}
	}
	if err := scanner.Err(); err != nil {
		return -1, err
	}
	return count, nil
}

