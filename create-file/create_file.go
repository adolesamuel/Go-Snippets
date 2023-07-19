package main

import (
	"log"
	"os"
)

func createFile() {
	// creation of an empty file
	// using the Create() method
	myFile, es := os.Create("Heyy.txt")
	if es != nil {
		log.Fatal(es)
	}
	log.Println(myFile)
	myFile.Close()
}
