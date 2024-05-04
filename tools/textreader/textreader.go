// Copyright 2024 Chris Palmer, https://noncombatant.org/

package main

// This program is the same as — should produce the bytewise-identical result as
// — cmd/pdf2txt. It uses [TextReader] instead of immediately writing to stdout.

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/noncombatant/pdf"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:\n\n")
	fmt.Fprintf(os.Stderr, "textreader pdf-file [...]\n\n")
	os.Exit(1)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() == 0 {
		usage()
	}

	for _, pathname := range flag.Args() {
		reader, e := pdf.Open(pathname)
		if e != nil {
			fmt.Fprintln(os.Stderr, e)
			continue
		}

		textReader := pdf.NewTextReader(reader)
		io.Copy(os.Stdout, textReader)
	}
}
