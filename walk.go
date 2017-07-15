package lambda

func foldType(fn func(Type) Type, ty Type) Type {
	switch n := ty.(type) {
	case *AtomicType, *TypeVariable:
		return fn(ty)
	case *FunctionType:
		return &FunctionType{
			Dom:   foldType(fn, n.Dom),
			Range: foldType(fn, n.Range),
		}
	case *TupleType:
		elts := make([]Type, len(n.Elts))
		for i, e := range n.Elts {
			elts[i] = foldType(fn, e)
		}
		return &TupleType{elts}
	default:
		panic(bad("foldType", ty))
	}
}
