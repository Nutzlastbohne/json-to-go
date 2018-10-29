package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/nutzlastbohne/json-to-go/internal/ahgenerator"
)

const usageMessage = `
json-to-go - generates a go-struct from a json schema. The result is written to stdout by default. Specify [output-file] to write into a file.

usage: json-to-go <json-schema> [output-file]
			
Options
	[output-file]
		Writes the generated output into the given file. If it already exists, it will be overwritten!
`

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), usageMessage)
	}

	flag.Parse()

	argCount := len(flag.Args())
	if argCount < 1 || argCount > 2 {
		log.Println("Error - invalid argument count:", argCount, flag.Args())
		flag.Usage()
		return
	}

	schemaPath := flag.Arg(0)
	outputTarget := flag.Arg(1)

	// read schema
	result, err := ahgenerator.ToStruct(schemaPath)

	if err != nil {
		fmt.Printf("generating structs failed: %v", err)
		return
	}

	// write result
	var w io.WriteCloser
	if outputTarget != "" {
		targetFile, err := os.Create(outputTarget)

		if err != nil {
			log.Panicf("writing results to '%v' failed: %v", outputTarget, err)
		}

		w = targetFile
	} else {
		w = os.Stdout
	}
	defer w.Close()

	fmt.Fprint(w, result)
}
