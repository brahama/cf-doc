package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/brahma/cf-doc/doc"
	"github.com/brahma/cf-doc/inject"
	"github.com/brahma/cf-doc/print"
)

var version = "v0.0.1"

const usage = `
  Usage:
    cf-doc [json | md | markdown] <file>...
    cf-doc inject <file> --output-file <readme>
    cf-doc -h | --help

  Examples:

    # View inputs and outputs
    $ cf-doc ./my-template.yaml

    # Generate a JSON of inputs and outputs
    $ cf-doc json ./my-template.yaml

    # Generate markdown tables of inputs and outputs
    $ cf-doc md ./my-template.yaml

    # Inject markdown docs into an existing Readme between
    # "<!-- cf-doc:start -->" and "<!-- cf-doc:end -->" markers
    $ cf-doc inject ./my-template.yaml --output-file Readme.md

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

	if argOut == "inject" {
		injectFlags := flag.NewFlagSet("inject", flag.ExitOnError)
		outputFile := injectFlags.String("output-file", "", "file containing cf-doc:start/cf-doc:end markers to inject docs into")
		injectFlags.Parse(args[3:])

		if *outputFile == "" {
			log.Fatal("inject requires --output-file <readme>")
		}

		if err := runInject(file, *outputFile); err != nil {
			log.Fatal(err)
		}
		return
	}

	_, err := os.Stat(file)
	if err != nil {
		log.Fatal(err)
	}

	content, err := ioutil.ReadFile(file)
	doc := doc.Create(content)

	var out string

	switch {
	case argOut == "markdown":
		out, err = print.Markdown(doc)
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

func runInject(templateFile, outputFile string) error {
	if _, err := os.Stat(templateFile); err != nil {
		return err
	}

	templateContent, err := ioutil.ReadFile(templateFile)
	if err != nil {
		return err
	}

	markdown, err := print.Markdown(doc.Create(templateContent))
	if err != nil {
		return err
	}

	existing, err := ioutil.ReadFile(outputFile)
	if err != nil {
		return err
	}

	injected, err := inject.Content(string(existing), markdown)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(outputFile, []byte(injected), 0644)
}
