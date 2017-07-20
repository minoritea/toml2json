package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	// "github.com/naoina/toml"
	"github.com/pelletier/go-toml"
	"io/ioutil"
	"os"
)

const help = `toml2json converts a TOML formated text to a JSON text.
usage: toml2json [flags] [file]
       If no files are specified, toml2json reads the input text from the standard input.
flags:
       -help
       -h    show the help message
`

func printHelpAndExit(exitStatus int) {
	fmt.Println(help)
	os.Exit(exitStatus)
}

func printErrAndHelpThenExit(err error) {
	fmt.Println(err)
	printHelpAndExit(1)
}

func tomlToJson(data []byte) ([]byte, error) {
	tree, err := toml.LoadBytes(data)
	if err != nil {
		return []byte{}, err
	}
	json, err := json.Marshal(tree.ToMap())
	if err != nil {
		return []byte{}, err
	}
	return json, nil
}

func main() {
	h := flag.Bool("h", false, "show the help message")
	help := flag.Bool("help", false, "show the help message")
	flag.Parse()
	if *h || *help {
		printHelpAndExit(0)
	}

	args := flag.Args()
	var input []byte
	var err error
	if len(args) < 1 {
		input, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			printErrAndHelpThenExit(err)
		}
	} else {
		f, err := os.Open(args[0])
		if err != nil {
			printErrAndHelpThenExit(err)
		}
		defer f.Close()
		input, err = ioutil.ReadAll(f)
		if err != nil {
			printErrAndHelpThenExit(err)
		}
	}
	if len(input) == 0 {
		printErrAndHelpThenExit(errors.New("The input text must be non-empty."))
	}

	json, err := tomlToJson(input)
	if err != nil {
		printErrAndHelpThenExit(err)
	}

	fmt.Println(string(json))
	os.Exit(0)
}
