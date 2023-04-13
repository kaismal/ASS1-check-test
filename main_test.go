package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number!"},
		{"not prime", 8, false, "8 is not a prime number because it is divisible by 2!"},
		{"zero", 0, false, "0 is not prime, by definition!"},
		{"one", 1, false, "1 is not prime, by definition!"},
		{"negative number", -11, false, "Negative numbers are not prime, by definition!"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s: expected true but got false", e.name)
		}

		if !e.expected && result {
			t.Errorf("%s: expected false but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s: expected %s but got %s", e.name, e.msg, msg)
		}
	}
}

func Test_checkNumbers(t *testing.T) {
	checkNumbersTests := []struct {
		name       string
		testString string
		expected   bool
		msg        string
	}{
		{"quitting", "q", true, ""},
		{"not a number", "123test", false, "Please enter a whole number!"},
		{"out of the range", "9223372036854775808", false, "Please enter a whole number!"},
		{"empty", "", false, "Please enter a whole number!"},
		{"whitespaces", "   ", false, "Please enter a whole number!"},
		{"prime", "7", false, "7 is a prime number!"},
		{"not prime", "8", false, "8 is not a prime number because it is divisible by 2!"},
		{"zero", "0", false, "0 is not prime, by definition!"},
		{"one", "1", false, "1 is not prime, by definition!"},
		{"negative number", "-11", false, "Negative numbers are not prime, by definition!"},
	}

	for _, e := range checkNumbersTests {
		scanner := bufio.NewScanner(strings.NewReader(e.testString))
		msg, res := checkNumbers(scanner)
		if e.expected && !res {
			t.Errorf("%s: expected true but got false", e.name)
		}

		if !e.expected && res {
			t.Errorf("%s: expected false but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s: expected %s but got %s", e.name, e.msg, msg)
		}
	}
}

func Test_intro(t *testing.T) {
	expectedOutput := "Is it Prime?\n------------\nEnter a whole number, and we'll tell you if it is a prime number or not. Enter q to quit.\n-> "

	last := os.Stdout
	reader, tempStdout, _ := os.Pipe()

	os.Stdout = tempStdout
	intro()

	tempStdout.Close()

	res, _ := io.ReadAll(reader)
	output := string(res)

	if output != expectedOutput {
		t.Errorf("intro: expected %s but got %s", expectedOutput, res)
	}

	os.Stdout = last
}

func Test_readUserInput(t *testing.T) {
	doneChan := make(chan bool)

	var stdin bytes.Buffer

	stdin.Write([]byte("1\nq\n"))

	go readUserInput(&stdin, doneChan)
	<-doneChan
	close(doneChan)
}
