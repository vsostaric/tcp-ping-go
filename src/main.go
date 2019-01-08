package main

import (
	"./catcher"
	"./pitcher"
	"flag"
)

const (
	StopCharacter = "\r\n\r\n"
)

var catcherCmdArgs = map[string]bool{"c": true, "catch": true, "catcher": true}
var pitcherCmdArgs = map[string]bool{"p": true, "pitch": true, "pitcher": true}

func main() {

	var communicatorType string
	flag.StringVar(&communicatorType, "ctype", "", "Usage")
	flag.Parse()

	if catcherCmdArgs[communicatorType] {
		catcher.Catch()
	} else if pitcherCmdArgs[communicatorType] {
		pitcher.Pitch()
	}

}
