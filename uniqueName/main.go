package main

import (
	"fmt"
	"reflect"
)

func uniqueNames(a, b []string) []reflect.Value {

	names := make(map[string]string)

	for _, name := range append(a, b...) {
		names[name] = name
	}

	return reflect.ValueOf(names).MapKeys()
}

func main() {
	// should print Ava, Emma, Olivia, Sophia

	fmt.Println(uniqueNames(
		[]string{"Ava", "Emma", "Olivia"},
		[]string{"Olivia", "Sophia", "Emma"},
	))
}
