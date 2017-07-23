package parse_test

import (
	"bytes"
	"io/ioutil"
	"log"
	"testing"

	"github.com/nelhage/gollum/parse"
	"github.com/nelhage/gollum/testutil"

	"github.com/kr/pretty"
)

func TestParser(t *testing.T) {
	good := testutil.ListDir(t, "good")
	for _, tc := range good {
		t.Run("good/"+tc.Name, func(t *testing.T) {
			ast := testutil.MustParse(t, tc)
			ioutil.WriteFile(
				tc.Path+".ast",
				[]byte(pretty.Sprint(ast)),
				0644)
		})
	}

	bad := testutil.ListDir(t, "bad")
	for _, tc := range bad {
		t.Run("bad/"+tc.Name, func(t *testing.T) {
			buf := bytes.NewBuffer(tc.Body)
			ast, err := parse.Program(buf, tc.Name)
			if err == nil {
				ioutil.WriteFile(
					tc.Path+".ast",
					[]byte(pretty.Sprint(ast)),
					0644)
				log.Fatalf("parse(%q): ok", tc.Name)
			}
		})
	}
}

func TestParseType(t *testing.T) {
	types := testutil.ListDir(t, "type")
	for _, tc := range types {
		t.Run("types/"+tc.Name, func(t *testing.T) {
			_, err := parse.Type(bytes.NewReader(tc.Body), tc.Path)
			if err != nil {
				t.Fatalf("parse(%q): %v", tc.Path, err)
			}
		})
	}
}
