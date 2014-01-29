package main

import (
	"./got"
	"fmt"
	"log"
	"time"
)

// This program take a .tex beamer presentation and
// produces as output three different presentation:
// 1 - normal presentation with appendix
// 2 - presentation without appendix for sharing
// 3 - handout notes
// .tex file need to be in the brunetto's format
func main () {
	tGlob0 := time.Now()
	
	got.InitCommands()
	got.GoTalkCmd.Execute()
	
	tGlob1 := time.Now()
	fmt.Println()
	log.Println("Wall time for compiling ", tGlob1.Sub(tGlob0))
}