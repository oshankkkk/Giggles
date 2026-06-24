package compiler

import (
	"fmt"
	//"lang/internal/compiler"
	//	"fmt"
	"lang/internal/lexer"
	"lang/internal/parser"
	"strconv"
)
type State struct {
	CounterTable  []int
	fixups []fixup
	Entrypoint int
	Buff []byte
	globals map[string]parser.VarDecl
	//globals []parser.VarDecl

}

func (c *State) Wrapper(ast parser.ASTNode,scope *Scope){
	c.globals=make(map[string]parser.VarDecl)
	c.Buff = append(c.Buff, byte(JMP),0)	
	c.ToBytes(ast,scope)
	c.Buff[1]=byte(c.Entrypoint)
	if err:=c.Fixpatchs();err!=nil{
	panic(err)
}
	
}
func (c *State) ToBytes(ast parser.ASTNode, scope *Scope) {
	if value, ok := ast.(parser.Program); ok {
		for _, val := range value.Statements {
			c.ToBytes(val, scope)
		}
	}
	if value, ok := ast.(parser.ExprStatement); ok {
		c.ToBytes(value.Expr, scope)
	}
	if value, ok := ast.(parser.Groups); ok {
		c.ToBytes(value.Value, scope)
	}
//if value, ok := ast.(parser.VarDecl); ok {
//    c.ToBytes(value.Value, scope)
//    value.Address = len(c.Buff)
//
//    if !value.IsLocal {
//        // globals unchanged
//        c.globals[value.GetName()] = value
//        c.CounterTable = append(c.CounterTable, value.Address)
//        c.Buff = append(c.Buff, byte(SETGLOBAL), byte(len(c.CounterTable)-1))
//    } else {
//        if existing, ok := scope.VarLookup(value.GetName()); ok {
//            // REASSIGNMENT — slot already exists, just write into it
//            c.CounterTable = append(c.CounterTable, existing.GetAddress())
//            c.Buff = append(c.Buff, byte(SETLOCAL), byte(len(c.CounterTable)-1))
//        } else {
//            // DECLARATION — allocate new slot
//            scope.AddSymbol(value)
//            c.CounterTable = append(c.CounterTable, value.Address) // <-- was missing
//            c.Buff = append(c.Buff, byte(SETLOCAL), byte(len(c.CounterTable)-1))
//        }
//    }
//}
	//assign values, and also declar them
	if value, ok := ast.(parser.VarDecl); ok {
		c.ToBytes(value.Value, scope)
		value.Address=len(c.Buff)
		c.CounterTable = append(c.CounterTable, value.Address)
		if !value.IsLocal{
			c.globals[value.GetName()]=value
			//the address of the push
			c.Buff = append(c.Buff, byte(SETGLOBAL), byte(len(c.CounterTable)-1))
		}else{
			scope.AddSymbol(value)
			c.Buff = append(c.Buff, byte(SETLOCAL), byte(len(c.CounterTable)-1))
		} 
	}

	if value, ok := ast.(parser.Identifier); ok {
		if info, ok := scope.VarLookup(value.Name.Value); ok {
				//the address of the push
			c.CounterTable = append(c.CounterTable,info.GetAddress())
			c.Buff = append(c.Buff, byte(GETLOCAL), byte(len(c.CounterTable)-1))
		}else if value,ok:=c.globals[value.Name.Value];ok{
			c.CounterTable = append(c.CounterTable, value.GetAddress())
			c.Buff = append(c.Buff, byte(GETGLOBAL), byte(len(c.CounterTable)-1))
		}else {
			panic("var undefined")
		}
	}
	if value, ok := ast.(parser.Call); ok {
		c.Buff = append(c.Buff, byte(NWFRM))
		// Push return address (instruction after the JMP)
		returnAddr := len(c.Buff) + 4 // address of instruction after JMP+operand
		c.CounterTable = append(c.CounterTable, returnAddr)
		c.Buff = append(c.Buff, byte(PUSH), byte(len(c.CounterTable)-1))
		fmt.Println("pushing",len(c.Buff)+5 )
		funcjmp := len(c.Buff) + 1
		c.Buff = append(c.Buff, byte(JMP), 0)
		if function, ok := scope.VarLookup(value.Function); ok {
			c.CounterTable = append(c.CounterTable,function.GetAddress())
			fmt.Println(function.GetAddress(),"<-- address in foo")
			fmt.Println(c.CounterTable,"this is the CounterTable")
			c.Buff[funcjmp] = byte(len(c.CounterTable) - 1)
		} else {

			fmt.Println(c.CounterTable,"this is the CounterTable")
			c.fixups = append(c.fixups, fixup{address: funcjmp, funname: value.Function, scopeNode: scope})
		}
	}
	if value, ok := ast.(parser.Function); ok {
		if value.Ismain{
			c.Entrypoint = len(c.CounterTable)
			c.CounterTable = append(c.CounterTable, len(c.Buff))
			c.Buff = append(c.Buff, byte(NWFRM))
				for _, n := range value.Content {
			c.ToBytes(n,scope)
		}
			c.Buff= append(c.Buff, byte(STOP))
		}else{
		value.Address=len(c.Buff)
		fmt.Println("address of foo",value.Address)
		scope.AddSymbol(value)
		local:=EnterScope(scope)
		//c.symboltb[value.Name] = len(c.Buff)
		for _, n := range value.Content {
			c.ToBytes(n,local)
		}
		c.Buff = append(c.Buff, byte(RMFRM))
	} 
}
	if value, ok := ast.(parser.Condition); ok {
		condPos := len(c.Buff)
		c.ToBytes(value.Condition, scope)
		jifPos := len(c.Buff)
		c.Buff = append(c.Buff, byte(JIF), 0)
		for _, r := range value.Result {
			c.ToBytes(r, scope)
		}
		if value.Looped {
			c.Buff = append(c.Buff, byte(JMP), byte(condPos))
			c.CounterTable = append(c.CounterTable, len(c.Buff))
			c.Buff[jifPos+1] = byte(len(c.CounterTable) - 1)
			for _, e := range value.ElseResult {
				c.ToBytes(e, scope)
			}
		} else if value.HasElse {
			jmpPos := len(c.Buff)
			c.Buff = append(c.Buff, byte(JMP), 0)
			c.CounterTable = append(c.CounterTable, len(c.Buff))
			c.Buff[jifPos+1] = byte(len(c.CounterTable) - 1)
			for _, e := range value.ElseResult {
				c.ToBytes(e, scope)
			}
			c.CounterTable = append(c.CounterTable, len(c.Buff))
			c.Buff[jmpPos+1] = byte(len(c.CounterTable) - 1)
		} else {
			c.CounterTable = append(c.CounterTable, len(c.Buff))
			c.Buff[jifPos+1] = byte(len(c.CounterTable) - 1)
		}
	}
	if value, ok := ast.(parser.Unary); ok {
		if value.Operator == lexer.MINUS {
			// Compile as 0 - value
			c.CounterTable = append(c.CounterTable, 0)
			c.Buff = append(c.Buff, byte(PUSH), byte(len(c.CounterTable)-1))
			c.ToBytes(value.Value, scope)
			c.Buff = append(c.Buff, byte(SUB))
		} else if value.Operator == lexer.NOT {
			// Compile as value == false
			c.ToBytes(value.Value, scope)
			c.Buff = append(c.Buff, byte(FALSE))
			c.Buff = append(c.Buff, byte(EQ))
		}
	}
	if value, ok := ast.(parser.Literal); ok {
		if value.Value.Type == lexer.TRUE {
			c.Buff = append(c.Buff, byte(TRUE))
			return
		} else if value.Value.Type == lexer.FALSE {
			c.Buff = append(c.Buff, byte(FALSE))
			return
		}
		index, err := strconv.Atoi(value.Value.Value)
		if err != nil {
			panic(err)
		}
		c.CounterTable = append(c.CounterTable, index)
		c.Buff = append(c.Buff, byte(PUSH), byte(len(c.CounterTable)-1))
	}
	if value, ok := ast.(parser.Binary); ok {

		if value.Operator == lexer.EQUAL {

			fmt.Println("ealalal{")
			c.ToBytes(value.Right, scope)
			if variable, ok := scope.VarLookup(value.Left.(parser.Identifier).Name.Value); ok {
				c.CounterTable = append(c.CounterTable, variable.GetAddress())
				c.Buff = append(c.Buff, byte(ASS), byte(len(c.CounterTable)-1))
			} else {
				panic("var undefined")
			}
			return
		}
		c.ToBytes(value.Left, scope)
		c.ToBytes(value.Right, scope)
		var opcode Opcode
		switch value.Operator {
		case lexer.MINUS:
			opcode = SUB
		case lexer.PLUS:
			opcode = ADD
		case lexer.SLASH:
			opcode = DIV
		case lexer.STAR:
			opcode = MUL
		case lexer.AND:
			opcode = AND
		case lexer.OR:
			opcode = OR
		case lexer.GREATER:
			opcode = GT
		case lexer.LESS:
			opcode = LT
		case lexer.GREATER_EQUAL:
			opcode = GTE
		case lexer.LESS_EQUAL:
			opcode = LTE
		case lexer.EQUAL_EQUAL:
			opcode = EQ
		case lexer.NOT_EQUAL:
			opcode = NEQ
		}
		c.Buff = append(c.Buff, byte(opcode))
		//fmt.Println(len(c.Buff)-1,"add address")
	}
}

type fixup struct{
	address int
	funname string
	scopeNode *Scope
}
func (c *State) Fixpatchs()error{
	// After the whole program is compiled, resolve all fixups
	for _, f := range c.fixups {
		symbol, ok := f.scopeNode.VarLookup(f.funname)
		if !ok {
			return fmt.Errorf("undefined function: %s", f.funname)
		}
		c.CounterTable = append(c.CounterTable, symbol.GetAddress())

			fmt.Println(c.CounterTable,"this is the CounterTable")
		c.Buff[f.address] =byte(len(c.CounterTable)-1)
				fmt.Println("f address",f.address)
	}
	return nil

}

// Write the real address into the 2-byte placeholder
// Adjust encoding to match your VM's word size / endianness
//c.Buff[f.address]   = byte(addr)   // high byte

//c.Buff[f.patchOffset+1] = byte(addr & 0xFF) // low byte
