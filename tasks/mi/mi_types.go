package mi

import (
	"fmt"
	"io"
	"time"
)

type Vertex struct {
	X, Y float64
}

type MyFloat float64

type Abser interface {
	abs() float64
}

type I interface {
	m()
}

type IWithNil interface {
	mWithNil()
}

type T struct {
	S string
}

type F float64

func (f F) m() {
	fmt.Println(f)
}

type Person struct {
	Name string // a person's name
	Age  int    //	a person's age
}

type IPAddr [4]byte

type MyError struct {
	When time.Time
	What string
}

type ErrNegativeSqrt float64 // a custom type for an error

type MyReader struct {
}

type Rot13Reader struct {
	r io.Reader
}

type Image struct{}
