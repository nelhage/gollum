package parse

import (
	"bytes"
	"io/ioutil"
	"log"
	"testing"

	"nelhage.com/lambda/lambdatest"

	"github.com/kr/pretty"
)

func TestParser(t *testing.T) {
	good := lambdatest.ListDir(t, "good")
	for _, tc := range good {
		t.Run("good/"+tc.Name, func(t *testing.T) {
			buf := bytes.NewBuffer(tc.Body)
			ast, err := Parse(buf, tc.Name)
			if err != nil {
				log.Fatalf("parse(%q): %v", tc.Name, err)
			}
			ioutil.WriteFile(
				tc.Path+".ast",
				[]byte(pretty.Sprint(ast)),
				0644)
		})
	}

	bad := lambdatest.ListDir(t, "bad")
	for _, tc := range bad {
		t.Run("bad/"+tc.Name, func(t *testing.T) {
			buf := bytes.NewBuffer(tc.Body)
			ast, err := Parse(buf, tc.Name)
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
