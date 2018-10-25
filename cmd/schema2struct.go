package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"json-to-struct/internal/ahgenerator"
)

func main() {
	var w io.WriteCloser

	flag.Parse()
	schemaPath := flag.Arg(0)
	target := flag.Arg(1)

	result, err := ahgenerator.ToStruct(schemaPath)

	if err != nil {
		fmt.Printf("generating structs failed: %v", err)
	}

	if target != "" {
		targetFile, err := os.Create(target)

		if err != nil {
			log.Panicf("writing results to '%v' failed: %v", target, err)
		}

		w = targetFile
	} else {
		w = os.Stdout
	}


	fmt.Fprint(w, result)
	w.Close()
}
