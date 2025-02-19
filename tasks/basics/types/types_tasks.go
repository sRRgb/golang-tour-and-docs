package types

import (
	"fmt"
	"golang.org/x/tour/wc"
	"math"
)
import "golang.org/x/tour/pic"

func ProcessAllTasks() {
	pointers()
	printStruct()
	structFields()
	structPointers()
	structLiterals()
	arrays()
	slices()
	slicesPointers()
	slicesLiterals()
	slicesDefaults()
	slicesLengthCapacity()
	nilSlices()
	makeSlices()
	slicesOfSlices()
	appendSlices()
	rangeSlices()
	rangeSlicesContinued()
	//sliceTask() // disabled due big response
	maps()
	mapLiterals()
	mapMutating()
	mapTask()
	functionValues()
	functionClosures()
	fibonacciFunc()
}

func pointers() {
	i, j := 42, 2701

	p := &i         // pointer to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // pointer to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

func printStruct() {
	v := Vertex{1, 2} // has type Vertex
	fmt.Println(v)
}

func structFields() {
	v := Vertex{1, 2} // has type Vertex
	v.X = 4           // change X field
	fmt.Println(v.X)
}

func structPointers() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

func structLiterals() {
	var (
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{X: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		p  = &Vertex{1, 2} // has type *Vertex
	)
	fmt.Println(v1, p, v2, v3)
}

func arrays() {
	var a [2]string         // array of 2 strings
	a[0] = "Hello"          // set first element
	a[1] = "World"          // set second element
	fmt.Println(a[0], a[1]) // print both elements
	fmt.Println(a)          // print the whole array

	primes := [6]int{2, 3, 5, 7, 11, 13} // array of 6 ints
	fmt.Println(primes)                  // [2 3 5 7 11 13]
}

func slices() {
	primes := [6]int{2, 3, 5, 7, 11, 13} // array of 6 ints

	var s []int = primes[1:4] // slice of the array from 1 to 4
	fmt.Println(s)            // [3 5 7]
}

func slicesPointers() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names) // [John Paul George Ringo]

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b) // [John Paul] [Paul George]

	b[0] = "XXX"
	fmt.Println(a, b)  // [John XXX] [XXX George]
	fmt.Println(names) // [John XXX George Ringo]
}

func slicesLiterals() {
	q := []int{2, 3, 5, 7, 11, 13} // slice of ints
	fmt.Println(q)                 // [2 3 5 7 11 13]

	r := []bool{true, false, true, true, false, true} // slice of bools
	fmt.Println(r)                                    // [true false true true false true]

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s) // [{2 true} {3 false} {5 true} {7 true} {11 false} {13 true}]
}

func slicesDefaults() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s) // [3 5 7]

	s = s[:2]
	fmt.Println(s) // [3 5]

	s = s[1:]
	fmt.Println(s) // [5]
}

func slicesLengthCapacity() {
	s := []int{2, 3, 5, 7, 11, 13}

	printSlice(s) // len=6 cap=6 [2 3 5 7 11 13]

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s) // len=0 cap=6 []

	// Extend its length.
	s = s[:4]
	printSlice(s) // len=4 cap=6 [2 3 5 7]

	// Drop its first two values.
	s = s[2:]
	printSlice(s) // len=2 cap=4 [5 7]
}

func nilSlices() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

func makeSlices() {
	a := make([]int, 5) // len(a)=5
	printSliceMake("a", a)

	b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	printSliceMake("b", b)

	c := b[:2] // len(c)=2, cap(c)=5
	printSliceMake("c", c)

	d := c[2:5] // len(d)=3, cap(d)=3
	printSliceMake("d", d)
}

func slicesOfSlices() {
	board := [][]string{ // slice of slices
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
	board[0][0] = "X" // set the first element
	board[2][2] = "O" // set the last element
	board[1][2] = "X" // set the middle right element
	board[1][0] = "O" // set the middle left element
	board[0][2] = "X" // set the first right element

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", board[i])
	}
}

func appendSlices() {
	var s []int
	printSliceAppend(s) // [] len=0 cap=0

	// append works on nil slices
	s = append(s, 0)
	printSliceAppend(s) // [0] len=1 cap=1

	// the slice grows as needed
	s = append(s, 1)
	printSliceAppend(s) // [0 1] len=2 cap=2

	// we can add more than one element at a time
	s = append(s, 2, 3, 4)
	printSliceAppend(s) // [0 1 2 3 4] len=5 cap=6
}

func rangeSlices() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func rangeSlicesContinued() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
		// << is a left shift operator that shifts the bits of a number to the left
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

func sliceTask() {
	pic.Show(Pic)
}

func Pic(dx, dy int) [][]uint8 {
	picture := make([][]uint8, dy) // slice of slices
	for i := range picture {       // iterate over the first slice
		picture[i] = make([]uint8, dx) // create a slice for each element
		for j := range picture[i] {    // iterate over the second slice
			picture[i][j] = uint8(i * j) // set the value
		}
	}

	return picture
}

func maps() {
	var m map[string]VertexMap // map of strings to VertexMap

	m = make(map[string]VertexMap) // create a map
	m["Bell Labs"] = VertexMap{
		40.68433, -74.39967,
	}

	fmt.Println(m["Bell Labs"]) // {40.68433 -74.39967}
}

func mapLiterals() {
	var m = map[string]VertexMap{ // map of strings to VertexMap
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m) // map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]
}

func mapMutating() {
	m := make(map[string]int) // map of strings to ints
	m["Answer"] = 42          // set the value
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48 // change the value
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer") // delete the value
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"] // check if the value is present
	fmt.Println("The value:", v, "Present?", ok)

	// ok is a bool that is true if the key is present in the map
}

func mapTask() {
	wc.Test(wordCount)
}

func functionValues() {
	hypot := func(x, y float64) float64 { // function value
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12)) // 13

	fmt.Println(compute(hypot))    // 5
	fmt.Println(compute(math.Pow)) // 81
}

func functionClosures() {
	pos, neg := adder(), adder() // function closures
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func fibonacciFunc() {
	f := fibonacci() // function closure
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
