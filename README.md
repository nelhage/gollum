# gollum -- a toy typed lambda calculus

Gollum implements a simply-typed lambda calculus, along with
Hindley-Milner-style unification, in Go.

For source examples see any of the `testdata/good` directories.

The typechecker (and inference/unification engine) is probbaly the
most interesting code, and lives in [typecheck.go](typecheck.go).
