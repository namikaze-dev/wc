package main_test

import (
	"errors"
	"io/fs"
	"os"
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/namikaze-dev/wc2"
)

func TestCountAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		fn := "file.txt"
		got, err := main.CountAllFromFs(testFS, fn)
		if err != nil {
			t.Fatalf("unexpected error from main.CountAllFromFs: %v", err)
		}

		want := main.CountResult{
			Lines: 4,
			Words: 9,
			Chars: 33,
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+v want %+v", got, want)
		}
	})
	
	t.Run("failure file doesn't exist", func(t *testing.T) {
		fn := "non-existing-file.txt"
		_, err := main.CountAllFromFs(testFS, fn)
		if !errors.Is(err, os.ErrNotExist) {
			t.Fatalf("want error %v got error %v", os.ErrNotExist, err)
		}
	})

	
	t.Run("failure is a directory", func(t *testing.T) {
		fn := "dir"
		_, err := main.CountAllFromFs(testFS, fn)
		if !errors.Is(err, os.ErrInvalid) {
			t.Fatalf("want error %v got error %v", os.ErrInvalid, err)
		}
	})
	
	t.Run("failure permission denied", func(t *testing.T) {
		_, err := main.CountAllFromFs(FailingPermFs{}, "")
		if !errors.Is(err, os.ErrPermission) {
			t.Fatalf("want error %v got error %v", os.ErrPermission, err)
		}
	})
}

var testFS = fstest.MapFS{
	"file.txt": {Data: []byte("line 1\n\nline 2\nline 3\tlast line 4")},
	"dir": {Mode: fs.ModeDir},
}

type FailingPermFs struct {}

func (fs FailingPermFs) Open(fn string) (fs.File, error) {
	return nil, os.ErrPermission
}