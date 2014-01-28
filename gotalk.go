package main

import (
	"./got"
)

// This program take a .tex beamer presentation and
// produces as output three different presentation:
// 1 - normal presentation with appendix
// 2 - presentation without appendix for sharing
// 3 - handout notes
// .tex file need to be in the brunetto's format
func main () {
	got.InitCommands()
	got.GoTalkCmd.Execute()
	
}