package got

import (
	"fmt"
	"io"
	"net/http"
	"log"
	"os"
// 	"os/exec"
	"os/user"
	"path/filepath"
	"time"
)

func Prepare (presentationName string, templateName string) () {
	
	var (
		u *user.User
		here string
		err error
		folderName string
		outFileName string
		outFile *os.File
		response *http.Response
		writtenBytes int64
	)
	
	const url = `http://brunettoziosi.eu/files/beamerTmpl.tex`
	
	tGlob0 := time.Now()
	
	if presentationName == "" {
		presentationName = time.Now().String()[:10] + "-location-event"
	}
	log.Println("Set presentation name to:")
	fmt.Println(presentationName)
	
	u, err = user.Current()
	here, err = os.Getwd()
	folderName = filepath.Join(here, presentationName)
	outFileName = filepath.Join(here, presentationName, u.Username+"-"+presentationName+".tex")
	
	log.Println("Creating folder ", folderName)
	if err = os.MkdirAll(folderName, 0700); err != nil {
		log.Fatal("Can't create folder ", err)
	}
	
	if templateName == "" {
		fmt.Println("")
		log.Println("Download template from ", url)
		fmt.Println("")

		// from https://github.com/thbar/golang-playground/blob/master/download-files.go
		if outFile, err = os.Create(outFileName); err != nil {
			log.Fatal("Error while creating", outFileName, "-", err)
        }
        defer outFile.Close()
		
		if response, err = http.Get(url); err != nil {
            log.Fatal("Error while downloading", url, "-", err)
        }
        defer response.Body.Close()
		
		if writtenBytes, err = io.Copy(outFile, response.Body); err != nil {
            log.Fatal("Error while downloading", url, "-", err)
        }
        
        log.Println("Downloaded ", writtenBytes, " bytes template")
		
	} else {
		log.Println("Coping the template")
		
// 		cmd := exec.Command("cp", templateName, outFileName)
// 		cmd.Stdout = os.Stdout
// 		cmd.Stderr = os.Stderr
// 		cmd.Run()
		if writtenBytes, err = CopyFile(templateName, outFileName); err != nil {
			log.Fatal("Can't copy ", templateName, " to ", outFileName, " with error ", err)
		} else {
			log.Println("Copies ", writtenBytes, " bytes")
		}
	}
	
	tGlob1 := time.Now()
	fmt.Println()
	log.Println("Wall time for preparation ", tGlob1.Sub(tGlob0))
}