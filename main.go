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

	PrintCountFromArgs(os.Stdout, os.DirFS("."), options, args)
}

func PrintCountFromArgs(w io.Writer, fsys fs.FS, options Options, files []string) {
	var total CountResult
	for _, fn := range files {
		res := PrintCountWithOptions(w, fsys, options, fn)
		total.Lines += res.Lines
		total.Words += res.Words
		total.Chars += res.Chars
	}
	if len(files) > 1 {
		fmt.Fprintf(w, "\t%v\t%v\t%v total\n", total.Lines, total.Words, total.Chars)
	}
}

func PrintCountWithOptions(w io.Writer, fsys fs.FS, options Options, fn string) CountResult {
	if noFlagPassed(options) {
		options = Options{true, true, true}
	}

	res, err := CountAllFromFs(fsys, fn)
	if err != nil {
		fmt.Fprintf(w, "wc: %v: %v\n", fn, err)
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
	return res
}

func PrintCountFromReader(w io.Writer, r io.Reader) {
	res := CountAll(r)
	fmt.Fprintf(w, "\t%v\t%v\t%v\n", res.Lines, res.Words, res.Chars)
}

func noFlagPassed(options Options) bool {
	return !options.CountLines && !options.CountWords && !options.CountChars
}
