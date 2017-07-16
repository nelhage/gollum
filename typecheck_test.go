package lambda_test

import (
	"io/ioutil"
	"regexp"
	"testing"

	"nelhage.com/lambda"
	"nelhage.com/lambda/testutil"
)

var (
	typeComment = regexp.MustCompile(`^# type: ([^\n]*)\n`)
)

func TestTypeCheck(t *testing.T) {
	good := testutil.ListDir(t, "good")
	for _, tc := range good {
		t.Run("good/"+tc.Name, func(t *testing.T) {
			ast := testutil.MustParse(t, tc)
			ty, err := lambda.TypeCheck(ast, lambda.GlobalEnv)

			if err != nil {
				t.Fatalf("typecheck(%q): %v", tc.Name, err)
			}
			ioutil.WriteFile(
				tc.Path+".out",
				[]byte(lambda.PrintType(ty)),
				0644)

			groups := typeComment.FindSubmatch(tc.Body)
			if groups != nil {
				got := lambda.PrintType(ty)
				want := groups[1]
				if string(want) != got {
					t.Errorf("want type=%q got=%q",
						want, got,
					)
				}
			}
		})
	}

	bad := testutil.ListDir(t, "bad")
	for _, tc := range bad {
		t.Run("bad/"+tc.Name, func(t *testing.T) {
			ast := testutil.MustParse(t, tc)
			ty, err := lambda.TypeCheck(ast, lambda.GlobalEnv)

			if err == nil {
				t.Fatalf("typecheck(%q): %s", tc.Name, lambda.PrintType(ty))
			}
		})
	}
}
