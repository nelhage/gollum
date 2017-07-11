package main

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/kr/pretty"

	"nelhage.com/lambda/asteval"
	"nelhage.com/lambda/parse"
	"nelhage.com/lambda/typecheck"
)

func main() {
	var (
		printAst  = flag.Bool("ast", false, "pretty-print the AST")
		printType = flag.Bool("type", false, "pretty-print the type")
	)
	flag.Parse()

	var r io.Reader
	var path string
	if len(os.Args) > 1 {
		path = os.Args[1]
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
	ast, err := parse.Parse(r, path)

	if err != nil {
		log.Fatalf("parse error: %v", err)
	}

	if *printAst {
		pretty.Println("ast: ", ast)
	}

	ty, tyErr := typecheck.TypeCheck(ast, typecheck.GlobalEnv)
	if tyErr != nil {
		log.Fatalf("typechecking: %v", tyErr)
	}
	if *printType {
		pretty.Println("type: ", ty)
	}

	v, e := asteval.Eval(ast, asteval.GlobalEnv)
	if e != nil {
		log.Fatalf("eval error: %v", e)
	}

	pretty.Println("value: ", v)
}
