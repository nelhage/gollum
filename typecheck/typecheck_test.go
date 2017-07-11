package typecheck

import (
	"bytes"
	"io/ioutil"
	"log"
	"testing"

	"nelhage.com/lambda/lambdatest"
	"nelhage.com/lambda/parse"
)

func TestTypeCheck(t *testing.T) {
	good := lambdatest.ListDir(t, "good")
	for _, tc := range good {
		t.Run("good/"+tc.Name, func(t *testing.T) {
			buf := bytes.NewBuffer(tc.Body)
			ast, err := parse.Parse(buf, tc.Name)
			if err != nil {
				log.Fatalf("parse(%q): %v", tc.Path, err)
			}
			ty, err := TypeCheck(ast, GlobalEnv)

			if err != nil {
				log.Fatalf("typecheck(%q): %v", tc.Name, err)
			}
			ioutil.WriteFile(
				tc.Path+".out",
				[]byte(PrintType(ty)),
				0644)
		})
	}

	bad := lambdatest.ListDir(t, "bad")
	for _, tc := range bad {
		t.Run("bad/"+tc.Name, func(t *testing.T) {
			buf := bytes.NewBuffer(tc.Body)
			ast, err := parse.Parse(buf, tc.Name)
			if err != nil {
				log.Fatalf("parse(%q): %v", tc.Path, err)
			}
			ty, err := TypeCheck(ast, GlobalEnv)

			if err == nil {
				log.Fatalf("typecheck(%q): %s", tc.Name, PrintType(ty))
			}
		})
	}
}
