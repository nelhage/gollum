package lambda_test

import (
	"io/ioutil"
	"testing"

	"nelhage.com/lambda"
	"nelhage.com/lambda/testutil"
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
