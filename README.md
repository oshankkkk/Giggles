# A tiny Ballerina-inspired compiler/interpreter in Go for learning compiler frontend, type checking, IR generation, and interpretation.
## (It compiles to IR and interpretes it through Go)

### How this works
lexer
→ parser
→ AST
→ symbol table
→ type checker
→ tiny IR
→ interpreter

### What we have
- int
- bool
- string
- variables
- functions
- if
- while
- return

#### Timeline
Phase 1: lexer + parser + AST
Phase 2: symbol table + scopes
Phase 3: type checker
Phase 4: tiny IR
Phase 5: IR interpreter
Phase 6: records / optional fields / tables

#####  From hydrogen lang tut to: 
+ Go implementation
+ Ballerina-like syntax
+ symbol/type system
+ IR interpreter instead of assembly

