# Go Notes

WHY GO?
---
- General Purpose Language - great backend
- Cheap and easy concurrency and threading
- Cloud and multicore infrastructure
- Scaled, Distributed and Dynamic Uses

GO USES
---
- Performant Apps 
- Scaled and distributed systems
- Docker Kubernetes, HashiCorp Vault

GO CHACTERISTICS
---
- Simple and readable syntax of a dynamically typed language like Python
- Efficiency and safety of a lower-level, statically typed language like C++
- Server-side or backend language
- Fast and resource-cheap
- Compiles into single binary (machine code)
- Memory freed automatically

GO PACKAGE  
---
A collection of code files.
Can be imported using `import("example")`  
**Shortcut :** 
```
import ("fmt")
var pl = fmt.Println
```

GO MAIN FUNCTION
---
From where code is executed -
```
func main(){
  pl("Hello World")
}
```

GO VARIABLES
---
`If Capital first letter - Variable is Exported`  
>**NOTE:**  Variables if declared must be used (same for packages)

*For unused variables - use underscore ( _ )*

**var and const** are available for declaring variables - var allows mutability while const does not.

**Types of declarations:**
```
var var_name string/int
var var_name = value
var v1, v2 = val1, val2
```
**Syntactic sugar:**
`var_name := value`  

  
GO DATATYPES
---
- int, float64, bool, string, rune (unicode)  
- 0,   0.0,     false, ""   , ''  -> Default Values

Go provides a rich set of built-in types, including:  
- **Basic types:** such as int, float64, string, bool, and complex64. [All - int64, uint8, uint16, uint32, uint64, int, uint, rune, byte, uintptr, float32, float64, complex64, complex128, bool, string].
- **Composite types:** such as array, slice, map, struct, and interface.
- **Pointer types:** denoted by * followed by the type, for example, *int represents a pointer to an integer.
- **Function types:** represent functions with a specific signature, allowing functions to be assigned to variables or passed as arguments to other functions.
- **Interface types:** specify a set of methods that a concrete type must implement to satisfy the interface.
- **Channel types:** used for communication and synchronization between goroutines, the lightweight concurrent units of execution in Go.
  
Get the type of variable:
```
import "reflect"
reflect.TypeOf(val/var)
```
**Casting**
```
import "strconv"
v1 := 1.5
v2 := int(v2)

v3 := "500000000"
v4 := strconv.Atoi(v3)

v5 := 500000000
v6 := strconv.Itoa(v5)

v7 := "9.12"
v8 := strconv.ParseFloat(v7, 64)

// Use Sprintf to make a variable string instead of Printf
v9 := fmt.Sprintf("%f", v8)
```

**type Keyword**
1. **Type Declarations:** The type keyword is used to create named types. It provides a way to define new types based on existing types.  
By creating a named type, one can assign more descriptive names to a particular type or create aliases for existing types.  
For example:
```
type MyInt int
var num MyInt = 42
```
2. **Type Definitions:** The type keyword is used to define new types based on existing types, creating a completely new type with its own set of fields and methods. This is commonly used to define struct types.
For example:
```
type Person struct {
    Name string
    Age  int
}

func main() {
    p := Person{
        Name: "John",
        Age:  30,
    }
    fmt.Println(p.Name)
}
```


GO PRINT
---
fmt.Printf - formatted data - template string contains text to be formatted  
`fmt.Printf("Text %v text", variable)`
 
**OR**

fmt.Print or fmt.Println  
- %.2f = float data type upto two decimal points
- %d - Integer
- %c - character
- %t - Boolean
- %s - String
- %o - Base8
- %x - Base16
- %v - any variable
- %T - type,
- %9f - +9 spaces before float value  

Command line arguments = os.Args



GO USER INPUT
---
`fmt.Scan(&var_name) //use pointer &var_name`

**OR**
```
import ("bufio" "os" "log")
reader := bufio.NewReader(os.Stdin)
name, err := reader.ReadString('\n') //read till newline
if err == nil{
	pl(name)
}
else{
	log.Fatal(err)
}
```


