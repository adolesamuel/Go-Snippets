package main

import (
	"log"
	"os"
)

func makeDirectory() {
	if er := os.Mkdir("a", os.ModePerm); er != nil {
		log.Fatal(er)
	}
}
