package main

import (
	"fmt"
	"sync"
	"time"
	// "time"
)

func SingleHash(in, out chan string, wg *sync.WaitGroup) {
	defer wg.Done()	
	val := ""
	for {
		val = <-in
		time.Sleep(1000000000)
		fmt.Println("SingleHash got: ", val)
		out <- val
	}
}

func MultiHash(in, out chan string, wg *sync.WaitGroup) {
	defer wg.Done()	
	val := ""
	for {
		val = <-in
		fmt.Println("MultiHash got: ", val)
		out <- val
	}
}

func CombineResults(in chan string, wg *sync.WaitGroup) {
	defer wg.Done()	
	val := ""
	for {
		val = <-in
		fmt.Println("CombineResults got: ", val)
	}
}

func main() {

	wg := &sync.WaitGroup{}
	vals := []string{"Vasya", "Kolya", "Petya"}
	Single := make(chan string, 1)
	Multi := make(chan string, 1)
	Combine := make(chan string, 1)

	wg.Add(1)
	go SingleHash(Single, Multi, wg)
	wg.Add(1)
	go MultiHash(Multi, Combine, wg)
	wg.Add(1)
	go CombineResults(Combine, wg)

	for _, val := range vals {
		Single <- val
	}
	// time.Sleep(1000000000)

	wg.Wait()
}
