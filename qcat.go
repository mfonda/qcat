package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var leftDelim = flag.String("l", `"`, "left delimiter")
var rightDelim = flag.String("r", `"`, "right delimiter")
var escape = flag.Bool("escape", true, "escape all quotes within each line")
var trim = flag.Bool("trim", true, "trim all whitespace from each line before quoting")

func main() {
	flag.Parse()
	if flag.NArg() == 0 || flag.Arg(1) == "-" {
		quotecat(os.Stdin)
	} else {
		for _, fn := range flag.Args() {
			file, err := os.Open(fn)
			if err != nil {
				panic(err)
			}
			quotecat(file)
		}
	}
}

func quotecat(file *os.File) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if *escape {
			line = strings.Replace(line, `"`, `\"`, -1)
		}
		if *trim {
			line = strings.Trim(line, " \r\n\t")
		}
		fmt.Printf("%s%s%s\n", *leftDelim, line, *rightDelim)
	}
}
