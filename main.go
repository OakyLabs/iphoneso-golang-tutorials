package main

import (
	"fmt"
	"sync"
	"time"
)

type Pokemon struct {
	// json tags <-> go and json
	// on the left side (you have the golang representation)
	// on the right side you have the key of the json
	PokemonName string `json:"name"` // json tag: tells golang what name to get it from the json object
	Order       int
}

// func main() {
// 	url := "https://pokeapi.co/api/v2/pokemon/ditto"

// 	request, err := http.NewRequest("GET", url, nil)

// 	if err != nil {
// 		panic(err)
// 	}

// 	client := &http.Client{}

// 	response, err := client.Do(request)

// 	if err != nil {
// 		panic(err)
// 	}

// 	// ditto := make(map[string]any)
// 	ditto := Pokemon{}

// 	err = json.NewDecoder(response.Body).Decode(&ditto)

// 	if err != nil {
// 		panic(err)
// 	}

// 	ditto.PokemonName = "Alfonso"

// 	// {
// 	// 	name: "Alfonso",
// 	// 	order: 214
// 	// }
// 	fmt.Printf("%+v\n", ditto)

// 	bytes, err := json.Marshal(&ditto)

// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(string(bytes))
// }

// go routines

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i <= 100; i++ {
		wg.Add(1)
		go printAndSleep(i, wg)
	}

	wg.Wait()

	fmt.Println("Done with everything")

	// these are channels also known as a subcriber
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)

		// this is how you call an event in a channel
		c1 <- 420
	}()

	go func() {
		time.Sleep(time.Second)
		c2 <- 10
		c2 <- 30
	}()

outer:
	for {
		select {
		//with select you check for events called on each channel
		case channel1Result := <-c1:
			fmt.Println("Received from c1", channel1Result)
			break outer
		case channel2Result := <-c2:
			fmt.Println("Received from c2", channel2Result)
			// break
		}
	}

	close(c1)
	close(c2)

	fmt.Println("Done")
}

func printAndSleep(i int, wg *sync.WaitGroup) {
	fmt.Println(i)
	time.Sleep(time.Second * 3)
	wg.Done()
}
