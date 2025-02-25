package mi

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"time"
)

func (v Vertex) abs() float64 { // a method is a function with a special receiver argument
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f MyFloat) abs() float64 { // you can define a method on any type you define in your package, not just structs
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

func (v *Vertex) scale(f float64) { // methods can be defined for either pointer or value receiver types
	v.X = v.X * f
	v.Y = v.Y * f
}

func abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func scale(v *Vertex, f float64) { // you can declare a method on non-struct types
	v.X = v.X * f
	v.Y = v.Y * f
}

func scaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func absFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (t T) m() {
	fmt.Println(t.S)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func describeNil(i IWithNil) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func (t *T) mWithNil() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func describeEmpty(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i) // an empty interface may hold values of any type
}

func swithOnType(i interface{}) { // type switches are a form of switch statement that compares the type of the value in an interface
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2) // in this case, the variable v will be of type int
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v)) // in this case, the variable v will be of type string
	default:
		fmt.Printf("I don't know about type %T!\n", v) // in this case, the variable v will be of the same interface{} type
	}
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v років)", p.Name, p.Age) // Stringer is a type that can describe itself as a string
}

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3]) // Stringer is a type that can describe itself as a string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What) // error is the conventional interface for representing an error
}

func runError() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func sqrtWithError(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	return math.Sqrt(x), nil
}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e)) // error is the conventional interface for representing an error
}

func (r MyReader) Read(b []byte) (int, error) { // Reader is the interface that wraps the basic Read method
	for i := range b {
		b[i] = 'A'
	}

	return len(b), nil
}

func (rot *Rot13Reader) Read(p []byte) (n int, err error) { // Reader is the interface that wraps the basic Read method
	n, err = rot.r.Read(p)
	for i := 0; i < n; i++ { // rot13Reader is a type that implements the io.Reader interface
		if (p[i] >= 'A' && p[i] <= 'Z') || (p[i] >= 'a' && p[i] <= 'z') {
			if (p[i] >= 'A' && p[i] <= 'M') || (p[i] >= 'a' && p[i] <= 'm') {
				p[i] += 13
			} else {
				p[i] -= 13
			}
		}
	}
	return
}

func (m Image) ColorModel() color.Model { // Image is the interface that wraps the basic ColorModel method
	return color.RGBAModel // the ColorModel method returns the Image's color model
}

func (m Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 100) // the Bounds method returns the domain for which At can return non-zero color
}

func (m Image) At(x, y int) color.Color {
	v := uint8(x * y % 255)           // the At method returns the color of the pixel at (x, y)
	return color.RGBA{v, v, 255, 255} //(0, 0) is black, (255, 255) is white
}
