package main

import (
	"bufio"
	"io/fs"
	"os"
	"strings"
)

func CountLinesFromFs(fileSys fs.FS, fn string) (int, error) {
	f, err := openFile(fileSys, fn)
	if err != nil {
		return 0, err
	}
	return countLines(f), nil
}

func CountWordsFromFs(fileSys fs.FS, fn string) (int, error) {
	f, err := openFile(fileSys, fn)
	if err != nil {
		return 0, err
	}
	return countWords(f), nil
}

func CountCharsFromFs(fileSys fs.FS, fn string) (int, error) {
	f, err := openFile(fileSys, fn)
	if err != nil {
		return 0, err
	}
	return countChars(f), nil
}

func openFile(fileSys fs.FS, fn string) (fs.File, error) {
	f, err := fileSys.Open(fn)
	if err != nil {
		return nil, err
	}

	fstat, err := f.Stat()
	if err != nil {
		return nil, err
	}

	if fstat.IsDir() {
		return nil, os.ErrInvalid
	}

	return f, nil
}

func countLines(f fs.File) int {
	var count int
	scn := bufio.NewScanner(f)
	for scn.Scan() {
		count++
	}
	return count
}

func countWords(f fs.File) int {
	var count int
	scn := bufio.NewScanner(f)
	for scn.Scan() {
		count += len(strings.Fields(scn.Text()))
	}
	return count
}

func countChars(f fs.File) int {
	var count int
	scn := bufio.NewScanner(f)
	scn.Split(bufio.ScanBytes)
	for scn.Scan() {
		count++
	}
	return count
}
