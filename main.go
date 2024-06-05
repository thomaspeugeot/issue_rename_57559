package main

// The Foo type is used with [Bar].
type Foo struct{}

// func Bar uses [Foo].
func Bar(Foo) {}

// some comment about [Foo]
