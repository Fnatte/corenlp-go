package main

import (
	"fmt"
	"flag"
	"os"
	"os/user"
	"path"
)

var corenlpPath string

func main() {

	ParseInput()
	PostprocessInput()
	PrintHeader()

	if !CheckPath() {
		fmt.Printf("\"%s\" is not a vaild corenpl directory.\n", corenlpPath)
		os.Exit(1)
	}
	
	// Start corenlp
	if err := StartCorenlp(); err != nil {
		fmt.Println("Failed to start corenlp.")
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func ParseInput() {
	flag.StringVar(&corenlpPath, "path", "", "Path to stanford corenlp folder")
	flag.Parse()

	if len(os.Args) > 1 && os.Args[1] == "help" {
		flag.Usage()
		os.Exit(0)
	}
}

func PostprocessInput() {
	// Resolve ~ to home directory
	if len(corenlpPath) > 1 && corenlpPath[0] == '~' {
		if usr, err := user.Current(); err == nil {
			corenlpPath = path.Join(usr.HomeDir, corenlpPath[1:])
		} else {
			fmt.Println("Failed to resolve home directory '~'. Error was:")
			fmt.Println(err.Error())
		}
	}
}

func PrintHeader() {
	fmt.Println("corelnp-go 0.1")
}
