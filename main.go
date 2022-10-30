package main

import (
	"fmt"
	"log"
	"os"
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
			r, regexErr := regexp.Compile("(.*)(?=\\.(epub|pdf)$)")
			if regexErr != nil {
				log.Panic(regexErr)
			}

			log.Printf("found: %s", r.FindAllString(v.Name(), -1))

			if err := os.Mkdir(r.FindString(v.Name()), os.ModePerm); err != nil {
				log.Fatal(err)
			} else {
				log.Printf("original:%s/%s, new:%s/%s/%s", f.Name(), v.Name(), f.Name(), v.Name(), v.Name())
				//os.Rename(f.Name()+v.Name(), "./"+v.Name()+"/"+v.Name())
			}
		}
	}
}
