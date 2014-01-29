package got

import (
	"fmt"
	"log"
	"path/filepath"
	"os"
	"os/user"
	"strings"
	"time"
)

func Publish (publishFolder string, pdfName string) () {
	
	var (
		u *user.User
		here string
		err error
		copied int64
		handoutsName string
		sharableName string
		pdfDest string
		handoutsDest string
		sharableDest string
		extension string
	)
	
	tGlob0 := time.Now()
	
	if publishFolder == "" {
		here, _ :=  os.Getwd()
		publishFolder =  filepath.Join(filepath.Dir(filepath.Dir(here)), "PresentationPrivate", here)
	}
	
	if pdfName == "" {
		u, err = user.Current()
		here, err = os.Getwd()	
		here = filepath.Base(here)
		pdfName = u.Username + "-" + here + ".pdf"
	}
	
	extension = filepath.Ext(pdfName)
	
	handoutsName = strings.TrimSuffix(texName, extension) + "-handouts.pdf" 
	sharableName = strings.TrimSuffix(texName, extension) + "-public.pdf"
	
	pdfDest = filepath.Join(publishFolder, here, pdfName)
	handoutsDest = filepath.Join(publishFolder, handoutsName)
	sharableDest = filepath.Join(publishFolder, sharableName)
	
	log.Println("Creating folder ", publishFolder)
	if err = os.MkdirAll(publishFolder, 0700); err != nil {
		log.Fatal("Can't create folder ", err)
	}
	
	log.Println("Copying presentation pdfs to ", publishFolder)
	
	if copied, err = CopyFile(pdfName, pdfDest); err != nil {
		log.Fatal("Can't copy ", pdfName, " to ", pdfDest, " with error ", err)
	} else {
		log.Println("Copies ", copied, " bytes")
	}
	
	if copied, err = CopyFile(handoutsName, handoutsDest); err != nil {
		log.Fatal("Can't copy ", handoutsName, " to ", handoutsDest, " with error ", err)
	} else {
		log.Println("Copies ", copied, " bytes")
	}
	
	if copied, err = CopyFile(sharableName, sharableDest); err != nil {
		log.Fatal("Can't copy ", sharableName, " to ", sharableDest, " with error ", err)
	} else {
		log.Println("Copies ", copied, " bytes")
	}
	
	tGlob1 := time.Now()
	fmt.Println()
	log.Println("Wall time for publishing ", tGlob1.Sub(tGlob0))
	
}