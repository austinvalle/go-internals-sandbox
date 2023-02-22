package main

import "fmt"

func main() {
	exStruct := ExampleStruct{
		strField1: "ayo",
		intField:  12,
		strField2: "waddup",
	}

	fmt.Printf("%v\n", exStruct)
}

type ExampleStruct struct {
	strField1 string `av:"field1"`
	intField  int    `av:"field2"`
	strField2 string `av:"field3"`
}
