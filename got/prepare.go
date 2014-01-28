package got

import (
	"fmt"
	"log"
	"os"
	"os/exec"
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
	)
	
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
	
	cmd := exec.Command("ls")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	
	if templateName == "" {
		fmt.Println("")
		log.Println("Download template with")
		fmt.Println("wget", "http://brunettoziosi.eu/files/colors.tex", "-O " + outFileName)
		fmt.Println("")
		cmd := exec.Command("wget", "http://brunettoziosi.eu/files/colors.tex", "-O " + outFileName)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	} else {
		log.Println("Coping the template")
		
		cmd := exec.Command("cp", templateName, outFileName)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}
	
	tGlob1 := time.Now()
	fmt.Println()
	log.Println("Wall time for preparation ", tGlob1.Sub(tGlob0))
}