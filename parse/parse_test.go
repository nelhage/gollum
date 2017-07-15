package parse_test

import (
	"bytes"
	"io/ioutil"
	"log"
	"testing"

	"nelhage.com/lambda/parse"
	"nelhage.com/lambda/testutil"

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
			ast, err := parse.Parse(buf, tc.Name)
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
