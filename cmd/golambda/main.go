package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/kr/pretty"

	"nelhage.com/lambda"
	"nelhage.com/lambda/asteval"
	"nelhage.com/lambda/parse"
)

func main() {
	var (
		printAst  = flag.Bool("ast", false, "pretty-print the AST")
		printType = flag.Bool("type", false, "pretty-print the type")
		untyped   = flag.Bool("untyped", false, "Don't typecheck")
	)
	flag.Parse()
	if *untyped && *printType {
		log.Fatal("can't print types if not typechecking")
	}

	var r io.Reader
	var path string
	if len(flag.Args()) > 0 {
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
	ast, err := parse.Parse(r, path)

	if err != nil {
		log.Fatalf("parse error: %v", err)
	}

	if *printAst {
		pretty.Println("ast: ", ast)
	}

	var ty lambda.Type
	if !*untyped {
		ty, err = lambda.TypeCheck(ast, lambda.GlobalEnv)
		if err != nil {
			log.Fatalf("typechecking: %v", err)
		}
		if *printType {
			fmt.Println("type: ", lambda.PrintType(ty))
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
			fmt.Println("value: <fun> : ", lambda.PrintType(ty))
		}
	} else {
		pretty.Println("value: ", v)
	}
}
