package main

import (
	"flag"
	"os"
	"bufio"
	"strings"
	"log"
)

var file string
var prefix string

func main() {

	flag.StringVar(&file, "file", "application.properties", "Valid path to file")
	flag.StringVar(&prefix, "prefix", "#", "String used to comment")
	flag.Parse()

	txt, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer txt.Close()

	nf, err := os.Create(file + ".new")
	if err != nil {
		panic(err)
	}
	defer nf.Close()

	w := bufio.NewWriter(nf)
	defer w.Flush()

	scanner := bufio.NewScanner(txt)
	for scanner.Scan() {
		text := scanner.Text()
		if !strings.HasPrefix(text, prefix) {
			_, err := w.WriteString(text+"\n")
			if err != nil {
				panic(err)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
