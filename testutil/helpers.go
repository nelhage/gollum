package testutil

import (
	"bytes"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/nelhage/gollum"
	"github.com/nelhage/gollum/parse"
)

// TestFile is a test case loaded from a testdata directory
type TestFile struct {
	Name string
	Path string
	Body []byte
}

// ListDir loads files from the named testdata/ subdirectory
func ListDir(t testing.TB, dir string) []TestFile {
	t.Helper()
	dir = path.Join("testdata", dir)
	f, e := os.Open(dir)
	if e != nil {
		t.Fatalf("open(%q): %v", dir, e)
	}
	defer f.Close()
	ents, e := f.Readdir(0)
	if e != nil {
		t.Fatalf("readdir(%q): %v", dir, e)
	}

	var out []TestFile
	for _, fi := range ents {
		if strings.HasPrefix(fi.Name(), ".") {
			continue
		}
		if !strings.HasSuffix(fi.Name(), ".gol") {
			continue
		}
		p := path.Join(dir, fi.Name())
		b, e := ioutil.ReadFile(p)
		if e != nil {
			t.Fatalf("read(%s): %v", p, e)
		}
		out = append(out, TestFile{Name: fi.Name(), Path: p, Body: b})
	}
	return out
}

// MustParse parses a TestFile into an AST, or aborts via t.Fatal
func MustParse(t testing.TB, tc TestFile) gollum.AST {
	t.Helper()
	buf := bytes.NewBuffer(tc.Body)
	ast, err := parse.Program(buf, tc.Name)
	if err != nil {
		t.Fatalf("parse(%q): %v", tc.Name, err)
	}
	return ast
}
