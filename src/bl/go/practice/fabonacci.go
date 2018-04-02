package main

import "fmt"

// FabClosure closure version
func FabClosure() func() int {
	pre, cur := 0, 1

	return func() int {
		pre, cur = cur, pre+cur
		return pre
	}
}

// FabChannel channel version
func FabChannel() {

}

// main entrance:
// go build -o ~/gopher/bin/fabonacci  fabonacci.go
// ~/gopher/bin/fabonacci
func main() {
	fmt.Println("Generating Fabonacci numbers: ")

	/*
		fab := FabClosure()
		for i := 0; i < 10; i++ {
			fmt.Println(fab())
		}
	*/

}
