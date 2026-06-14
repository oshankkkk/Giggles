# Giggles

#### TODOS
- type system
- string support
- if-then syntax sugar with end
- make all variables cons by default
- make all var global by default
- local keyword for local vars
- scope handling
- func support
- print() and scanner() functions in stdlib


## A toy interpreter 
References:
- find my notes on https://oshanswe.com/
- https://craftinginterpreters.com/
- https://dpvipracollege.ac.in/wp-content/uploads/2023/01/Alfred-V.-Aho-Monica-S.-Lam-Ravi-Sethi-Jeffrey-D.-Ullman-Compilers-Principles-Techniques-and-Tools-Pearson_Addison-Wesley-2007.pdf
- https://youtu.be/ENKT0Z3gldE?si=_dIebGGxXJ5CDZFK
- https://youtu.be/MMxMeX5emUA?si=xtlSl7JEcxy8GSme
- https://youtu.be/SToUyjAsaFk?si=ktX25YS_9reX7E05

No args → repl() (interactive console)
One arg → runFile() (run a .lox script)
Anything else → print usage error


#### It compiles to bytecode interpretes it through Go
+ Go implementation
+ lua like synta (ish)
+ symbol/type system
+ IR VM interpreter 

#### Features
- numbers
- strings
- booleans
- nil
- variables
- functions
--- later --- 
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

### Features i wanna implement eventually
#### (borrowed from project dreambird)
Write five or more equals signs to start a new file. This removes the need for multiple files or any build process.

```
const const score = 5!
print(score)! //5

=====================

const const score = 3!
print(score)! //3

New for 2022!
Thanks to recent advances in technology, you can now give files names.

======= add.gom =======
function add(a, b) => {
   return a + b!
}

```

Many languages allow you to import things from specific files. In GulfOfMexico, importing is simpler. Instead, you export to specific files!

```
===== add.gom ==
function add(a, b) => {
   return a + b!
}

export add to "main.gom"!

===== main.gom ==
import add!
add(3, 2)!

```


For maximum compatibility with other languages, you can also use the className keyword when making classes.
```
This makes things less complicated.

className Player {
   const var health = 10!
}

```

In response to some recent criticism about this design decision, we would like to remind you that this is part of the JavaScript specification, and therefore — out of our control.
```

Gulf of Mexico features AEMI, which stands for Automatic-Exclamation-Mark-Insertion. If you forget to end a statement with an exclamation mark, Gulf of Mexico will helpfully insert one for you!

print("Hello world") // This is fine

Similarly... Gulf of Mexico also features ABI, which stands for Automatic-Bracket-Insertion. If you forget to close your brackets, Gulf of Mexico will pop some in for you!

print("Hello world" // This is also fine

Similarly.... Gulf of Mexico also features AQMI, which stands for Automatic-Quotation-Marks-Insertion. If you forget to close your string, Gulf of Mexico will do it for you!

print("Hello world // This is fine as well

This can be very helpful in callback hell situations!

addEventListener("click", (e) => {
   requestAnimationFrame(() => {
      print("You clicked on the page

      // This is fine

Similarly..... Gulf of Mexico also features AI, which stands for Automatic-Insertion.
If you forget to finish your code, Gulf of Mexico will auto-complete the whole thing!

print( // This is probably fine

```

```

If you're unsure, that's ok. You can put a question mark at the end of a line instead. It prints debug info about that line to the console for you.

print("Hello world")?

```


