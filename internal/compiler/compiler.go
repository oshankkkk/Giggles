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
	symboltb map[string]int
	fixups []fixup
	Entrypoint int
}
type fixup struct{
	address int
	funname string
}
func (c *State) Wrapper(ast parser.ASTNode,scope *Scope)[]byte{
	c.symboltb=make(map[string]int)
	var bytelist []byte
	b:=c.ToBytes(ast,scope)
	bytelist = append(bytelist, byte(JMP),byte(c.Entrypoint))
	bytelist=append(bytelist,b...)	
	bytelist = append(bytelist, byte(STOP))
	return  bytelist
}
func (c *State) ToBytes(ast parser.ASTNode,scope *Scope) []byte{
	var bytecode []byte
	
	if value, ok := ast.(parser.Program); ok {
		for _, val := range value.Statements {
			bytecode = append(bytecode, c.ToBytes(val,scope)...)
		}
	}
	if value, ok := ast.(parser.ExprStatement); ok {
		bytecode = append(bytecode, c.ToBytes(value.Expr,scope)...)
	}
	if value, ok := ast.(parser.Groups); ok {
		bytecode = append(bytecode, c.ToBytes(value.Value,scope)...)
	}

	if value, ok := ast.(parser.VarDecl); ok {
		//check for scope and put the flat bytes here with pointer address and var table index value	
		bytecode = append(bytecode, c.ToBytes(value.Value,scope)...)
		// the variable name is after vardec
		v:=scope.AddVariable(value.Name.Value)

		c.CounterTable=append(c.CounterTable,v.id )	
		bytecode = append(bytecode, byte(SETGLOBAL),byte(len(c.CounterTable)-1))
		//		bytecode = append(bytecode, byte(SETGLOBAL))
		//what this opcode does
	}

	if value, ok := ast.(parser.Identifier); ok {
		fmt.Println("ee343")
		//get local/global
		if info,ok:=scope.VarLookup(value.Name.Value);ok{
			//this means the var is in the vm now

			c.CounterTable=append(c.CounterTable,info.id )	
			fmt.Println(c.CounterTable)	
			bytecode = append(bytecode, byte(GETGLOBAL),byte(len(c.CounterTable)-1))

		}else{

			panic("var undefined")
		}
	}

	if value, ok := ast.(parser.Call); ok {
		bytecode = append(bytecode, byte(NWFRM))

		funcjmp:=len(bytecode)+1
		bytecode = append(bytecode, byte(JMP),0)

		if fnaddress,ok:=c.symboltb[value.Function];ok{

		c.CounterTable=append(c.CounterTable,fnaddress)	
		bytecode[funcjmp] =  byte(len(c.CounterTable)-1)
			
		}else{
			c.fixups = append(c.fixups, fixup{address: funcjmp,funname: value.Function})	
		}
		// 0 is place holder for jmp
		//we check for the function if its not there, we will keep in touch.
	}
	if value, ok := ast.(parser.Function); ok {
		if value.Name==string(lexer.MAIN){
		c.Entrypoint=len(bytecode)
		}
		c.symboltb[value.Name]=len(bytecode)
		for _,n:=range value.Content{
			bytecode = append(bytecode, c.ToBytes(n,scope)...)
		}
		bytecode = append(bytecode, byte(RMFRM))
	}


	if value, ok := ast.(parser.Condition); ok {
		condPos := len(bytecode)
		bytecode = append(bytecode, c.ToBytes(value.Condition,scope)...)
		jifPos := len(bytecode)

		//placeholder 0 
		bytecode = append(bytecode, byte(JIF), 0)

		resultCode := []byte{}
		//local :=EnterScope(scope)
		for _, r := range value.Result {
			resultCode = append(resultCode, c.ToBytes(r,scope)...)
		}
		bytecode = append(bytecode, resultCode...)
		if value.Looped {
			bytecode = append(bytecode, byte(JMP), byte(condPos))
			elseCode := []byte{}
			//local :=EnterScope(scope)
			for _, e := range value.ElseResult {
				elseCode = append(elseCode, c.ToBytes(e,scope)...)
			}
			c.CounterTable=append(c.CounterTable,len(bytecode) )	
			bytecode[jifPos+1] = byte(len(c.CounterTable)-1)

			bytecode = append(bytecode, elseCode...)
		} else if value.HasElse {
			elseCode := []byte{}

			for _, e := range value.ElseResult {
				elseCode = append(elseCode, c.ToBytes(e,scope)...)
			}
			jmpPos := len(bytecode)

			bytecode = append(bytecode, byte(JMP), 0)

			c.CounterTable=append(c.CounterTable,len(bytecode) )	
			bytecode[jifPos+1] = byte(len(c.CounterTable)-1)

			bytecode = append(bytecode, elseCode...)

			c.CounterTable=append(c.CounterTable,len(bytecode) )	
			bytecode[jmpPos+1] = byte(len(c.CounterTable)-1)

		} else {
			c.CounterTable=append(c.CounterTable,len(bytecode) )	
			bytecode[jifPos+1] = byte(len(c.CounterTable)-1)

		}
	}

	if value, ok := ast.(parser.Literal); ok {
		if value.Value.Type == lexer.TRUE {
			return append(bytecode, byte(TRUE))
		} else if value.Value.Type == lexer.FALSE {
			return append(bytecode, byte(FALSE))
		}
		index,err:=strconv.Atoi(value.Value.Value)
		if err!=nil{
			panic(err)
		}
		c.CounterTable = append(c.CounterTable, index)
		return append(bytecode, byte(PUSH), byte(len(c.CounterTable)-1))
	}
	// ... rest of binary ops
	if value, ok := ast.(parser.Binary); ok {
		if value.Operator == lexer.EQUAL {
			// For assignment: only compile the right-hand side
			bytecode = append(bytecode, c.ToBytes(value.Right, scope)...)
			if info, ok := scope.VarLookup(value.Left.(parser.Identifier).Name.Value); ok {
				c.CounterTable = append(c.CounterTable, info.id)
				bytecode = append(bytecode, byte(ASS), byte(len(c.CounterTable)-1))
			} else {
				panic("var undefined")
			}
			return bytecode
		}

		bytecode = append(bytecode, c.ToBytes(value.Left,scope)...)
		bytecode = append(bytecode, c.ToBytes(value.Right,scope)...)
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

		bytecode = append(bytecode, byte(opcode))
	}
	return bytecode
}
func (c *State) Fixpatchs(bytecode []byte)(error,[]byte){
	// After the whole program is compiled, resolve all fixups
	for _, f := range c.fixups {
		addr, ok := c.symboltb[f.funname]
		if !ok {
			return fmt.Errorf("undefined function: %s", f.funname),make([]byte,1)
		}
		// Write the real address into the 2-byte placeholder
		// Adjust encoding to match your VM's word size / endianness
		bytecode[f.address]   = byte(addr)   // high byte
		//bytecode[f.patchOffset+1] = byte(addr & 0xFF) // low byte
	}
	return nil,bytecode
}


//		if value.IsLocal{
//			bytecode = append(bytecode, byte(SETLOCAL))
//			fmt.Println("heheh")
//		}else{

//	if value, ok:=ast.(parser.Call);ok{
//		bytecode = append(bytecode, byte(CALLFUNC))
//	}
//	if value, ok:=ast.(parser.Function);ok{
		//local :=EnterScope(scope)
		//start:=len(bytecode)
		//framestart:=len(bytecode)
		//bytecode = append(bytecode,byte( ENTERFUNC))
	//	for _,stmnt:=range value.Content{
		//bytecode = append(bytecode, c.ToBytes(stmnt,scope)...)
		//	}
			//end:=len(bytecode)
		//c.CounterTable=append(c.CounterTable,end )	
		//fmt.Println(framestart,"dddrfff")
		//fmt.Println(len(c.CounterTable)-2,"ffffffff")

		//bytecode = append(bytecode, byte(JMP), byte(len(c.CounterTable)-2))
	//}
