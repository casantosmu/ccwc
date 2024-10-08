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
	filename   string
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

func newConfig(filename string, countLines, countWord, countChars, countBytes bool) config {
	if !countLines && !countWord && !countChars && !countBytes {
		return config{
			filename:   filename,
			countLines: true,
			countWord:  true,
			countBytes: true,
		}
	}

	return config{
		filename:   filename,
		countLines: countLines,
		countWord:  countWord,
		countChars: countChars,
		countBytes: countBytes,
	}
}

func count(r io.Reader, conf config) (counters, error) {
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

func print(counters counters, conf config) {
	parts := []string{}

	if conf.countLines {
		parts = append(parts, fmt.Sprint(counters.lineCount))
	}
	if conf.countWord {
		parts = append(parts, fmt.Sprint(counters.wordCount))
	}
	if conf.countChars {
		parts = append(parts, fmt.Sprint(counters.charCount))
	}
	if conf.countBytes {
		parts = append(parts, fmt.Sprint(counters.byteCount))
	}
	if conf.filename != "" {
		parts = append(parts, conf.filename)
	}

	output := strings.Join(parts, " ")
	log.Printf("%s\n", output)
}

func main() {
	log.SetFlags(0)

	countLines := flag.BoolP("lines", "l", false, "print the newline counts")
	countWords := flag.BoolP("words", "w", false, "print the word counts")
	countChars := flag.BoolP("chars", "m", false, "print the character counts")
	countBytes := flag.BoolP("bytes", "c", false, "print the byte counts")
	flag.Parse()

	filename := flag.Arg(0)

	var r io.Reader
	if filename == "" {
		r = os.Stdin
	} else {
		fi, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer fi.Close()
		r = fi
	}

	conf := newConfig(filename, *countLines, *countWords, *countChars, *countBytes)

	counters, err := count(r, conf)
	if err != nil {
		log.Fatal(err)
	}

	print(counters, conf)
}
