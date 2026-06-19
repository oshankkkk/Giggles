package compiler
import "fmt"
type Variable struct {
	id int
	Name  string
	//Type 
}

type Scope struct {
	Parent *Scope
	locals map[string]Variable
	varcounter int
	//startpointer int
}

// InitScope creates a brand new root scope (no parent)
func InitScope() *Scope {
	return &Scope{
		Parent: nil,
		locals: make(map[string]Variable),
		//startpointer: 0,
	}
}

// EnterScope creates a child scope — call this when entering a block/function
func EnterScope(parent *Scope,) *Scope {
	return &Scope{
		Parent: parent,
		locals: make(map[string]Variable),
	}
}

// LeaveScope returns the parent scope — call this when exiting a block/function
func (s *Scope) LeaveScope() *Scope {
	if s.Parent == nil {
		panic("cannot leave the global scope")
	}
	return s.Parent
}

// AddVariable defines a variable in the CURRENT scope only
func (s *Scope) AddVariable(name string) {
	
	s.locals[name] = Variable{Name: name, id: s.varcounter}
	fmt.Println("idid",s.varcounter)

	s.varcounter++
}

// VarLookup walks up the scope chain to find a variable by name
func (s *Scope) VarLookup(name string) (Variable, bool) {
	// Check current scope first
	if v, ok := s.locals[name]; ok {
		return v, true
	}
	// Not here — ask the parent (this is the key recursive step)
	if s.Parent != nil {
		return s.Parent.VarLookup(name)
	}
	// Reached the root and still not found
	return Variable{}, false
}


