package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/nutzlastbohne/json-to-go/internal/ahgenerator"
)

func main() {
	var w io.WriteCloser

	flag.Parse()
	schemaPath := flag.Arg(0)
	target := flag.Arg(1)

	result, err := ahgenerator.ToStruct(schemaPath)

	if err != nil {
		fmt.Printf("generating structs failed: %v", err)
		return
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
	defer w.Close()


	fmt.Fprint(w, result)
}
