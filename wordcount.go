package main

import (
	"bufio"
	"bytes"
	"io"
	"io/fs"
	"os"
	"strings"
)

type CountResult struct {
	Lines, Words, Chars int
}

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

func CountAllFromFs(fileSys fs.FS, fn string) (CountResult, error) {
	f, err := openFile(fileSys, fn)
	if err != nil {
		return CountResult{}, err
	}
	return countAll(f), nil
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

func countLines(f io.Reader) int {
	var count int
	scn := bufio.NewScanner(f)
	for scn.Scan() {
		count++
	}
	return count
}

func countWords(f io.Reader) int {
	var count int
	scn := bufio.NewScanner(f)
	for scn.Scan() {
		count += len(strings.Fields(scn.Text()))
	}
	return count
}

func countChars(f io.Reader) int {
	var count int
	scn := bufio.NewScanner(f)
	scn.Split(bufio.ScanBytes)
	for scn.Scan() {
		count++
	}
	return count
}

func countAll(r io.Reader) CountResult {
	// use io.TeeReader to duplicate r
	var buf = &bytes.Buffer{}
	rd := io.TeeReader(r, buf)
	lines, words := countLinesAndWords(rd)
	return CountResult{
		Lines: lines,
		Words: words,
		Chars: countChars(buf),
	}
}

func countLinesAndWords(f io.Reader) (lines int, words int) {
	scn := bufio.NewScanner(f)
	for scn.Scan() {
		lines += 1
		words += len(strings.Fields(scn.Text()))
	}
	return
}