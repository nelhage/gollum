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
		panic(bad("foldType", ty))
	}
}
