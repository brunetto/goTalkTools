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
		absHere string
		err error
		copied int64
		handoutsName string
		sharableName string
		pdfSource string
		handoutsSource string
		sharableSource string
		pdfDest string
		handoutsDest string
		sharableDest string
		extension string
	)
	
	tGlob0 := time.Now()
	
	absHere, err = os.Getwd()	
	here = filepath.Base(absHere)
	
	if publishFolder == "" {
		publishFolder =  filepath.Join(filepath.Dir(filepath.Dir(absHere)), "PresentationsPublic", here)
	}
	
	if pdfName == "" {
		u, err = user.Current()
		pdfName = u.Username + "-" + here + ".pdf"
	}
	
	extension = filepath.Ext(pdfName)
	
	handoutsName = strings.TrimSuffix(pdfName, extension) + "-handouts.pdf" 
	sharableName = strings.TrimSuffix(pdfName, extension) + "-public.pdf"
	
	pdfSource = filepath.Join(absHere, pdfName)
	handoutsSource = filepath.Join(absHere, handoutsName)
	sharableSource = filepath.Join(absHere, sharableName)
	
	pdfDest = filepath.Join(publishFolder, pdfName)
	handoutsDest = filepath.Join(publishFolder, handoutsName)
	sharableDest = filepath.Join(publishFolder, sharableName)
	
	log.Println("Creating folder ", publishFolder)
	if err = os.MkdirAll(publishFolder, 0700); err != nil {
		log.Fatal("Can't create folder ", err)
	}
	
	log.Println("Copying presentation pdfs to ", publishFolder)
	
	if copied, err = CopyFile(pdfSource, pdfDest); err != nil {
		log.Fatal("Can't copy ", pdfSource, " to ", pdfDest, " with error ", err)
	} else {
		log.Println("Copies ", copied, " bytes")
	}
	
	if copied, err = CopyFile(handoutsSource, handoutsDest); err != nil {
		log.Fatal("Can't copy ", handoutsSource, " to ", handoutsDest, " with error ", err)
	} else {
		log.Println("Copies ", copied, " bytes")
	}
	
	if copied, err = CopyFile(sharableSource, sharableDest); err != nil {
		log.Fatal("Can't copy ", sharableSource, " to ", sharableDest, " with error ", err)
	} else {
		log.Println("Copies ", copied, " bytes")
	}
	
	tGlob1 := time.Now()
	fmt.Println()
	log.Println("Wall time for publishing ", tGlob1.Sub(tGlob0))
	
}