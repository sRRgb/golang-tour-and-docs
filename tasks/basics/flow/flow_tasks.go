package flow

import (
	"fmt"
)

func ProcessFlowTasks() {
	forLoop()
	forWithoutParams()
	forAsWhile()
	foreverLoop()
	ifStatement()
	shortIfStatement()
	ifElseStatement()
	loopsAndFunctions()
	switchStatement()
	switchOrdering()
	switchAlwaysTrue()
	deferStatement()
	multipleDeferStatements()
}

func forLoop() { // For loop
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)
}

func forWithoutParams() { // For loop without init and post statements
	sum := 1

	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)
}

func forAsWhile() { // For loop as while
	sum := 1

	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)
}

func foreverLoop() { // Forever loop
	for {
		break // or return to exit the loop
	}
}

func ifStatement() { // If statement
	fmt.Println(sqrtWithIf(2), sqrtWithIf(-4))
}

func shortIfStatement() { // If statement with a short statement
	fmt.Println(powWithShortIf(3, 2, 10), powWithShortIf(3, 3, 20))
}

func ifElseStatement() { // Short If statement with an else statement
	fmt.Println(powWithIfElse(3, 2, 10), powWithIfElse(3, 3, 20))
}

func loopsAndFunctions() { // Loops and functions
	fmt.Println(newtonsMethodSqrt(2))
}

func switchStatement() { // Switch statement
	processSwitchStatement()
}

func switchOrdering() { // Switch evaluation order
	processSwitchOrdering()
}

func switchAlwaysTrue() { // Switch with no condition
	processSwitchAlwaysTrue()
}

func deferStatement() { // Defer statement
	defer fmt.Println("World") // Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in

	fmt.Println("Hello")
}

func multipleDeferStatements() {
	processMultipleDefers()
}
