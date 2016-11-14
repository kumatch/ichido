package main

import (
	"fmt"

	"github.com/kumatch/ichido"
)

func provider() []*ichido.ContentHolder {
	return []*ichido.ContentHolder{
		&ichido.ContentHolder{Value: "foo"},
		&ichido.ContentHolder{Value: "bar"},
		&ichido.ContentHolder{Value: "baz"},
	}
}

func checker(c *ichido.ContentHolder) bool {
	if c.Value == "bar" {
		return false
	}

	return true
}

func invoker(c *ichido.ContentHolder) {
	fmt.Printf("c = %v", c)
}

func marker(c *ichido.ContentHolder) {
	// log content value.
}

func main() {
	i := ichido.Ichido{}
	i.RegisterContentsProvider(provider)
	i.RegisterNewlyChecker(checker)
	i.RegisterInvoker(invoker)
	i.RegisterInvokedMarker(marker)
	i.Run()
}
