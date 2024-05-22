package models

type Flags struct {
	Verbose bool
	Quiet   bool
	Source  string
	Branch  string
}

var NukeFlags Flags
var RecommitFlags Flags
var OwnFlags Flags
var BlameFlags Flags
var RetimeFlags Flags
