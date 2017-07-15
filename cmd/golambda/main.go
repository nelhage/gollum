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
	)
	flag.Parse()

	var r io.Reader
	var path string
	if len(flag.Args()) > 1 {
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

	ty, tyErr := lambda.TypeCheck(ast, lambda.GlobalEnv)
	if tyErr != nil {
		log.Fatalf("typechecking: %v", tyErr)
	}
	if *printType {
		fmt.Println("type: ", lambda.PrintType(ty))
	}

	v, e := asteval.Eval(ast, asteval.GlobalEnv)
	if e != nil {
		log.Fatalf("eval error: %v", e)
	}

	if _, ok := v.(*asteval.Closure); ok {
		fmt.Println("value: <fun> : ", lambda.PrintType(ty))
	} else {
		pretty.Println("value: ", v)
	}
}
