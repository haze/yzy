package main

import (
	"fmt"
	. "github.com/hvze/yzy/lib"
	"os"
	"strings"
)

const (
	MAJOR = 0
	MINOR = 1
)

func main() {
	argv := os.Args[1:]
	if len(argv) != 1 {
		PrintHelp()
	} else {
		file := argv[0]
		if strings.TrimSpace(file) == "" {
			QuickPanic(ElaborateError{
				Name:    EmptyFile,
				Message: EmptyFileMsg,
				Parser:  nil,
				Child:   nil,
			})
		}
	}
}
func PrintHelp() {
	fmt.Printf("yzy - v%d.%d\nauthor: haze <me@haze.sh>\nusage:\n\t./yzy [opts] <file>",
		MAJOR, MINOR)
}
