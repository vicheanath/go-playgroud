package main

import (
	"fmt"
)

// T /**
type T struct {
	name  string // name of the object
	value int    // its value
}

// run gofmt -w helloworld.go to format the code

func main() {
	fmt.Println("Hello, BF")
}

// Create object owner with getter and setter
type Owner struct {
	name  string
	value int
}

//go func() { for { dst <- <-src } }()

func f() {

}

func g() {

}

//if i < f(){ //
//	g()
//}

//if i < f()  // wrong!
//{           // wrong!
//g()
//}

//Control structures

// If statement in go it simple

//if x > 0{
//	return y
//}
