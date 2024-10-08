package main_test

import (
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func runBinary(args []string, stdin io.Reader) ([]byte, error) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("could not get current dir: %v", err)
	}
	binaryPath := filepath.Join(dir, "ccwc")
	cmd := exec.Command(binaryPath, args...)

	if stdin != nil {
		cmd.Stdin = stdin
	}

	return cmd.CombinedOutput()
}

func TestCLI(t *testing.T) {
	t.Run("-l flag", func(t *testing.T) {
		output, err := runBinary([]string{"-l", "test.txt"}, nil)
		if err != nil {
			t.Fatal(err)
		}

		actual := string(output)
		expected := "7145 test.txt\n"

		if actual != expected {
			t.Fatalf("actual = %s, expected = %s", actual, expected)
		}
	})

	t.Run("-w flag", func(t *testing.T) {
		output, err := runBinary([]string{"-w", "test.txt"}, nil)
		if err != nil {
			t.Fatal(err)
		}

		actual := string(output)
		expected := "58164 test.txt\n"

		if actual != expected {
			t.Fatalf("actual = %s, expected = %s", actual, expected)
		}
	})

	t.Run("-m flag", func(t *testing.T) {
		output, err := runBinary([]string{"-m", "test.txt"}, nil)
		if err != nil {
			t.Fatal(err)
		}

		actual := string(output)
		expected := "339292 test.txt\n"

		if actual != expected {
			t.Fatalf("actual = %s, expected = %s", actual, expected)
		}
	})

	t.Run("-c flag", func(t *testing.T) {
		output, err := runBinary([]string{"-c", "test.txt"}, nil)
		if err != nil {
			t.Fatal(err)
		}

		actual := string(output)
		expected := "342190 test.txt\n"

		if actual != expected {
			t.Fatalf("actual = %s, expected = %s", actual, expected)
		}
	})

	t.Run("no flags", func(t *testing.T) {
		output, err := runBinary([]string{"test.txt"}, nil)
		if err != nil {
			t.Fatal(err)
		}

		actual := string(output)
		expected := "7145 58164 342190 test.txt\n"

		if actual != expected {
			t.Fatalf("actual = %s, expected = %s", actual, expected)
		}
	})

	t.Run("read from standard input", func(t *testing.T) {
		fi, err := os.Open("test.txt")
		if err != nil {
			t.Fatal(err)
		}
		defer fi.Close()
		output, err := runBinary([]string{}, fi)
		if err != nil {
			t.Fatal(err)
		}

		actual := string(output)
		expected := "7145 58164 342190\n"

		if actual != expected {
			t.Fatalf("actual = %s, expected = %s", actual, expected)
		}
	})

	t.Run("-lwmc flags", func(t *testing.T) {
		output, err := runBinary([]string{"-lwmc", "test.txt"}, nil)
		if err != nil {
			t.Fatal(err)
		}

		actual := string(output)
		expected := "7145 58164 339292 342190 test.txt\n"

		if actual != expected {
			t.Fatalf("actual = %s, expected = %s", actual, expected)
		}
	})
}
