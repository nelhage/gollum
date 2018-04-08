package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/kr/pretty"
	"github.com/nelhage/gollum"
	"github.com/nelhage/gollum/asteval"
	"github.com/nelhage/gollum/parse"
)

func main() {
	var (
		printAst  = flag.Bool("ast", false, "pretty-print the AST")
		printType = flag.Bool("type", false, "pretty-print the type")
		untyped   = flag.Bool("untyped", false, "Don't typecheck")
		eval      = flag.String("e", "", "Evaluate code on the command line")
	)
	flag.Parse()
	if *untyped && *printType {
		log.Fatal("can't print types if not typechecking")
	}

	var r io.Reader
	var path string
	if *eval != "" {
		r = strings.NewReader(*eval)
		path = "-e"
	} else if len(flag.Args()) > 0 {
		path = flag.Arg(0)
		f, e := os.Open(path)
		if e != nil {
			log.Fatalf("open(%q): %v", path, e)
		}
		defer f.Close()
		r = f
	} else {
		r = os.Stdin
		path = "<stdin>"
	}
	ast, err := parse.Program(r, path)

	if err != nil {
		log.Fatalf("parse error: %v", err)
	}

	if *printAst {
		pretty.Println("ast: ", ast)
	}

	var ty gollum.Type
	if !*untyped {
		ty, err = gollum.TypeCheck(ast, gollum.GlobalEnv)
		if err != nil {
			log.Fatalf("typechecking: %v", err)
		}
		if *printType {
			fmt.Println("type: ", gollum.PrintType(ty))
		}
	}

	v, e := asteval.Eval(ast, asteval.GlobalEnv)
	if e != nil {
		log.Fatalf("eval error: %v", e)
	}

	if cl, ok := v.(*asteval.Closure); ok {
		if *untyped {
			fmt.Printf("value: <fun>/%d\n", len(cl.Args))
		} else {
			fmt.Println("value: <fun> : ", gollum.PrintType(ty))
		}
	} else {
		pretty.Println("value: ", v)
	}
}
