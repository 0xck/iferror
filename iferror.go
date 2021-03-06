package iferror

import (
	"fmt"
	"log"
	"os"
)

// HandleWith handles error with given functions
func HandleWith(f func(error), e error, fs ...func(error)) bool {
	if e != nil {
		f(e)
		for _, g := range fs {
			g(e)
		}
		return true
	}
	return false
}

// HandleWithErrors handles error with some other efunc functions
// it handles only functions without extra arguments
// other functions returns will be ignored
func HandleWithErrors(f func(error) bool, e error, fs ...func(error) bool) bool {
	if e != nil {
		f(e)
		for _, g := range fs {
			g(e)
		}
		return true
	}
	return false
}

// PrintError prints line with error
func PrintError(e error) bool {
	if e != nil {
		fmt.Println(e)
		return true
	}
	return false
}

// LogWith logs error with given logger
func LogWith(l *log.Logger, e error) bool {
	if e != nil {
		l.Println(e)
		return true
	}
	return false
}

// LogAndExit logs error and exit app
func LogAndExit(e error) bool {
	if e != nil {
		log.Fatalln(e)

	}
	return false
}

// ExitDueTo performs exit with optionally given code
// By default code is 1
func ExitDueTo(e error, c ...int) bool {
	if e != nil {
		code := 1
		if len(c) > 0 {
			code = c[0]
		}
		os.Exit(code)
	}
	return false
}

// Exit performs exit with code 1
func Exit(e error) bool {
	return ExitDueTo(e)
}

// Nothing just does nothing
func Nothing(error) bool { return false }

// Check checks error state and returns true or false
func Check(e error) bool {
	if e != nil {
		return true
	}
	return false
}
