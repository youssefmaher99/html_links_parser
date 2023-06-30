package main

import (
	"fmt"
	"log"
	"os"

	"github.com/youssefmaher99/link_parser/parser"
)

func main() {
	fmt.Println()
	files := []string{"ex1.html", "ex2.html", "ex3.html", "ex4.html"}
	for _, file := range files {
		fmt.Printf("------------- %s -------------\n", file)
		fopen, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		links, err := parser.Parse(fopen)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", links)
	}
}