GO STRING
---
Array of bytes

`import "strings"`

**Replace**  
```
s1 := "A GOAT"
replacer := strings.NewReplacer("A","The")
s2 := replacer.Replace(s1)

s3 := strings.Replace(s2, "A", "1", -1) // -1 -> all instances; enter a number for the number of times replace action should be performed
```
**Length**  
`len(s1)`

**Contains**  
`var := strings.Contains(s2, "value")`

Similarly, HasPrefix and HasSuffix

**First Index**  
`var := strings.Index(s1, "O")`

**Trim**  
`s4 = strings.TrimSpace(s1)`

**Split**  
`s5, s6 := strings.Split(s1, " ")`

**Convert**  
```
strings.ToLower(string)
strings.ToUpper(string)
```
 

GO RUNES
---
Characters in GO are called runes
```
import "unicode/utf8"
str := "asdf"
num := utf8.RuneCountInString(str)
```
%#U - Unicode  
%c - Rune/Character  

GO LOOP
---
**Only For and for-each loop**

for condition {  
	// code  
}

for index, element := range slice_name -- range iterates over diff data structures  
if index is not used then use _  
```
for x:=1; x<=5; x++{
	pl(x)
}

x := 0
for x<5{
	pl(x)
	x++
}

for true{
	// code
}
```

GO IF...ELSE
---

if condition {  
  
}  
  
else if condition{  
  
}   

else{  
  
}  


> NOT of a value - !value

**break** - terminates execution of the current loop  
**continue** - skip the remaining portion of the loop, and return to the top of the loop and continue a new iteration



GO SWITCH
---
```
switch var {
 case "value1":
	// code
 case "value2", "value3":
	// code
 default:
	// code
}
```


GO TIME
---
```
import "time"

now := time.Now()
var year := now.Year()
var month := now.Month()
var day := now.Day()
var hour := now.Hour()
var minute := now.Minute()
var second := now.Second()
```


GO OPERATORS
---
+, -, *, /, %

*Shorthand  
var := value  
var += 1  
var++*


GO RAND
---
The seed() method is used to initialize the random number generator. The random number generator needs a number to start with (a seed value), to be able to generate a random number. By default the random number generator uses the current system time.  

The rand.Seed() function is used to set a seed value to generate random numbers. 
If the Seed value is the same then rand.Intn() function will generate the same series of random numbers. If we change the seed value then it will generate different series of random numbers.
```
seedVal := time.Now().Unix()
rand.Seed(seedVal)
randNum := rand.Intn(50)+1
```

GO MATH
---
`import "math"`  

Abs, Pow, Sqrt, Cbrt, Ceil, Floor, Round, Log2, Log10, Log, Exp, Max, Min, Pi, Many Trig Functions


GO FUNCTIONS
---
```
func func_name(inparam type) outparamtype{  
	// code  
	return var  
}
```
>Go functions can return many variable values

**For returning more than one values:**  
`func func_name(inparam type) (outparamtype1, outparamtype2){}`  
`func func_name(inparam type) (outparam1 type, outparam2 type){}`

**Call the function:**  
`fmt.Println(func_name(value))`

**Many input:**   
`func func_name(nums ...int) int{} //treat as array`

**With Pointers**
```
func f1(ptr *int){
*ptr = new_val
}

func main(){
var := val
f1(&var)
}
```

**Export Functions:**  
Capital Starting Letter of Function  
Global Variable - Capital Starting Letter
  
GO DATA STRUCTURES
---
**Arrays:** Arrays in Go are fixed-size collections of elements of the same type. Once an array is defined, its size cannot be changed. Arrays are useful when one knows the exact number of elements to be stored in advance.  

**Slices:** Slices are similar to arrays but with a dynamic size. They are references to underlying arrays and allow for flexible resizing. Slices are widely used in Go due to their versatility and ease of use.  

**Maps:** Maps in Go are unordered collections of key-value pairs. They provide efficient lookup and retrieval of values based on unique keys. Maps are commonly used to store and retrieve data based on a specific identifier or key.  

