package lib

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

const (
	EmptyFile    = "Empty Input File"
	EmptyFileMsg = "Expected an input file, got none"
	ParserError  = "Parser Error"
)

type ElaborateError struct {
	Name, Message string
	Parser        *Parser
	Child         error
}

func (e ElaborateError) Headline() string {
	white := color.New(color.FgHiWhite).SprintFunc()
	blue := color.New(color.FgHiBlack).SprintfFunc()
	return fmt.Sprintf("%s ~ %s [%s:%s]",
		white(e.Name),
		white(e.Parser.FileMeta.DisplayName()),
		blue("%d", e.Parser.Cursor.Line),
		blue("%d", e.Parser.Cursor.Column),
	)
}

func generateLines(errorLine int) []int {
	return []int{errorLine - 1, errorLine, errorLine + 1}
}

// ease of life function to quickly return
func QuickPanic(err ElaborateError) {
	if err.Child != nil || err.Parser == nil {
		// nil parser, this was in the creation of the parsers
		// or non-nil child
		fmt.Println(err.Message)
	} else {
		fmt.Println(err.Print())
	}
	os.Exit(1)
}

func (e ElaborateError) FilePreview() string {
	buf := ""
	lines := generateLines(e.Parser.Cursor.Line)
	for i := range lines {
		buf += fmt.Sprintf("%d |\t%s\n", i, e.Parser.GetLine(i))
	}
	return buf
}

func (e ElaborateError) Print() string {
	return fmt.Sprintf("%s\n%s\n",
		e.Headline(),
		e.FilePreview(),
	)
}
