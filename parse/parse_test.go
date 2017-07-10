package parse

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/kr/pretty"
)

type testFile struct {
	name string
	body []byte
}

func listDir(t *testing.T, dir string) []testFile {
	f, e := os.Open(dir)
	if e != nil {
		t.Fatalf("open(%q): %v", dir, e)
	}
	defer f.Close()
	ents, e := f.Readdir(0)
	if e != nil {
		t.Fatalf("readdir(%q): %v", dir, e)
	}

	var out []testFile
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
		out = append(out, testFile{fi.Name(), b})
	}
	return out
}

func TestParser(t *testing.T) {
	good := listDir(t, "zoo/good")
	for _, tc := range good {
		t.Run(tc.name, func(t *testing.T) {
			buf := bytes.NewBuffer(tc.body)
			ast, err := Parse(buf, tc.name)
			if err != nil {
				log.Fatalf("parse(%q): %v", tc.name, err)
			}
			ioutil.WriteFile(
				path.Join("zoo/good", tc.name+".ast"),
				[]byte(pretty.Sprint(ast)),
				0644)
		})
	}
}
