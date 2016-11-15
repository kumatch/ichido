package main

import (
	"fmt"

	"github.com/kumatch/ichido"
)

type content struct {
	value string
	key   int
}

func (c *content) GetValue() interface{} {
	return c.value
}

func (c *content) GetKey() interface{} {
	return c.key
}

func provider() []ichido.ContentHolder {
	return []ichido.ContentHolder{
		&content{value: "foo", key: 1},
		&content{value: "bar", key: 2},
		&content{value: "baz", key: 3},
	}
}

func checker(c ichido.ContentHolder) bool {
	key, ok := c.GetKey().(int)

	if !ok {
		return false
	}
	if key == 2 {
		return false
	}

	return true
}

func invoker(c ichido.ContentHolder) {
	fmt.Printf("content value = %v\n", c.GetValue())
}

func marker(c ichido.ContentHolder) {
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
