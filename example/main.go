package main

import (
	"fmt"

	"github.com/kumatch/ichido"
)

type content struct {
	key   int
	value string
}

func provider() []interface{} {
	return []interface{}{
		&content{value: "foo", key: 1},
		&content{value: "bar", key: 2},
		&content{value: "baz", key: 3},
	}
}

func checker(c interface{}) bool {
	co, ok := c.(*content)
	if !ok {
		return false
	}
	if co.key == 2 {
		return false
	}

	return true
}

func invoker(c interface{}) {
	fmt.Printf("content value = %v\n", c)
}

func marker(c interface{}) {
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