**Structs:** Structs are user-defined types that allows one to group together different fields of different types into a single entity. They are similar to classes in object-oriented programming and are commonly used to represent complex data structures.  

**make keyword:**  
The make keyword in Go is a built-in function that is used to allocate and initialize certain types of data structures. It is primarily used for creating slices, maps, and channels.  

When using make with slices, one specifies the slice type and the desired length and capacity. The function then allocates a new underlying array for the slice and initializes it with zero values.

`slice := make([]int, 5, 10)`  
In this example, make([]int, 5, 10) creates a new slice of type []int with a length of 5 and a capacity of 10. The underlying array has a length of 10, but the slice is currently limited to the first 5 elements.  

GO ARRAY
---
Fixed Sized - Same Datatype
```
var var_name [50]string
// or
var var_name = [50]string
// or
var var_name = [50]string{"dcgh", "trfdsgf", "trfdc"}

// Assign
var_name[i] = "xfgc" 

// String to Rune
str := "abcd"
rArr := []rune(str)

// Byte array to String
byteArr := []byte{'a', 'b', 'c'}
str := string(byteArr[:])
```

GO SLICE  
---
Flexible - Variable length
```
var var_name = []string
// OR
var_name := make([]string, 6)

var_name = append(var_name, value)
len(var_name)

var[0:2] // zero index and first index
var[:2] // till second index
var[3:] // from fourth index

var1 = var[0:2]
// Changes in any variable will reflect on both
```

GO MAP
---
`var var_name = make(map[string]string)` - map key->string and value->string  
**OR**  
`var := map[int]string{1: "val1", 2: "val2" }`

List of map of string key and value  
`make([]map[string]string, 0)`


GO STRUCT
---
```
type name_of_struct struct {
	var1 type
	var2 type
}
```
*type refers to user type*
```
var var_name = name_of_struct {
	var1 = value1,
	var2 = value2,
};

var3 := var_name.var1
```

GO GENERICS
---
Generics in Go allows one to define functions or types that can operate on a variety of different types. To use generics in Go, one can define a generic type or function using type parameters. These type parameters act as placeholders for actual types that will be provided when using the generic code. The type parameters are specified within angle brackets < > after the function or type declaration.  

```
func Swap[T any](a, b T) (T, T) {
    return b, a
}
```

```
type MyType interface {
	int | float64
}

func newFunc[T MyType](x T, y T) T {
	return x+y
}

// Both int and float can be used
```

GO INTERFACE
---
Interfaces are a powerful feature that allows one to define a set of method signatures. An interface type represents a contract that specifies the behavior a type must implement in order to satisfy the interface.

**Method Signatures:** An interface is defined by specifying a set of method signatures. These method signatures describe the behavior that types implementing the interface should have. The actual implementation of the methods is provided by the types that satisfy the interface.

**Polymorphism:** Interfaces in Go enable polymorphism, which means that multiple types can satisfy the same interface. This allows one to write code that works with different types as long as they satisfy the required interface. This promotes code reusability and flexibility.

**Implicit Implementation:** In Go, there is no explicit declaration that a type implements an interface. If a type has the required methods with matching signatures, it is said to implicitly satisfy the interface. This makes it easy to use existing types with interfaces without modifying their source code.  

```
type A interface {
	F1()
	F2()
}

type B string

func (b B) F1() {
	fmt.Println("F1 called")
}

func (b B) F2() {
	fmt.Println("F2 called")
}

func main() {
	var var1 A
	var1 = B("value")
	var1.F1()

	var var2 A = var1.(B) // type conversion
	var2.F2()
}

```
```
// File Interface
type File interface {
    Read(p []byte) (n int, err error)
    Write(p []byte) (n int, err error)
    Close() error
}

// Sorter Interface
type Sorter interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}
```

```
type Shape interface {
    Area() float64
    Perimeter() float64
}

type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

func main() {
    rect := Rectangle{Width: 5, Height: 3}
    var shape Shape
    
    shape = rect
    
    fmt.Println("Area:", shape.Area())
    fmt.Println("Perimeter:", shape.Perimeter())
}
```

