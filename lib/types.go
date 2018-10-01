package lib

import "path/filepath"

//////////////////////////////
// MetaRuntime              //
// - keeps information      //
//   about the syntax tree, //
//   file, cursor etc       //
//////////////////////////////
type MetaRuntime struct {
	ParserMap map[string]Pointer
	Entrance  string // file w/ entry point
}

type Pointer struct {
	Line, Column int
}

type File string

func (f File) DisplayName() string {
	return filepath.Base(string(f))
}

type Parser struct {
	Cursor   Pointer
	FileMeta File
	Meta     *MetaRuntime
}
