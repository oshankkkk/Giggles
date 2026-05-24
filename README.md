## References:
- https://craftinginterpreters.com/
- https://dpvipracollege.ac.in/wp-content/uploads/2023/01/Alfred-V.-Aho-Monica-S.-Lam-Ravi-Sethi-Jeffrey-D.-Ullman-Compilers-Principles-Techniques-and-Tools-Pearson_Addison-Wesley-2007.pdf
- https://youtu.be/ENKT0Z3gldE?si=_dIebGGxXJ5CDZFK
- https://youtu.be/MMxMeX5emUA?si=xtlSl7JEcxy8GSme
- https://youtu.be/SToUyjAsaFk?si=ktX25YS_9reX7E05

#### A tiny Ballerina-inspired compiler/interpreter in Go for learning compiler frontend, type checking, IR generation, and interpretation.
##### (It compiles to IR and interpretes it through Go)

+ Go implementation
+ Ballerina-like syntax
+ symbol/type system
+ IR interpreter instead of assembly

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

#### Syntax:
##### (lua inspired)
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

## Expression Grammer
expression → literal
            | unary
            | binary
            | grouping ;

literal → NUMBER | STRING | "true" | "false" | "nil" ;
grouping → "(" expression ")" ;
unary → ( "-" | "!" ) expression ;
binary → expression operator expression ;
operator → "==" | "!=" | "<" | "<=" | ">" | ">="
| "+" | "-" | "*" | "/" ;

### Grammar so ALU ops
Literal1, 2, 3, 4, false — raw values
Grouping(2 * 3) — anything in parentheses
Binary*, -, <, == — anything with a left and right side
Unary-1 or !true — one operator, one value
