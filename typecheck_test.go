package gollum_test

import (
	"bytes"
	"io/ioutil"
	"regexp"
	"testing"

	"github.com/nelhage/gollum"
	"github.com/nelhage/gollum/parse"
	"github.com/nelhage/gollum/testutil"
)

var (
	typeComment = regexp.MustCompile(`^# type: ([^\n]*)\n`)
)

func TestTypeCheck(t *testing.T) {
	good := testutil.ListDir(t, "good")
	for _, tc := range good {
		t.Run("good/"+tc.Name, func(t *testing.T) {
			ast := testutil.MustParse(t, tc)
			ty, err := gollum.TypeCheck(ast, gollum.GlobalEnv)

			if err != nil {
				t.Fatalf("typecheck(%q): %v", tc.Name, err)
			}
			ioutil.WriteFile(
				tc.Path+".out",
				[]byte(gollum.PrintType(ty)),
				0644)

			groups := typeComment.FindSubmatch(tc.Body)
			if groups != nil {
				t.Logf("%q: want: %q", tc.Name, groups[1])
				ast, err := parse.Type(bytes.NewReader(groups[1]), "<#type>")
				var want gollum.Type
				if err == nil {
					want, err = gollum.ParseType(ast, gollum.GlobalEnv)
				}
				if err == nil {
					_, err = gollum.Unify(want, ty)
				}
				if err != nil {
					t.Fatalf(err.Error())
				}
			}
		})
	}

	bad := testutil.ListDir(t, "bad")
	for _, tc := range bad {
		t.Run("bad/"+tc.Name, func(t *testing.T) {
			ast := testutil.MustParse(t, tc)
			ty, err := gollum.TypeCheck(ast, gollum.GlobalEnv)

			if err == nil {
				t.Fatalf("typecheck(%q): %s", tc.Name, gollum.PrintType(ty))
			}
		})
	}
}

func BenchmarkTypeCheck(b *testing.B) {
	good := testutil.ListDir(b, "good")
	for _, tc := range good {
		b.Run(tc.Name, func(b *testing.B) {
			ast := testutil.MustParse(b, tc)
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_, err := gollum.TypeCheck(ast, gollum.GlobalEnv)

				if err != nil {
					b.Fatalf("typecheck(%q): %v", tc.Name, err)
				}
			}
		})
	}
}
