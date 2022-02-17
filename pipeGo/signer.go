package main

import (
	"fmt"
	// "strings"
	// "sync"
	// "time"
	"time"
)
func Single(vals ...string) <-chan string {
	out := make(chan string, 1)
	go func() {
		for _, n := range vals {
			out <- n
		}
		close(out)
	}()
	return out
}

func Multi(in <-chan string) <-chan string {
    out := make(chan string, 1)
    go func() {
        for n := range in {
            out <- n + "1-"
			time.Sleep(10000000000)
        }
        close(out)
    }()
    return out
}

func Combine(in <-chan string) <-chan string {
    out := make(chan string, 3)
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
