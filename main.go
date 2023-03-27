package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var options struct {
		countlines bool
	}

	flag.BoolVar(&options.countlines, "l", true, "print count of lines")
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		return
	}

	linesCount, err := CountLinesFromFs(os.DirFS("."), args[0])
	if err != nil {
		fmt.Printf("./wc: %v", err)
	} else {
		fmt.Printf("\t\t%v %v", linesCount, args[0])
	}
}