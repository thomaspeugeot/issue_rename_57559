package main

type Foo struct {
}

// Rename does not work for Foo
// struct [Foo] is a nice struct ...

func FuncA() {
}

// Rename works for FuncA
// FuncB is same as [FuncA].
func FuncB() {
}

func main() {

}
