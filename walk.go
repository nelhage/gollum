package gollum

func mapVars(fn func(*TypeVariable) Type, ty Type) Type {
	switch n := ty.(type) {
	case *TypeVariable:
		return fn(n)
	case *AtomicType:
		return n
	case *FunctionType:
		return &FunctionType{
			Dom:   mapVars(fn, n.Dom),
			Range: mapVars(fn, n.Range),
		}
	case *TupleType:
		elts := make([]Type, len(n.Elts))
		for i, e := range n.Elts {
			elts[i] = mapVars(fn, e)
		}
		return &TupleType{elts}
	default:
		panic(bad("mapVars", ty))
	}
}

func eachVar(fn func(*TypeVariable), ty Type) {
	switch n := ty.(type) {
	case *TypeVariable:
		fn(n)
	case *AtomicType:
	case *FunctionType:
		eachVar(fn, n.Dom)
		eachVar(fn, n.Range)
	case *TupleType:
		for _, e := range n.Elts {
			eachVar(fn, e)
		}
	case *Forall:
		for _, e := range n.Vars {
			eachVar(fn, e)
		}
		eachVar(fn, n.Type)
	default:
		panic(bad("eachVar", ty))
	}
}
