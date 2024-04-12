package main

import (
	"fmt"

	"github.com/fonsogms/iphoneso/person"
)

/*
interface withWheels {
	printAmountOfWheels(): void
}

class Car implements withWheels {
	age: number
	model: string
	color: string


	printAmountOfWheels() {
		console.log(4)
	}
}
*/

type car struct {
	age   int
	model string
	color string
}

type Alfonso struct {
	age  string
	from string
}

// func main() {
func (c car) PrintAmountOfWheels() {
	fmt.Println("I have 4 wheels")
}

func (this Alfonso) PrintAmountOfWheels() {
	fmt.Println("I am not something with wheels")
}

type withWheels interface {
	PrintAmountOfWheels()
}

func doBirthday(birthdayHaver person.HasBirthday) {
	birthdayHaver.Birthday()
}

// fn
// fun
// def
func main() {
	p := person.Person{31, false}

	doBirthday(&p)

	fmt.Println(fmt.Sprintf("The Person before birthday %+v", p))

	p.Birthday()

	fmt.Println(fmt.Sprintf("The Person after birthday %+v", p))
	p.DeathDay()
	fmt.Println(fmt.Sprintf("The Person after DeathDay %+v", p))
	// propertyEstruct := attempt.Estruct{
	// 	Property: 55,
	// }
	// // var aCar car
	// aCar := car{
	// 	age:   13,
	// 	model: "Auris",
	// 	color: "red",
	// }

	// printWheels(aCar)
	// printWheels(propertyEstruct)

	// // System.out.println
	// // console.log
	// // print
	// // fmt.Println
	// /*
	// 	Syntax for comments
	// */

	// fmt.Println("Hello, world!")

	// var variable1 string = "A variable"
	// fmt.Println("Variable 1, value: ", variable1)
	// var variable2 string
	// fmt.Println("Variable 2, value: ", len(variable2))
	// variable2 = "Another variable"
	// fmt.Println("Variable 2, value: ", variable2)
	// variable3 := "Another another variable"
	// fmt.Println("Variable 3, value: ", variable3)

	// floats := 3.14

	// result := floats + 5

	// fmt.Println(result)

	// _ = map[string]string{
	// 	"alfonso":        "garcia",
	// 	"espannnnnnnnha": "esspain",
	// }

	// _ = true || false

	// // TODO: Check undefined - pointers

	// // slices
	// slice := []int{6, 1, 2, 3}

	// // value := append(slice, 8)

	// slice = append(slice, 5)
	// _ = []int{7, 1, 2, 4}
	// // arrays

	// _ = [5]int{1, 2, 3, 4, 5}

	// // { 1 }

	// fmt.Println(sayHello("Alfonso"))

	// _, _ = twoThings()

	// _, _ = useState()
}

func sayHello(name string) string {
	return "hello " + name
}

func twoThings() (bool, string) {
	return false, "hello"
}

func useState() (string, func()) {
	return "", func() {}
}

func printWheels(somethingWithWheels withWheels) {
	somethingWithWheels.PrintAmountOfWheels()
}

// return { }

// { name: string, age: int }
// {name; "", age: 0}
