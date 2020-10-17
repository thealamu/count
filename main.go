package main

import (
	"fmt"
	"log"
	"os"
)

var usage = `Usage: count WORD FILE
Count occurrences of WORD in FILE
Text can also be piped in:
cat FILE | count WORD
`

//isPipe returns true if f is a pipe
func isPipe(f *os.File) bool {
	fi, err := f.Stat()
	if err != nil {
		log.Println(err)
		return false
	}
	return fi.Mode()&os.ModeNamedPipe != 0
}

func main() {
	params := os.Args[1:]
	if isPipe(os.Stdin) {
		countPiped(params)
		return
	}
	countFile(params)
}

func countFile(args []string) {
	if len(args) < 2 {
		exitWithUsage()
	}
	word, filename := args[0], args[1]
	fd, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	count, err := countOccurrences(fd, word)
	if err != nil {
		panic(err)
	}
	fmt.Println(count)
}

func countPiped(args []string) {
	if len(args) < 1 {
		exitWithUsage()
	}
	count, err := countOccurrences(os.Stdin, args[0])
	if err != nil {
		panic(err)
	}
	fmt.Println(count)
}

func exitWithUsage() {
	fmt.Print(usage)
	os.Exit(1)
}
