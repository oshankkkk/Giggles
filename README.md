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

## Crafting interpreters with lux language 
#### Features
- numbers
- strings
- booleans
- nil
- variables
- functions
- classes
- instances
- methods
- inheritance

# What we build: lua flavoured lux language

#### Syntax:

```lua
-- numbers
local age = 21
local price = 99.50
local total = age + price

-- strings
local name = "Oshan"
local message = "Hello, " .. name

-- booleans
local isReady = true
local isDone = false

if isReady then
  print("ready")
end

-- nil
local value = nil

if value == nil then
  print("nothing here")
end

-- variables
local x = 10
x = x + 5

local langName = "Lua-ish Lox"

-- functions
function add(a, b)
  return a + b
end

print(add(2, 3))

-- classes (Lua-style using tables + metatables)
Person = {}
Person.__index = Person

function Person.new(name)
  local self = setmetatable({}, Person)
  self.name = name
  return self
end

-- instances
local p = Person.new("Oshan")
print(p.name)

-- methods
function Person:sayHello()
  print("Hello, " .. self.name)
end

p:sayHello()

-- inheritance
Animal = {}
Animal.__index = Animal

function Animal.new(name)
  local self = setmetatable({}, Animal)
  self.name = name
  return self
end

function Animal:speak()
  print(self.name .. " makes a sound")
end

Dog = setmetatable({}, { __index = Animal })
Dog.__index = Dog

function Dog.new(name)
  local self = setmetatable(Animal.new(name), Dog)
  return self
end

function Dog:speak()
  print(self.name .. " barks")
end

local d = Dog.new("Rex")
d:speak()
```


## What is a lexer
***lexer was just sequentially going through the characters and categorizing them into groups every time it finds a break point (an invalid character, space, operator, etc).***

## Why a token

- What kind of thing is this?
- What exact text did it come from?
- Where was it in the source file?
- What value does it represent, if any?

## Token we should build

```
LEFT_PAREN      (
RIGHT_PAREN     )
LEFT_BRACE      {
RIGHT_BRACE     }
LEFT_BRACKET    [
RIGHT_BRACKET   ]
COMMA           ,
DOT             .
SEMICOLON       ;
COLON           :

```

```
PLUS            +
MINUS           -
STAR            *
SLASH           /
PERCENT         %

```

```
EQUAL           =
EQUAL_EQUAL     ==
BANG            !
BANG_EQUAL      !=
LESS            <
LESS_EQUAL      <=
GREATER         >
GREATER_EQUAL   >=

```

```
IDENTIFIER      variableName
NUMBER          123, 3.14
STRING          "hello"
```

```
IF
ELSE
WHILE
FOR
RETURN
FUNCTION
VAR
TRUE
FALSE
NULL
EOF
```

































