package mi

import (
	"fmt"
	"golang.org/x/tour/pic"
	"golang.org/x/tour/reader"
	"image"
	"io"
	"math"
	"os"
	"strings"
)

func ProcessAllTasks() {
	methods()
	methodsContinued()
	methodsPointers()
	methodsPointersAndFunctions()
	indirection()
	indirectionValues()
	methodsWithPointerReceivers()
	interfaces()
	interfacesAreSatisfiedImplicitly()
	interfacesValues()
	interfaceValuesWithNil()
	nilInterfaceValues()
	emptyInterface()
	typeAssertions()
	typeSwitches(nil)
	stringers()
	stringerTask()
	errors()
	doSqrtWithError()
	readerTask()
	infinityReader()
	rot13ReaderTask()
	images()
	//imageTask()
}

func methods() { // methods are functions with a special receiver argument
	v := Vertex{3, 4} // define a struct
	fmt.Println(v.abs())
}

func methodsContinued() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.abs())
}

func methodsPointers() {
	v := Vertex{3, 4}
	v.scale(10)
	fmt.Println(v.abs())
}

func methodsPointersAndFunctions() {
	v := Vertex{3, 4}
	scale(&v, 10) // you can declare a method on non-struct types
	fmt.Println(abs(v))
}

func indirection() {
	v := Vertex{3, 4}
	v.scale(2)        // this is the same as (&v).Scale(2)
	scaleFunc(&v, 10) // this is the same as Scale(&v, 10)

	p := &Vertex{4, 3}
	p.scale(3)      // this is the same as (&p).Scale(3)
	scaleFunc(p, 8) // this is the same as Scale(p, 8)

	fmt.Println(v, p)
}

func indirectionValues() {
	v := Vertex{3, 4}
	fmt.Println(v.abs())    //.abs() is a method with a value receiver
	fmt.Println(absFunc(v)) //absFunc is a function with a value argument

	p := &Vertex{4, 3}
	fmt.Println(p.abs())     // .abs() is a method with a pointer receiver
	fmt.Println(absFunc(*p)) // absFunc is a function with a value argument
}

func methodsWithPointerReceivers() {
	v := Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.abs())
	v.scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.abs())
}

func interfaces() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // MyFloat implements Abser
	a = &v // *Vertex implements Abser

	a = v // Vertex does not implement Abser
	fmt.Println(a.abs())
}

func interfacesAreSatisfiedImplicitly() {
	var i I = T{"hello"}
	i.m()
}

func interfacesValues() {
	var i I

	i = &T{"Hello"} // a pointer of type T
	describe(i)     // (&{Hello}, *mi.T)
	i.m()           // Hello

	i = F(math.Pi) // a float64
	describe(i)    // (3.141592653589793, mi.F)
	i.m()          // 3.141592653589793
}

func interfaceValuesWithNil() {
	var i IWithNil // an interface value that holds a nil concrete value
	var t *T       // a nil pointer of type T
	i = t          // a nil pointer of type T
	describeNil(i) // (<nil>, *mi.T)
	i.mWithNil()   // <nil>
}

func nilInterfaceValues() {
	var i I     // an interface value that holds a nil concrete value
	describe(i) // (<nil>, <nil>)
	//i.m()       // this will cause a panic
}

func emptyInterface() {
	var i interface{} // an empty interface may hold values of any type
	describeEmpty(i)  // (<nil>, <nil>)

	i = 42
	describeEmpty(i) // (42, int)

	i = "hello"
	describeEmpty(i) // (hello, string)
}

func typeAssertions() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// f = i.(float64) // this will cause a panic
	// fmt.Println(f)
}

func typeSwitches(i interface{}) {
	swithOnType(21)
	swithOnType("hello")
	swithOnType(true)
}

func stringers() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}

func stringerTask() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ipAddr := range hosts {
		fmt.Printf("%v: %v\n", name, ipAddr)
	}
}

func errors() {
	if err := runError(); err != nil { // if the function returns an error, check it
		fmt.Println(err)
	}
}

func doSqrtWithError() {
	fmt.Println(sqrtWithError(2))  // 1.4142135623730951 <nil>
	fmt.Println(sqrtWithError(-2)) // 0 cannot Sqrt negative number: -2
}

func readerTask() {
	r := strings.NewReader("Hello, Reader!") // a string reader
	b := make([]byte, 8)                     // a slice of bytes
	for {
		n, err := r.Read(b)                               // read 8 bytes into the slice
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b) // print the number of bytes read, the error, and the slice
		fmt.Printf("b[:n] = %q\n", b[:n])                 // print the slice up to n
		if err == io.EOF {                                // if the error is the end of the file
			break
		}
	}
}

func infinityReader() {
	reader.Validate(MyReader{})
}

func rot13ReaderTask() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!") // a string reader
	r := Rot13Reader{s}                             // a rot13 reader
	_, err := io.Copy(os.Stdout, &r)
	if err != nil { // if there is an error, print it
		return

	} // copy the rot13 reader to the standard output
}

func images() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100)) // create a new image
	fmt.Println(m.Bounds())                        // print the bounds of the image
	fmt.Println(m.At(0, 0).RGBA())                 // print the color of the pixel at (0, 0)
}

func imageTask() {
	m := Image{}     // create an image
	pic.ShowImage(m) // display the image
}
