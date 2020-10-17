package main

import (
	"github.com/matryer/is"
	"os"
	"testing"
)

func TestCount(t *testing.T) {
	is := is.New(t)
	fd, err := os.Open("testfixtures/testdata.txt")
	is.NoErr(err)

	count, err := countOccurrences(fd, "of")
	is.NoErr(err)
	is.Equal(7, count)
}