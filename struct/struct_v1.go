package main

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	FileRegex       = `%s/[\w\ \&\-\/]*\.(%s)`
	folderURI       = "UQ/ST3/AsianCommon"
	validExtensions = []string{"jpg", "png", "jpeg", "gif", "pdf", "mp4", "json"}
)

func main() {
	type test struct {
		status bool
	}
	fmt.Println("Hello! Welcome to  Sturct!")
	nwTest := new(test)
	nwTest.status = true
	//fmt.Println(nwTest)
	regX := regexp.MustCompile(fmt.Sprintf(FileRegex, folderURI, strings.Join(validExtensions, "|")))
	tg := "UQ/ST3/AsianCommons/imagesgoods/000976/red_shirt_2.jpgpng"
	key := regX.FindString(tg)
	//key := regX.MatchString(tg)
	fmt.Println(regX)
	fmt.Printf("%T, %v\n", key, key)
	fmt.Println(len(key))
	RegxContent := `^([\w-]+\/)*[\w-]+([\w-]+\.(%s))?$`
	regX2 := fmt.Sprintf(RegxContent, strings.Join(validExtensions, "|")) //^([\w-]+\/)*[\w-]+([\w-]+\.(jpg|png|jpeg|gif|pdf|mp4|json))?$
	fmt.Println(regX2)
}
