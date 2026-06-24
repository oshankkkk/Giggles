
package compiler

import (
	"lang/internal/parser"
)

type Scope struct {
	Parent *Scope
	localsymbols map[string]parser.Symbols

}

func InitScope() *Scope {
	return &Scope{
		Parent: nil,
		localsymbols: make(map[string]parser.Symbols),
	}
}

func EnterScope(parent *Scope,) *Scope {
	return &Scope{
		Parent: parent,
		localsymbols: make(map[string]parser.Symbols),
	}
}

func (s *Scope) LeaveScope() *Scope {
	if s.Parent == nil {
		panic("cannot leave the global scope")
	}
	return s.Parent
}

func (s *Scope) AddSymbol(symbol parser.Symbols) {
	name:=symbol.GetName()
//	fmt.Println(symbol.GetName(),"this is the function name" )
    s.localsymbols[name] = symbol

}
//for identifier to check if symbol are parsed
func (s *Scope) VarLookup(name string) (parser.Symbols, bool) {
	if v, ok := s.localsymbols[name]; ok {
		return v, true
	}
//
//	if s.Parent != nil {
//		return s.Parent.VarLookup(name)
//	}
	return nil, false
}
func (s *Scope) GlobalVarLookup(name string) (parser.Symbols, bool) {
	if v, ok := s.localsymbols[name]; ok {
		return v, true
	}

	if s.Parent != nil {
		return s.Parent.VarLookup(name)
	}

	return nil, false
}



