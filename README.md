# Giggles

## A toy interpreter 
References:
- https://craftinginterpreters.com/
- https://dpvipracollege.ac.in/wp-content/uploads/2023/01/Alfred-V.-Aho-Monica-S.-Lam-Ravi-Sethi-Jeffrey-D.-Ullman-Compilers-Principles-Techniques-and-Tools-Pearson_Addison-Wesley-2007.pdf
- https://youtu.be/ENKT0Z3gldE?si=_dIebGGxXJ5CDZFK
- https://youtu.be/MMxMeX5emUA?si=xtlSl7JEcxy8GSme
- https://youtu.be/SToUyjAsaFk?si=ktX25YS_9reX7E05

No args → repl() (interactive console)
One arg → runFile() (run a .lox script)
Anything else → print usage error

#### A tiny Ballerina-inspired compiler/interpreter in Go for learning compiler frontend, type checking, IR generation, and interpretation.
##### (It compiles to IR and interpretes it through Go)

+ Go implementation
+ lua like syntax
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
##### (lua inspired syntax)

```
-- numbers
int age = 21
int price = 99.50
int total = age + price

-- strings
string name = "Oshan"
string message = "Hello, " .. name

-- booleans
bool isReady = true
bool isdone = false

if isready then
  print("ready")
end

-- nil
int value = nil

if value == nil then
  print("nothing here")
end

-- variables
int mut x = 10
x = x + 5

string langName = "LuLox"

-- functions
function add(a, b)
  return a + b
end

print(add(2, 3))

```


