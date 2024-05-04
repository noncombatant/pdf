package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/noncombatant/pdf"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:\n\n")
	fmt.Fprintf(os.Stderr, "pdf2txt pdf-file [...]\n\n")
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

		for i := 1; i <= reader.NumPage(); i++ {
			y := 0.0
			for _, t := range reader.Page(i).Content().Text {
				if t.Y != y {
					y = t.Y
					os.Stdout.WriteString("\n")
				}
				os.Stdout.WriteString(t.S)
			}
		}
	}
}
