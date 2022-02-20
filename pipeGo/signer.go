package main

import (
	// "fmt"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func ExecutePipeline(jobs ...job) {
	var wg sync.WaitGroup
	in := make(chan interface{})

	for _, job := range jobs {
		wg.Add(1)
		out := make(chan interface{})
		go ExecuteJob(&wg, job, in, out)
		in = out
	}
	wg.Wait()
}

func ExecuteJob(wg *sync.WaitGroup, curJob job, in, out chan interface{}) {
	defer wg.Done()
	defer close(out)
	curJob(in, out)
}

func SingleHash(in, out chan interface{}) {
	var wg sync.WaitGroup
	var mu = sync.Mutex{}
	for val := range in {
		d, _ := val.(int)
		s := strconv.Itoa(d)
		wg.Add(1)
		SingleHashWork(&wg, &mu, in, out, s)
	}
	wg.Wait()
}

func SingleHashWork(wg *sync.WaitGroup, mu *sync.Mutex ,in, out chan interface{}, s string) {
	defer wg.Done()

	mu.Lock()
	md5 := DataSignerMd5(s)
	mu.Unlock()
	
	crc32chan := make(chan interface{})
	go Crc32Job(s, crc32chan)
	src := <- crc32chan
	go Crc32Job(md5, crc32chan)
	mdc := <- crc32chan
	fmt.Println( src.(string) + "~" + mdc.(string))
	out <- src.(string) + "~" + mdc.(string)
}

func Crc32Job(s string, out chan interface{}) {
	out <- DataSignerCrc32(s)
}

func MultiHash(in, out chan interface{}) {
	ans := ""
	crc32chan := make(chan interface{})
	for val := range in {
		ans = ""
		s, _ := val.(string)
		for j := 0; j < 6; j++ {
			go Crc32Job(strconv.Itoa(j) + s, crc32chan)
		}
		for i:=0; i < 6; i++ {
			add := <- crc32chan
			fmt.Println(add)
			ans = ans + add.(string)
		}
		out <- ans
	}
}

func CombineResults(in, out chan interface{}) {
	res := []string{}

	for val := range in {
		res = append(res, val.(string))
	}
	sort.Strings(res)
	resStr := strings.Join(res, "_")
	out <- resStr
}
