//
// tsort.go (go-coreutils) 0.1
// Copyright (C) 2014, The GO-Coreutils Developers.
//
// Written By: Akira Hayakawa
//
package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	help_text string = `
    Usage: tsort [OPTIONS] FILE
    
    Topological sort the strings in FILE. Strings are defined as any sequence of tokes separated by
    whitespace (tab, space, or newline). If FILE it not passed, stdin is used instead.

        --help        display this help and exit
        --version     output version information and exit
    `
	version_text = `
    tsort (go-coreutils) 0.1

    Copyright (C) 2014, The GO-Coreutils Developers.
    This program comes with ABSOLUTELY NO WARRANTY; for details see
    LICENSE. This is free software, and you are welcome to redistribute 
    it under certain conditions in LICENSE.
`
)

var (
	help    = flag.Bool("help", false, help_text)
	version = flag.Bool("version", false, version_text)
)

func processFlags() {
	if *help {
		fmt.Println(help_text)
		os.Exit(0)
	}

	if *version {
		fmt.Println(version_text)
		os.Exit(0)
	}
}

func main() {
	flag.Parse()
	processFlags()
}
