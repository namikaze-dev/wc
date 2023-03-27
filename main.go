package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var options struct {
		countlines bool
		countwords bool
		countchars bool
	}

	flag.BoolVar(&options.countlines, "l", false, "print count of lines")
	flag.BoolVar(&options.countwords, "w", false, "print count of words")
	flag.BoolVar(&options.countchars, "c", false, "print count of chars")
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		return
	}

	if noFlagPassed(options.countlines, options.countwords, options.countchars) {
		countall(args[0])
	}

	if options.countlines {
		countlines(args[0])
	} else if options.countwords {
		countwords(args[0])
	} else if options.countchars {
		countchars(args[0])
	}
}

func countlines(fn string) {
	linesCount, err := CountLinesFromFs(os.DirFS("."), fn)
	if err != nil {
		fmt.Printf("./wc: %v", err)
	} else {
		fmt.Printf("\t\t%v %v", linesCount, fn)
	}
}

func countwords(fn string) {
	wordsCount, err := CountWordsFromFs(os.DirFS("."), fn)
	if err != nil {
		fmt.Printf("./wc: %v", err)
	} else {
		fmt.Printf("\t\t%v %v", wordsCount, fn)
	}
}

func countchars(fn string) {
	chars, err := CountCharsFromFs(os.DirFS("."), fn)
	if err != nil {
		fmt.Printf("./wc: %v", err)
	} else {
		fmt.Printf("\t\t%v %v", chars, fn)
	}
}

func countall(fn string) {
	res, err := CountAllFromFs(os.DirFS("."), fn)
	if err != nil {
		fmt.Printf("./wc: %v", err)
		return
	} 

	fmt.Printf("\t%v\t%v\t%v %v", res.Lines, res.Words, res.Chars, fn)
}

func noFlagPassed(clines, cwords, cchars bool) bool {
	return !clines && !cwords && !cchars
}