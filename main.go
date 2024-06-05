package main

// The Foo type is used with [Bar].
type Foo struct {
	Name string
}

// func Bar uses [Foo] and likes [Foo.Name].
func Bar(Foo) {}
