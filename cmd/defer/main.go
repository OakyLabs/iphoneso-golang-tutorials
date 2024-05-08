package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("What is defer? or shall i say better finally?")

	doStuff()
}

// per function stack
// every defer pushes to the stack
// for length-- stack.pop()()
func doStuff() {
	fmt.Println("Make an API call")
	time.Sleep(time.Second * 3)
	fmt.Println("Make a DB call only once. After first db write, we close db forever")
	defer fmt.Println("CloseDB()")
	// defer func() {
	// 	fmt.Println("CloseDB()")
	// }()
	defer fmt.Println("Do some hard work")

	time.Sleep(time.Millisecond * 300)

	fmt.Println("Done with hard work")
}

// try {writeToDB()} catch(err) {dealWithError(err)} finally {closeDB()}
//
