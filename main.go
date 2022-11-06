package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	f, err := os.Open(".")
	if err != nil {
		fmt.Println(err)
		return
	}
	files, err := f.ReadDir(0)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range files {
		if !v.IsDir() {
			fmt.Println(v.Name())
			r := regexp.MustCompile(`(?i).*\.(epub|pdf)$`)
			directoryEntry := r.FindAllString(v.Name(), -1)
			if len(directoryEntry) > 0 && len(directoryEntry[0]) > 0 {
				log.Printf("found: %s", directoryEntry[0])

				var newDirectoryName string = v.Name()[:len(v.Name())-len(filepath.Ext(v.Name()))]
				log.Printf("Neues Verzeichnis: %v", newDirectoryName)

				if err := os.Mkdir(newDirectoryName, os.ModePerm); err != nil {
					log.Fatal(fmt.Sprintf("Verzeichnis kann nicht angelegt werden: %v", err.Error()))
				} else {
					originalPath := filepath.Join(f.Name(), v.Name())
					newPath := filepath.Join(f.Name(), newDirectoryName, v.Name())
					log.Printf("original:%s, new:%s", originalPath, newPath)
					os.Rename(originalPath, newPath)
				}
			}
		}
	}
}
