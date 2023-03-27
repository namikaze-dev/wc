package main

import (
	"flag"
	"fmt"
	"os"
)

type Options struct {
	countlines bool
	countwords bool
	countchars bool
}

func main() {
	var options Options

	flag.BoolVar(&options.countlines, "l", false, "print count of lines")
	flag.BoolVar(&options.countwords, "w", false, "print count of words")
	flag.BoolVar(&options.countchars, "c", false, "print count of chars")
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		return
	}

	if noFlagPassed(options) {
		printCountWithOptions(Options{true, true, true}, args[0])
		return
	}

	printCountWithOptions(options, args[0])
}

func printCountWithOptions(options Options, fn string) {
	res, err := CountAllFromFs(os.DirFS("."), fn)
	if err != nil {
		fmt.Printf("./wc: %v", err)
		return
	}

	var output string
	if options.countlines {
		output = fmt.Sprintf("\t%v", res.Lines)
	}
	if options.countwords {
		output = fmt.Sprintf("%v\t%v", output, res.Words)
	}
	if options.countchars {
		output = fmt.Sprintf("%v\t%v", output, res.Chars)
	}
	output += fmt.Sprintf(" %v", fn)

	fmt.Println(output)
}

func noFlagPassed(options Options) bool {
	return !options.countlines && !options.countwords && !options.countchars
}
