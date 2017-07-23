package gollum

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
	v  *TypeVariable
	ty Type
}

type tcState struct {
	nextSym int64
	soln    map[*TypeVariable]*typeSub
}

func (tcs *tcState) gensym() Type {
	n := tcs.nextSym
	tcs.nextSym++

	return &TypeVariable{n, nil}
}

// TypeCheck typechecks an AST and returns the type of the AST
// structure
func TypeCheck(ast AST, env *TypeEnv) (Type, error) {
	tcs := tcState{soln: make(map[*TypeVariable]*typeSub)}
	ty, err := tcs.typeCheck(ast, env)
	if ty != nil {
		ty = tcs.mapTypes(ty)
	}
	return ty, err
}

func (tcs *tcState) mapTypes(ty Type) Type {
	return mapVars(func(v *TypeVariable) Type {
		if ent, ok := tcs.soln[v]; ok {
			mapped := tcs.mapTypes(ent.ty)
			ent.ty = mapped
			return ent.ty
		}
		return v
	}, ty)
}

func (tcs *tcState) addMapping(from *TypeVariable, to Type) {
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

func (tcs *tcState) generalize(ty Type, e *TypeEnv) Type {
	bound := make(map[*TypeVariable]struct{})
	for _, b := range e.BoundVars() {
		bound[b] = struct{}{}
	}
	free := make(map[*TypeVariable]struct{})
	mapVars(func(tv *TypeVariable) Type {
		if _, ok := bound[tv]; !ok {
			free[tv] = struct{}{}
		}
		return tv
	}, ty)
	var quantify []*TypeVariable
	for f := range free {
		quantify = append(quantify, f)
	}
	return &Forall{
		Vars: quantify,
		Type: ty,
	}
}

func (tcs *tcState) instantiate(ty Type) Type {
	forall, ok := ty.(*Forall)
	if !ok {
		return ty
	}
	rename := make(map[*TypeVariable]Type, len(forall.Vars))
	for _, v := range forall.Vars {
		rename[v] = tcs.gensym()
	}
	return mapVars(func(tv *TypeVariable) Type {
		if newv, ok := rename[tv]; ok {
			return newv
		}
		return tv
	}, forall.Type)
}

func (tcs *tcState) unify(cs []constraint) error {
	for len(cs) > 0 {
		c := cs[0]
		cs = cs[1:]

		left := tcs.mapTypes(c.left)
		right := tcs.mapTypes(c.right)

		if debug {
			log.Printf("unify %s = %s | %s = %s",
				PrintType(c.left), PrintType(c.right),
				PrintType(left), PrintType(right),
			)
		}

		if left == right {
			continue
		}

		if v, ok := left.(*TypeVariable); ok {
			if occur(v, right) {
				return &OccurCheck{c.node, left, right}
			}

			tcs.addMapping(v, right)
		} else if v, ok := right.(*TypeVariable); ok {
			if occur(v, left) {
				return &OccurCheck{c.node, left, right}
			}
			tcs.addMapping(v, left)
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

func syntacticValue(ast AST) bool {
	_, ok := ast.(*Abstraction)
	return ok
}

func (tcs *tcState) typeCheck(ast AST, env *TypeEnv) (Type, error) {
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
		return tcs.instantiate(t), nil
	case *Abstraction:
		var names []string
		var types []Type
		var bound []*TypeVariable

		for _, v := range n.Vars {
			var argType Type
			tv := v.(*TypedName)
			if tv.Type == nil {
				argType = tcs.gensym()
				bound = append(bound, argType.(*TypeVariable))
			} else {
				var e error
				argType, e = ParseType(tv.Type)
				if e != nil {
					return nil, e
				}
			}
			names = append(names, tv.Name)
			types = append(types, argType)
		}
		env := env.Extend(names, types, bound)

		rtype, e := tcs.typeCheck(n.Body, env)

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
		ftype, err := tcs.typeCheck(n.Func, env)
		if err != nil {
			return nil, err
		}

		var args []Type
		for _, a := range n.Args {
			argType, err := tcs.typeCheck(a, env)
			if err != nil {
				return nil, err
			}
			args = append(args, argType)
		}
		argType := &TupleType{Elts: args}
		rng := tcs.gensym()
		if err := tcs.unify([]constraint{{
			node: ast,
			left: ftype,
			right: &FunctionType{
				Dom:   argType,
				Range: rng,
			},
		}}); err != nil {
			return nil, err
		}
		return rng, nil

	case *If:
		cdType, err := tcs.typeCheck(n.Condition, env)
		if err != nil {
			return nil, err
		}
		conType, err := tcs.typeCheck(n.Consequent, env)
		if err != nil {
			return nil, err
		}
		altType, err := tcs.typeCheck(n.Alternate, env)
		if err != nil {
			return nil, err
		}
		if err := tcs.unify([]constraint{
			{ast, boolType, cdType},
			{ast, conType, altType},
		}); err != nil {
			return nil, err
		}

		return conType, nil
	case *Let:
		parent := env
		var names []string
		var types []Type
		var vars []*TypeVariable
		for _, b := range n.Bindings {
			nb := b.(*NameBinding)
			tn := nb.Var.(*TypedName)
			ty := tcs.gensym()
			vars = append(vars, ty.(*TypeVariable))
			names = append(names, tn.Name)
			types = append(types, ty)
		}

		if n.Recursive {
			env = env.Extend(names, types, vars)
		}

		for i, b := range n.Bindings {
			nb := b.(*NameBinding)
			tn := nb.Var.(*TypedName)
			var ty Type
			vty, err := tcs.typeCheck(nb.Value, env)
			if err != nil {
				return nil, err
			}
			if err := tcs.unify([]constraint{
				{nb, types[i], vty},
			}); err != nil {
				return nil, err
			}
			if tn.Type != nil {
				var err error
				ty, err = ParseType(tn.Type)
				if err != nil {
					return nil, err
				}
				if err := tcs.unify([]constraint{
					{nb, ty, vty},
				}); err != nil {
					return nil, err
				}
			}
			if syntacticValue(nb.Value) {
				vty = tcs.generalize(vty, parent)
			}
			if debug {
				log.Printf("let %s : %s", tn.Name, PrintType(vty))
			}
			types[i] = vty
		}
		if n.Recursive {
			env.SetLocal(names, types)
		} else {
			env = env.Extend(names, types, vars)
		}
		bty, err := tcs.typeCheck(n.Body, env)
		if err != nil {
			return nil, err
		}
		return bty, nil

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
		return ok && t == rv
	default:
		panic(fmt.Sprintf("unhandled type: %#v", l))
	}
}
