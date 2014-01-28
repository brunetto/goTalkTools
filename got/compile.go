package got

import (
	"bitbucket.org/brunetto/goutils/readfile"
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
	"time"
)

func Compile (texName string) () {
	log.Println(texName)
	
	var (
		u *user.User
		here string
		err error
		texFile *os.File
		texReader *bufio.Reader
		handoutsName string
		handoutsFile *os.File
		handoutsWriter *bufio.Writer
		sharableFile *os.File
		sharableWriter *bufio.Writer
		sharableName string
		line string
		note1 string
		note2 string
		note3 string
		appendixSection bool = false
		extension string
	)
	
	tGlob0 := time.Now()
	
	if texName == "" {
		u, err = user.Current()
		here, err = os.Getwd()	
		here = filepath.Base(here)
		texName = u.Username + "-" + here + ".tex"
	}
	
	if texFile, err = os.Open(texName); err != nil {log.Fatal(err)}
	defer texFile.Close()
	texReader = bufio.NewReader(texFile)
	
	log.Println("Creating files for handouts and public presentation version:")
	
	extension = filepath.Ext(texName)
	
	handoutsName = strings.TrimSuffix(texName, extension) + "-handouts.tex" 
	sharableName = strings.TrimSuffix(texName, extension) + "-public.tex"
	
	fmt.Println(handoutsName)
	fmt.Println(sharableName)
	
	if handoutsFile, err = os.Create(handoutsName); err != nil {log.Fatal(err)}
	defer handoutsFile.Close()
	handoutsWriter = bufio.NewWriter(handoutsFile)

	if sharableFile, err = os.Create(sharableName); err != nil {log.Fatal(err)}
	defer sharableFile.Close()
	sharableWriter = bufio.NewWriter(sharableFile)
	
	log.Println("Preparing contents")
	
	// Strings for handouts check
	
	note1 = `\setbeameroption{show notes}`
	note2 = `\usepackage{pgfpages}`
	note3 = `\pgfpagesuselayout{8 on 1}[a4paper]%, landscape]`

	// Read the presentation .tex file and copy it to the handout and sharable version 
	for {
		// Read line by line
		if line, err = readfile.Readln(texReader); err != nil {
			if err.Error() != "EOF" {
				log.Fatal("Non EOF error while reading ", err)
			} else {
				break
			}
		}

		// Handouts activation?
		if strings.Contains(line, note1) || strings.Contains(line, note2) || strings.Contains(line, note3) {
			if _, err = handoutsWriter.WriteString(strings.TrimLeft(line, "%") + "\n"); err != nil {log.Fatal(err)}
		} else {
			if _, err = handoutsWriter.WriteString(line + "\n"); err != nil {log.Fatal(err)}
		}
		
		// Are we in the appendix?
		if !appendixSection {
			// If appendixSection is false, check if we arrived there
			appendixSection = strings.Contains(line, `\appendix`)
			if _, err = sharableWriter.WriteString(line + "\n"); err != nil {log.Fatal(err)}
		} else {
			// We already are in the appendix and appendixSection is true
			// and we don't want it to change in false again
		}
		
	}
	
	// Compile .tex files
	// Presentation one
	talkCmd := exec.Command("pdflatex", texName)
	talkCmd.Stdout = os.Stdout
	talkCmd.Stderr = os.Stderr
	talkCmd.Run()
	talkCmd.Run()
	talkCmd.Run()
	
	// Handouts version
	handoutCmd := exec.Command("pdflatex", handoutsName)
	handoutCmd.Stdout = os.Stdout
	handoutCmd.Stderr = os.Stderr
	handoutCmd.Run()
	handoutCmd.Run()
	handoutCmd.Run()
	
	// Sharable version
	sharableCmd := exec.Command("pdflatex", sharableName)
	sharableCmd.Stdout = os.Stdout
	sharableCmd.Stderr = os.Stderr
	sharableCmd.Run()
	sharableCmd.Run()
	sharableCmd.Run()

	tGlob1 := time.Now()
	fmt.Println()
	log.Println("Wall time for compiling ", tGlob1.Sub(tGlob0))
}