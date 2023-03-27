package main

import (
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
)

type Options struct {
	CountLines bool
	CountWords bool
	CountChars bool
}

func main() {
	var options Options

	flag.BoolVar(&options.CountLines, "l", false, "print count of lines")
	flag.BoolVar(&options.CountWords, "w", false, "print count of words")
	flag.BoolVar(&options.CountChars, "c", false, "print count of chars")
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		PrintCountFromReader(os.Stdout, os.Stdin)
		return
	}

	PrintCountWithOptions(os.Stdout, os.DirFS("."), options, args[0])
}

func PrintCountWithOptions(w io.Writer, fsys fs.FS, options Options, fn string) {
	if noFlagPassed(options) {
		options = Options{true, true, true}
	}

	res, err := CountAllFromFs(fsys, fn)
	if err != nil {
		fmt.Fprintf(w, "wc: %v: %v\n", fn, err)
		return
	}

	var output string
	if options.CountLines {
		output = fmt.Sprintf("\t%v", res.Lines)
	}
	if options.CountWords {
		output = fmt.Sprintf("%v\t%v", output, res.Words)
	}
	if options.CountChars {
		output = fmt.Sprintf("%v\t%v", output, res.Chars)
	}
	output += fmt.Sprintf(" %v", fn)

	fmt.Fprintln(w, output)
}

func PrintCountFromReader(w io.Writer, r io.Reader) {
	res := CountAll(r)
	fmt.Fprintf(w, "\t%v\t%v\t%v\n", res.Lines, res.Words, res.Chars)
}

func noFlagPassed(options Options) bool {
	return !options.CountLines && !options.CountWords && !options.CountChars
}
