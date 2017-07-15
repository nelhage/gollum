package lambda

import (
	"fmt"
)

// TypeCheck typechecks an AST and returns the type of the AST
// structure
func TypeCheck(ast AST, env *TypeEnv) (Type, error) {
	switch n := ast.(type) {
	case *Boolean:
		return boolType, nil
	case *String:
		return strType, nil
	case *Integer:
		return intType, nil
	case *Variable:
		t := env.Lookup(n.Var)
		if t == nil {
			return nil, UnboundVariable{ast, n.Var}
		}
		return t, nil
	case *Abstraction:
		var names []string
		var types []Type
		for _, v := range n.Vars {
			tv := v.(*TypedName)
			if tv.Type == nil {
				return nil, UntypedName{tv, tv.Name}
			}
			argType, e := ParseType(tv.Type)
			if e != nil {
				return nil, e
			}
			names = append(names, tv.Name)
			types = append(types, argType)
		}
		env := env.Extend(names, types)

		rtype, e := TypeCheck(n.Body, env)

		if e != nil {
			return nil, e
		}
		return &FunctionType{
			Dom: &TupleType{
				Elts: types,
			},
			Range: rtype,
		}, nil
	case *Application:
		ftype, err := TypeCheck(n.Func, env)
		if err != nil {
			return nil, err
		}

		fnt, ok := ftype.(*FunctionType)
		if !ok {
			return nil, TypeError{
				Node:     n.Func,
				Expected: "function",
				Got:      ftype,
			}
		}

		var args []Type
		for _, a := range n.Args {
			argType, err := TypeCheck(a, env)
			if err != nil {
				return nil, err
			}
			args = append(args, argType)
		}
		argType := &TupleType{Elts: args}
		if !Equal(fnt.Dom, argType) {
			return nil, TypeError{
				Node:       n,
				ExpectedTy: fnt.Dom,
				Got:        argType,
			}
		}
		return fnt.Range, nil

	case *If:
		cdType, err := TypeCheck(n.Condition, env)
		if err != nil {
			return nil, err
		}
		conType, err := TypeCheck(n.Consequent, env)
		if err != nil {
			return nil, err
		}
		altType, err := TypeCheck(n.Alternate, env)
		if err != nil {
			return nil, err
		}

		if !Equal(boolType, cdType) {
			return nil, TypeError{
				Node:       n.Condition,
				Got:        cdType,
				ExpectedTy: boolType,
			}
		}
		if !Equal(conType, altType) {
			return nil, TypeError{
				Node:       n.Alternate,
				Got:        altType,
				ExpectedTy: conType,
			}
		}
		return conType, nil

	case *TypedName, *TyName, *TyArrow:
		panic(fmt.Sprintf("bad toplevel ast: %#v", ast))
	default:
		panic(fmt.Sprintf("unhandled ast: %#v", ast))
	}
}

// ParseType parses an AST that refers to a type into a type object
func ParseType(ast AST) (Type, error) {
	switch n := ast.(type) {
	case *TyName:
		ty := GlobalTypes[n.Type]
		if ty == nil {
			return nil, UnboundType{ast, n.Type}
		}
		return ty, nil
	case *TyArrow:
		dom, err := ParseType(n.Dom)
		if err != nil {
			return nil, err
		}
		if _, ok := dom.(*TupleType); !ok {
			dom = &TupleType{
				Elts: []Type{dom},
			}
		}
		range_, err := ParseType(n.Range)
		if err != nil {
			return nil, err
		}
		return &FunctionType{
			Dom:   dom,
			Range: range_,
		}, nil
	case *TyTuple:
		var tys []Type
		for _, elt := range n.Elts {
			ty, e := ParseType(elt)
			if e != nil {
				return nil, e
			}
			tys = append(tys, ty)
		}
		return &TupleType{Elts: tys}, nil
	default:
		panic(fmt.Sprintf("bad ast node to ParseType: %#v", ast))
	}
}

// Equal checks two types for equality
func Equal(l, r Type) bool {
	switch t := l.(type) {
	case *AtomicType:
		ra, ok := r.(*AtomicType)
		if !ok {
			return false
		}
		return t.Name == ra.Name
	case *FunctionType:
		rf, ok := r.(*FunctionType)
		if !ok {
			return false
		}
		return Equal(t.Dom, rf.Dom) && Equal(t.Range, rf.Range)
	case *TupleType:
		rt, ok := r.(*TupleType)
		if !ok {
			return false
		}
		if len(t.Elts) != len(rt.Elts) {
			return false
		}
		for i, e := range t.Elts {
			if !Equal(e, rt.Elts[i]) {
				return false
			}
		}
		return true
	default:
		panic(fmt.Sprintf("unhandled type: %#v", l))
	}
}
