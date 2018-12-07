package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/brahama/cf-doc/doc"
	"github.com/brahama/cf-doc/print"
)

var version = "v0.0.1"

const usage = `
  Usage:
    cf-doc [json | md | markdown] <file>...
    cf-doc -h | --help

  Examples:

    # Generate a JSON of inputs and outputs
    $ cf-doc json ./my-template.yaml

    # Generate markdown tables of inputs and outputs
    $ cf-doc md ./my-template.yaml

  Options:
    -h, --help     show help information

`

func main() {

	// Lets change how args are parsed to use flags.
	args := os.Args
	if len(args) <= 2 {
		log.Fatal(usage)
	}

	argOut := args[1]
	file := args[2]

	_, err := os.Stat(file)
	if err != nil {
		log.Fatal(err)
	}

	content, err := ioutil.ReadFile(file)
	doc := doc.Create(content)

	var out string

	switch {
	case argOut == "markdown":
		fallthrough
	case argOut == "md":
		out, err = print.Markdown(doc)
	case argOut == "json":
		out, err = print.Pretty(doc)
	default:
		out, err = print.Pretty(doc)
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out)
}
