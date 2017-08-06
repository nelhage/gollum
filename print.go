package gollum

import (
	"bytes"
	"fmt"
	"io"
)

// PrintType returns a string representation of a type
func PrintType(t Type) string {
	uvars := make(map[string]struct{})
	gvars := make(map[*TypeVariable]struct{})
	var gvo []*TypeVariable

	eachVar(func(v *TypeVariable) {
		if v.Name != "" {
			uvars[v.Name] = struct{}{}
		} else if _, ok := gvars[v]; !ok {
			gvars[v] = struct{}{}
			gvo = append(gvo, v)
		}
	}, t)

	names := make(map[*TypeVariable]string)
	var i int64
	for _, v := range gvo {
		var n string
		for {
			n = base26name(i)
			i++
			if _, ok := uvars[n]; !ok {
				break
			}
		}
		names[v] = n
	}

	var buf bytes.Buffer
	p := printer{&buf, names}
	p.print(t)
	return buf.String()
}

type printer struct {
	buf   io.Writer
	names map[*TypeVariable]string
}

func (p *printer) write(s string) {
	io.WriteString(p.buf, s)
}

func (p *printer) print(t Type) {
	switch n := t.(type) {
	case *AtomicType:
		p.write(n.Name)
	case *FunctionType:
		dom := n.Dom
		if dtup, ok := n.Dom.(*TupleType); ok && len(dtup.Elts) == 1 {
			dom = dtup.Elts[0]
		}

		if _, ok := dom.(*FunctionType); ok {
			p.write("(")
			p.print(dom)
			p.write(")")
		} else {
			p.print(dom)
		}
		p.write(" -> ")
		p.print(n.Range)
	case *TupleType:
		p.write("(")
		for i, e := range n.Elts {
			p.print(e)
			if i != len(n.Elts)-1 {
				p.write(", ")
			}
		}
		p.write(")")
	case *TypeVariable:
		if n.Name != "" {
			p.write(n.Name)
		} else if name, ok := p.names[n]; ok {
			p.write(name)
		} else {
			panic(fmt.Sprintf("unexpected variable: %#v", n))
		}
	case *Forall:
		p.write("∀")
		for i, v := range n.Vars {
			p.print(v)
			if i != len(n.Vars)-1 {
				p.write(",")
			}
		}
		p.write(".")
		p.print(n.Type)
	default:
		panic(fmt.Sprintf("unknown type: %#v", t))
	}
}

func base26name(v int64) string {
	// ceil(log_26(2**64))
	var buf [14]byte
	i := len(buf)

	for {
		i--
		buf[i] = byte('a' + (v % 26))
		v /= 26
		if v == 0 {
			break
		}
	}

	return string(buf[i:])
}
