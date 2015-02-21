package main

import (
	"fmt"
	"os"
	"regexp"
	"flag"
	"strings"
	"path/filepath"
)

var localefile = flag.String("localefile", "", "A path to one of your localization files (ie: /path/to/your/en.json)")
var viewfolder = flag.String("viewfolder", "", "A path to your view folder (ie: /path/to/your/viewfolder)")
var L10nKeys []string
var unusedlocale []string
var viewKeys []string
var viewstrings []string

func init() {
	flag.StringVar(localefile, "L", "", "Description")
	flag.StringVar(viewfolder, "V", "", "Description")
}

func main() {
	flag.Parse()
	L10nKeys  = append(getLocaleKeys(*localefile))
	filepath.Walk(*viewfolder,walkpath)
	listUnusedKeys(unusedlocale)
}

func listUnusedKeys(keys []string){
	for i := 0; i < len(keys); i++ {
		fmt.Println(keys[i])
	}
}

func walkpath(path string, f os.FileInfo, err error) error {

	if f.Size() > 0{
		viewKeys = append(getViewKeys(path))

		if len(viewKeys) > 0 {
			unusedlocale = append(compareKeys(L10nKeys, viewKeys))
		}
	}
	return nil
}

func getLocaleKeys(path string) []string {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil{
		fmt.Println(err)
	}

	data := make([]byte, stat.Size())
	_, err = file.Read(data)
	if err != nil {
		fmt.Println(err)
	}

	datastr := string(data)
	keysCleaned, err := regexp.Compile(`.+\:`)
	var localearray []string

	locales := keysCleaned.FindAllString(datastr,-1)
	for i := 0; i < len(locales); i++ {
		localestr := strings.Replace(locales[i],`:`,``, -1)
		matchlocales, _ := regexp.Compile("'(.*?)'")
		locale := matchlocales.FindAllString(localestr,-1)
		localearray = append(localearray,locale[0])
	}
	return localearray
}

func getViewKeys(path string) []string{
	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
	}

	data := make([]byte, stat.Size())

	_, err = file.Read(data)
	if err != nil {
		fmt.Println(err)
	}

	datastr := string(data)
	keysCleaned, _ := regexp.Compile("#{__(.*?)}")
	viewvars := keysCleaned.FindAllString(datastr,-1)

	var viewvar []string

	for i := 0; i < len(viewvars); i++ {
		localevar := strings.Replace(viewvars[i], `"`, `'`, -1)
		keysCleaned, _ := regexp.Compile("'(.*?)'")
		viewvar = append(keysCleaned.FindAllString(localevar,-1))
		viewstrings = append(viewstrings,viewvar[0])
	}
	return viewstrings

}

func compareKeys(localeKeys, viewKeys []string) []string {
	keymap := make(map[string]int)

	for _, key := range viewKeys {
		keymap[key]++
	}

	var unusedkey []string
	for _, index := range localeKeys {

		if keymap[index] > 0 {
			keymap[index]--
			continue
		}
		unusedkey = append(unusedkey, index)
	}
	return unusedkey
}
