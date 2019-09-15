package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/yene/gohtml"
)

func main() {
	gohtml.Condense = true

	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) != 1 {
		log.Fatalln("Please provide path to the HTML files.")
	}

	source := argsWithoutProg[0]

	if _, err := os.Stat(source); err != nil {
		if os.IsNotExist(err) {
			log.Fatalln("Please provide path to the HTML files.")
		} else {
			log.Fatalln(err)
		}
	}

	subDirToSkip := "skip"

	err := filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() && info.Name() == subDirToSkip {
			log.Printf("skipping a dir without errors: %+v \n", info.Name())
			return filepath.SkipDir
		}
		if !info.IsDir() && filepath.Ext(path) == ".html" {
			err := tidyHTML(path)
			if err != nil {
				log.Println(err)
			}
		}
		return nil
	})
	if err != nil {
		log.Printf("error walking the path %q: %v\n", source, err)
		return
	}

}

func tidyHTML(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	res := gohtml.FormatBytes(data)
	var perm os.FileMode = 0644
	err = ioutil.WriteFile(filename, res, perm)
	if err != nil {
		return err
	}
	return nil
}
