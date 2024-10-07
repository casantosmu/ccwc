package main_test

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func runBinary(args []string) ([]byte, error) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("could not get current dir: %v", err)
	}
	binaryPath := filepath.Join(dir, "ccwc")
	cmd := exec.Command(binaryPath, args...)
	return cmd.CombinedOutput()
}

func TestCLi(t *testing.T) {
	t.Run("-l flag", func(t *testing.T) {
		output, err := runBinary([]string{"-l", "test.txt"})
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
		output, err := runBinary([]string{"-w", "test.txt"})
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
		output, err := runBinary([]string{"-m", "test.txt"})
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
		output, err := runBinary([]string{"-c", "test.txt"})
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
		output, err := runBinary([]string{"test.txt"})
		if err != nil {
			t.Fatal(err)
		}

		actual := string(output)
		expected := "7145 58164 342190 test.txt\n"

		if actual != expected {
			t.Fatalf("actual = %s, expected = %s", actual, expected)
		}
	})

	t.Run("-lwmc flags", func(t *testing.T) {
		output, err := runBinary([]string{"-lwmc", "test.txt"})
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
