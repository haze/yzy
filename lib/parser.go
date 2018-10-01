package lib

import (
	"io/ioutil"
	"strings"
)

///////////////
//  parser   //
// functions //
///////////////

// gets the raw contents of the parsers file
func (p *Parser) Content() string {
	dat, err := ioutil.ReadFile(string(p.FileMeta))
	if err != nil {
		QuickPanic(ElaborateError{
			Name:    ParserError,
			Message: err.Error(),
			Parser:  nil,
			Child:   err,
		})
	}
	return string(dat)
}

// returns all lines in the parsers file
func (p *Parser) Lines() []string {
	return strings.Split(p.Content(), "\n")
}

func (p *Parser) GetLine(ind int) string {
	return p.Lines()[ind]
}
