package main

import (
	"log"
	"os"

	"github.com/youssefmaher99/link_parser/parser"
)

func main() {
	files := []string{"ex1.html", "ex2.html", "ex3.html", "ex4.html"}
	for _, file := range files {
		fopen, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		parser.Parse(fopen)
	}
}
