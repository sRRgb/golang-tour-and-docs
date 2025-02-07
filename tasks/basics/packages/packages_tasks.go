package packages

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

const Pi = 3.14         // float64 constant
const World = "World"   // string constant
const Truth = true      // boolean constant
const Big = 1 << 100    // numeric constant
const Small = Big >> 99 // numeric constant

var c, python, java bool // Variables on the package level

func ProcessAllTasks() {
	helloWorld()
	welcome()
	packages()
	imports()
	exportedNames()
	multipleParams()
	multipleParamsContinued()
	multipleReturn()
	nakedReturn()
	variables()
	initializedVariables()
	shortVariableDeclarations()
	baseTypes()
	zeroValues()
	typeConversions()
	typeInference()
	constants()
	numericConstants()
}

func helloWorld() { // hello world
	fmt.Println("Hello, World!")
}

func welcome() { // welcome to the playground
	fmt.Println("Welcome to Playground")
	fmt.Println("Time -", time.Now())
}

func packages() { // Packages
	fmt.Println("My favorite number is:", rand.Intn(10))
}

func imports() { // Imports
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}

func exportedNames() { // Exported names
	//fmt.Println(math.pi) // Error: cannot refer to unexported name math.pi
	fmt.Println(math.Pi)
}

func multipleParams() { // Function with multiple parameters of the same type
	fmt.Println(add(42, 13))
}

func multipleParamsContinued() { // Function with multiple parameters of the same type (continued)
	fmt.Println(addContinued(42, 13))
}

func multipleReturn() { // Multiple return results
	a, b := swap("Hello", "World")

	fmt.Println(a, b)
}

func nakedReturn() { // Named return values x and y (aka naked return)
	fmt.Println(split(17))
}

func variables() { // Variables on the function level and package level
	var i int

	fmt.Println(i, c, python, java)
}

func initializedVariables() { // Variables with initializers
	var c, python, java = true, false, "ні!" // Variables on the function level
	var i, j int = 1, 2                      // Variables on the function level

	fmt.Println(i, j, c, python, java)
}

func shortVariableDeclarations() { // Short variable declarations
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "ні!"

	fmt.Println(i, j, k, c, python, java)
}

func baseTypes() { // Basic types
	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

func zeroValues() { // Zero values
	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

func typeConversions() { // Type conversions
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	//var f float64 = math.Sqrt((x*x + y*y))	// Error: cannot use x * x + y * y (type int) as type float64 in argument to math.Sqrt
	var z uint = uint(f)
	//var z uint = (f)	// Error: cannot use f (type float64) as type uint in assignment

	fmt.Println(x, y, z)
}

func typeInference() { // Type inference (omit the type from the declaration)
	v := 42 // change me!
	fmt.Printf("v is of type %T\n", v)

	i := 42           // int
	f := 3.142        // float64
	g := 0.867 + 0.5i // complex128

	fmt.Printf("variables are of type %T %T %T", i, f, g)
}

func constants() { // Constants
	fmt.Println("Hello", World)          //	Hello World
	fmt.Println("Good", Pi, "Morning")   //	 Good 3.14 Morning
	fmt.Println("Are you human?", Truth) // Are you human? true
}

func numericConstants() { // Numeric constants
	fmt.Println(needInt(Small))   // 21
	fmt.Println(needFloat(Small)) // 0.2
	fmt.Println(needFloat(Big))   // 1.2676506002282295e+29
	//fmt.Println(needInt(Big))     // Error: constant 1267650600228229401496703205376 overflows int
}
