package main_test

import (
	"bytes"
	"strings"
	"testing"

	main "github.com/namikaze-dev/wc2"
)

func TestPrintCountWithOptionsSuccesses(t *testing.T) {
	cases := []struct {
		name    string
		fn      string
		options main.Options
		want    string
	}{
		{
			name:    "success no options",
			fn:      "file.txt",
			options: main.Options{},
			want:    "\t4\t9\t33 file.txt\n",
		},
		{
			name:    "success all options",
			fn:      "file.txt",
			options: main.Options{},
			want:    "\t4\t9\t33 file.txt\n",
		},
		{
			name:    "success lines and chars",
			fn:      "file.txt",
			options: main.Options{CountChars: true, CountLines: true},
			want:    "\t4\t33 file.txt\n",
		},
		{
			name:    "success lines lonly",
			fn:      "file.txt",
			options: main.Options{CountLines: true},
			want:    "\t4 file.txt\n",
		},
		{
			name:    "success words lonly",
			fn:      "file.txt",
			options: main.Options{CountWords: true},
			want:    "\t9 file.txt\n",
		},
		{
			name:    "success chars lonly",
			fn:      "file.txt",
			options: main.Options{CountChars: true},
			want:    "\t33 file.txt\n",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			stdout := &bytes.Buffer{}
			main.PrintCountWithOptions(stdout, testFS, c.options, c.fn)
			got := stdout.String()
			if got != c.want {
				t.Errorf("got %q want %q", got, c.want)
			}
		})
	}
}

func TestPrintCountWithOptionsFailures(t *testing.T) {
	t.Run("failure not exist", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		main.PrintCountWithOptions(stdout, testFS, main.Options{}, "not-existing-file.txt")
		got := stdout.String()
		if !strings.Contains(got, "not exist") {
			t.Errorf("want error containing 'not exist', got %v", got)
		}
	})

	t.Run("failure directory", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		main.PrintCountWithOptions(stdout, testFS, main.Options{}, "dir")
		got := stdout.String()
		if !strings.Contains(got, "invalid argument") {
			t.Errorf("want error containing 'invalid argument', got %v", got)
		}
	})

	t.Run("failure permission denied", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		main.PrintCountWithOptions(stdout, FailingPermFs{}, main.Options{}, "dir")
		got := stdout.String()
		if !strings.Contains(got, "permission denied") {
			t.Errorf("want error containing 'permission denied', got %v", got)
		}
	})
}

func TestPrintCountFromReader(t *testing.T) {
	stdout := &bytes.Buffer{}
	input := strings.NewReader("line 1\n\nline 2\nline 3\tlast line 4")
	main.PrintCountFromReader(stdout, input)

	got := stdout.String()
	want := "\t4\t9\t33\n"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestPrintCountFromArgs(t *testing.T) {
	stdout := &bytes.Buffer{}
	main.PrintCountFromArgs(stdout, testFS, main.Options{}, []string{"file.txt", "file2.txt"})

	got := stdout.String()
	want := "\t4\t9\t33 file.txt\n\t1\t2\t12 file2.txt\n\t5\t11\t45 total\n"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
