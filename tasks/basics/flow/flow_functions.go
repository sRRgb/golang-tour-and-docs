package flow

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func sqrtWithIf(x float64) string { // If statement
	if x < 0 {
		return sqrtWithIf(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

func powWithShortIf(x, n, lim float64) float64 { // If statement with a short statement
	if v := math.Pow(x, n); v < lim {
		return v
	}

	return lim
}

func powWithIfElse(x, n, lim float64) float64 { // Short If statement with an else statement
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}

	return lim
}

func newtonsMethodSqrt(x float64) float64 { // Newton's method
	z := 1.0
	//z := float64(1)
	z -= (z*z - x) / (2 * z)

	return z
}

func processSwitchStatement() { // Switch statement
	fmt.Print("Go runs on ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func processSwitchOrdering() { // Switch evaluation order
	fmt.Println("Коли Субота?")

	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("After 2 days.")
	default:
		fmt.Println("Someday..")
	}
}

func processSwitchAlwaysTrue() { // Switch without a condition (always true)
	t := time.Now()

	switch { //always true
	case t.Hour() < 12:
		fmt.Println("Доброго ранку!")
	case t.Hour() < 17:
		fmt.Println("Добрий день.")
	default:
		fmt.Println("Добрий вечір.")
	}
}

func processMultipleDefers() { // more details about defer here https://go.dev/blog/defer-panic-and-recover
	fmt.Println("Counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Done")
}