GO I/O
---
```
f := os.Create("file.txt")
defer f.Close() 		// wait till function scope ends
_ := f.WriteString(var)		// write

f1 = os.Open("file.txt")
scann := bufio.NewScanner(f)	// scanner
for scann.Scan(){		// each line
	fmt.Printf(scann.Text())// content
}

os.Stat("file.txt")
f := os.OpenFile("file.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // modes and permission
```

GO DEFER, PANIC AND RECOVER
---
A defer statement pushes a function call onto a list. The list of saved calls is executed after the surrounding function returns.  
Defer is commonly used to simplify functions that perform various clean-up actions.

Panic is a built-in function that stops the ordinary flow of control and begins panicking.  
When the function F calls panic, execution of F stops, any deferred functions in F are executed normally, and then F returns to its caller. To the caller, F then behaves like a call to panic. The process continues up the stack until all functions in the current goroutine have returned, at which point the program crashes. Panics can be initiated by invoking panic directly. They can also be caused by runtime errors, such as out-of-bounds array accesses.  

Recover is a built-in function that regains control of a panicking goroutine. Recover is only useful inside deferred functions. During normal execution, a call to recover will return nil and have no other effect. If the current goroutine is panicking, a call to recover will capture the value given to panic and resume normal execution.
  
When a panic occurs, the program enters a panic state. In this state, the execution of the surrounding code is halted, and the program starts to unwind the stack, executing any deferred functions (using the defer keyword) in reverse order.
```
func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()

    fmt.Println("Before panic")
    panic("Something went wrong!")
    fmt.Println("After panic") // This line is never reached
}
```

GO CONCURRENCY
---
**GOROUTINES**  
Goroutine is a lightweight thread of execution managed by the Go runtime. Goroutines allow functions or methods to be executed concurrently without the need for explicit thread management. To create a goroutine, one uses the go keyword followed by a function call.  
For example, `go f(x, y, z)` will execute the function f concurrently in a goroutine.  
Goroutines are multiplexed onto multiple OS threads by the Go runtime, meaning that a program with multiple goroutines can run concurrently on multiple CPU cores, taking advantage of parallelism when available. Goroutines are very lightweight, with their initial stack size being only a few kilobytes, allowing one to create thousands or even millions of goroutines without much overhead.


**GO CONCURRENCY SYNC**  
> FROM sync package

Use **WaitGroup** - waits for the goroutine to finish if main thread is stopped.  
A **WaitGroup** in Go is a synchronization primitive that allows you to wait for a collection of goroutines to finish their execution before proceeding.  
The WaitGroup has three main methods:  
**Add(delta int):** This method is used to add or subtract from the WaitGroup counter. One typically calls Add(1) before starting a goroutine to indicate that a goroutine has been launched. The delta value can be positive or negative, indicating the number of goroutines to add or subtract from the WaitGroup counter.  
**Done():** This method is used to decrement the WaitGroup counter. It is called by each goroutine once it has finished its execution.  
**Wait():** This method blocks until the WaitGroup counter becomes zero. It is typically called by the main goroutine to wait for all the other goroutines to finish before proceeding.  

```
var wg = sync.WaitGroup{}
wg.Add(1)
wg.Wait()
wg.Done()
```
```
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	// Perform some work...
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("All workers have completed")
}
```
Go Channnels


Go Mutex Lock

type Var struct{
	var1 datatype
	lock sync.Mutex
}

func (v *Var) GetVar1 datatype{
	v.lock.Lock()
	defer v.lock.Unlock()
	return v.var1
}


Go Closures
Functions that don't have to be associated with an identifier
Go closure is a nested function that allows us to access variables 
of the outer function even after the outer function is closed.

Nested Functions
Returning a function

func f1(f2 func(int, int) int, x, y int){
	fmt.Println(f2(x, y))
}

func f2(x, y int) int {
	return x+y
}

f1(f2, 7, 8)


Let’s create an anonymous version of a function:
func (message string) {
    fmt.Println(message)
}("Hello!")

Directly implement after defining the function.
