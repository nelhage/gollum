package lambda

import (
	"fmt"
	"log"
)

const debug = false

type constraint struct {
	node        AST
	left, right Type
}

type typeSub struct {
	v  int64
	ty Type
}

type typeMap []*typeSub

type tcState struct {
	nextSym int64
	soln    map[int64]*typeSub
}

func (tcs *tcState) gensym() Type {
	n := tcs.nextSym
	tcs.nextSym++

	return &TypeVariable{n}
}

// TypeCheck typechecks an AST and returns the type of the AST
// structure
func TypeCheck(ast AST, env *TypeEnv) (Type, error) {
	tcs := tcState{soln: make(map[int64]*typeSub)}
	return tcs.typeCheck(ast, env)
}

func (tcs *tcState) mapTypes(ty Type) Type {
	return foldType(func(t Type) Type {
		if v, ok := t.(*TypeVariable); ok {
			if ent, ok := tcs.soln[v.Var]; ok {
				mapped := tcs.mapTypes(ent.ty)
				ent.ty = mapped
				return ent.ty
			}
		}
		return t
	}, ty)
}

func (tcs *tcState) addMapping(from int64, to Type) {
	tcs.soln[from] = &typeSub{from, to}
}

func occur(v *TypeVariable, ty Type) bool {
	switch n := ty.(type) {
	case *AtomicType:
		return false
	case *TypeVariable:
		return n.Var == v.Var
	case *FunctionType:
		return occur(v, n.Dom) || occur(v, n.Range)
	case *TupleType:
		for _, e := range n.Elts {
			if occur(v, e) {
				return true
			}
		}
		return false
	default:
		panic(bad("occur", ty))
	}
}

func (tcs *tcState) unify(cs []constraint) error {
	for len(cs) > 0 {
		c := cs[0]
		cs = cs[1:]

		left := tcs.mapTypes(c.left)
		right := tcs.mapTypes(c.right)

		if left == right {
			continue
		}

		if v, ok := left.(*TypeVariable); ok {
			if occur(v, right) {
				return &OccurCheck{c.node}
			}

			tcs.addMapping(v.Var, right)
		} else if v, ok := right.(*TypeVariable); ok {
			if occur(v, left) {
				return &OccurCheck{c.node}
			}
			tcs.addMapping(v.Var, left)
		} else if lf, ok := left.(*FunctionType); ok {
			rf, ok := right.(*FunctionType)
			if !ok {
				return &TypeError{
					Node:     c.node,
					Got:      right,
					Expected: left,
				}
			}
			cs = append(cs, constraint{
				c.node, lf.Dom, rf.Dom,
			}, constraint{
				c.node, lf.Range, rf.Range,
			})
		} else if lt, ok := left.(*TupleType); ok {
			rt, ok := right.(*TupleType)
			if !ok || len(lt.Elts) != len(rt.Elts) {
				return &TypeError{
					Node:     c.node,
					Got:      right,
					Expected: left,
				}
			}
			for i, le := range lt.Elts {
				cs = append(cs, constraint{
					c.node, le, rt.Elts[i],
				})
			}
		} else if la, ok := left.(*AtomicType); ok {
			ra, ok := right.(*AtomicType)
			if !ok || ra.Name != la.Name {
				return &TypeError{
					Node:     c.node,
					Got:      right,
					Expected: left,
				}
			}
		} else {
			panic(fmt.Sprintf("occurs: unexpected lhs: %#v", left))
		}
	}
	return nil
}

func (tcs *tcState) typeCheck(ast AST, env *TypeEnv) (Type, error) {
	ty, cs, err := tcs.constraints(ast, env)
	if err != nil {
		return nil, err
	}
	if debug {
		log.Printf("type: %s", PrintType(ty))
		log.Printf("constraints: ")
		for _, c := range cs {
			log.Printf("  %s = %s", PrintType(c.left), PrintType(c.right))
		}
	}
	for i := range cs {
		cs[i].left = tcs.mapTypes(cs[i].left)
		cs[i].right = tcs.mapTypes(cs[i].right)
	}
	err = tcs.unify(cs)
	if err != nil {
		return nil, err
	}
	mapped := tcs.mapTypes(ty)
	if debug {
		log.Printf("mapped: %s", PrintType(mapped))
	}
	return mapped, nil
}

func (tcs *tcState) constraints(ast AST, env *TypeEnv) (Type, []constraint, error) {
	switch n := ast.(type) {
	case *Boolean:
		return boolType, nil, nil
	case *String:
		return strType, nil, nil
	case *Integer:
		return intType, nil, nil
	case *Variable:
		t := env.Lookup(n.Var)
		if t == nil {
			return nil, nil, UnboundVariable{ast, n.Var}
		}
		return t, nil, nil
	case *Abstraction:
		var names []string
		var types []Type
		for _, v := range n.Vars {
			var argType Type
			tv := v.(*TypedName)
			if tv.Type == nil {
				argType = tcs.gensym()
			} else {
				var e error
				argType, e = ParseType(tv.Type)
				if e != nil {
					return nil, nil, e
				}
			}
			names = append(names, tv.Name)
			types = append(types, argType)
		}
		env := env.Extend(names, types)

		rtype, cs, e := tcs.constraints(n.Body, env)

		if e != nil {
			return nil, nil, e
		}
		return &FunctionType{
			Dom: &TupleType{
				Elts: types,
			},
			Range: rtype,
		}, cs, nil
	case *Application:
		ftype, err := tcs.typeCheck(n.Func, env)
		if err != nil {
			return nil, nil, err
		}

		var args []Type
		for _, a := range n.Args {
			argType, err := tcs.typeCheck(a, env)
			if err != nil {
				return nil, nil, err
			}
			args = append(args, argType)
		}
		argType := &TupleType{Elts: args}
		rng := tcs.gensym()
		constraints := []constraint{{
			node: ast,
			left: ftype,
			right: &FunctionType{
				Dom:   argType,
				Range: rng,
			},
		}}
		return rng, constraints, nil

	case *If:
		cdType, err := tcs.typeCheck(n.Condition, env)
		if err != nil {
			return nil, nil, err
		}
		conType, err := tcs.typeCheck(n.Consequent, env)
		if err != nil {
			return nil, nil, err
		}
		altType, err := tcs.typeCheck(n.Alternate, env)
		if err != nil {
			return nil, nil, err
		}
		constraints := []constraint{
			{ast, boolType, cdType},
			{ast, conType, altType},
		}

		return conType, constraints, nil

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
	if l == r {
		return true
	}
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
	case *TypeVariable:
		rv, ok := r.(*TypeVariable)
		return ok && t.Var == rv.Var
	default:
		panic(fmt.Sprintf("unhandled type: %#v", l))
	}
}
