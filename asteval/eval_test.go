package asteval

import (
	"path"
	"testing"

	"nelhage.com/lambda/testutil"
)

func TestEval(t *testing.T) {
	good := testutil.ListDir(t, "good")
	for _, tc := range good {
		t.Run(path.Join("good", tc.Name), func(t *testing.T) {
			ast := testutil.MustParse(t, tc)
			_, err := Eval(ast, GlobalEnv)

			if err != nil {
				t.Fatalf("eval(%q): %v", tc.Name, err)
			}
		})
	}

	bad := testutil.ListDir(t, "bad")
	for _, tc := range bad {
		t.Run(path.Join("bad", tc.Name), func(t *testing.T) {
			ast := testutil.MustParse(t, tc)
			v, err := Eval(ast, GlobalEnv)

			if err == nil {
				t.Fatalf("typecheck(%q): %v", tc.Name, v)
			}
		})
	}
}
