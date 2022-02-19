package main

import (
	"fmt"
	"runtime"
	"time"
	// "strings"
	"sync"
)

func	SingleWorker(val string, out chan string, wg sync.WaitGroup) {
		defer wg.Done()
		runtime.Gosched()
		val = val + "Singled"
		time.Sleep(1000000000)
		out <- val
		return
}


func Single(vals ...string) (<-chan string) {
    out := make(chan string, 3)
	var wg sync.WaitGroup 
    go func() {
        for _, n := range vals {
			fmt.Println("Single got", n)
			wg.Add(1)
			go SingleWorker(n, out, wg)
		}
        // close(out)
		wg.Wait()
    }()
	return out
}

func Multi(in <-chan string) (<-chan string) {
    out := make(chan string, 1)
    go func() {
        for n := range in {
			fmt.Println("Milti got", n)
            out <- n + "1-"
        }
        close(out)
    }()
    return out
}

func Combine(in <-chan string) (<-chan string) {
    out := make(chan string, 1)
    go func() {
        for n := range in {
            out <- n + "2"
        }
        close(out)
    }()
    return out
}

func main() {

	vals := []string{"Vasya", "Kolya", "Petya"}
	s := Single(vals...)
	m := Multi(s)
	out := Combine(m)

	for val := range out {
		fmt.Println(val)
	}
}
