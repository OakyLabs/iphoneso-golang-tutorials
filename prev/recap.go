package main

import "fmt"

// ts
// class has methods + data
// interfaces describe methods + data
// types describe methods + data

// go
// struct -> only data
// interfaces -> only methods

type Alfonso struct {
	Name string
}

func (a Alfonso) SayHi() {
	fmt.Println("Hey there, my name is " + a.Name)
}

type Greeter interface {
	SayHi()
}

func main() {
	fmt.Println("Hello world")

	alfonso := Alfonso{Name: "Alfonso"}

	greet(alfonso)

	dictionaryInPython := map[string]string{}

	dictionaryInPython["person"] = "someone"

	fmt.Println(dictionaryInPython)

	// if dictionaryInPython.person
	// if 'person' in dictionaryInPython

	_, isThereInDict := dictionaryInPython["person_1"]

	if !isThereInDict {
		fmt.Println("Did not find it")
	}
}

func greet(greeter Greeter) {
	greeter.SayHi()
}
