package main

import "fmt"

const (
	info       = "Generating Fabonacci numbers: "
	verClosure = "Closure version"
	verChannel = "Channel version"
)

// FabClosure closure version
func FabClosure() func() int {
	pre, cur := 0, 1

	return func() int {
		pre, cur = cur, pre+cur
		return pre
	}
}

// FabChannel channel version
func FabChannel(ch chan int) {
	cl := cap(ch)
	pre, cur := 0, 1
	for i := 0; i < cl; i++ {
		pre, cur = cur, pre+cur
		ch <- pre
	}
	close(ch)
}

// main entrance:
// go build -o ~/gopher/bin/fabonacci  fabonacci.go
// ~/gopher/bin/fabonacci
func main() {
	fmt.Println(info)

	// Closure version
	fab := FabClosure()
	for i := 0; i < 10; i++ {
		fmt.Println(verClosure, fab())
	}

	// Channel version
	ch := make(chan int, 10)
	go FabChannel(ch)
	for i := range ch {
		fmt.Println(verChannel, i)
	}
}
