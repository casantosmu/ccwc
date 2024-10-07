package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode/utf8"

	flag "github.com/spf13/pflag"
)

type config struct {
	countLines bool
	countWord  bool
	countChars bool
	countBytes bool
}

type counters struct {
	lineCount int
	wordCount int
	charCount int
	byteCount int
}

func newConfig(countLines, countWord, countChars, countBytes bool) config {
	if !countLines && !countWord && !countChars && !countBytes {
		return config{
			countLines: true,
			countWord:  true,
			countBytes: true,
		}
	}

	return config{
		countLines,
		countWord,
		countChars,
		countBytes,
	}
}

func counter(r io.Reader, conf config) (counters, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return counters{}, err
	}

	c := counters{}
	if conf.countLines {
		c.lineCount = bytes.Count(b, []byte{'\n'})
	}
	if conf.countWord {
		words := strings.Fields(string(b))
		c.wordCount = len(words)
	}
	if conf.countChars {
		c.charCount = utf8.RuneCount(b)
	}
	if conf.countBytes {
		c.byteCount = len(b)
	}

	return c, nil
}

func print(filename string, c counters, conf config) {
	output := ""
	if conf.countLines {
		output += fmt.Sprintf("%d ", c.lineCount)
	}
	if conf.countWord {
		output += fmt.Sprintf("%d ", c.wordCount)
	}
	if conf.countChars {
		output += fmt.Sprintf("%d ", c.charCount)
	}
	if conf.countBytes {
		output += fmt.Sprintf("%d ", c.byteCount)
	}

	log.Printf("%s%s\n", output, filename)
}

func main() {
	log.SetFlags(0)

	countLines := flag.BoolP("lines", "l", false, "print the newline counts")
	countWords := flag.BoolP("words", "w", false, "print the word counts")
	countChars := flag.BoolP("chars", "m", false, "print the character counts")
	countBytes := flag.BoolP("bytes", "c", false, "print the byte counts")
	flag.Parse()

	filename := flag.Arg(0)

	fi, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer fi.Close()

	conf := newConfig(*countLines, *countWords, *countChars, *countBytes)

	c, err := counter(fi, conf)
	if err != nil {
		log.Fatal(err)
	}

	print(filename, c, conf)
}
