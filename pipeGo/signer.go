package main

import (
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
		go SingleHashWork(&wg, &mu, in, out, s)
	}
	wg.Wait()
}

func SingleHashWork(wg *sync.WaitGroup, mu *sync.Mutex, in, out chan interface{}, s string) {
	defer wg.Done()

	crc32chan := make(chan interface{})
	md5chan := make(chan interface{})

	mu.Lock()
	md5 := DataSignerMd5(s)
	mu.Unlock()

	go Crc32Job(s, crc32chan)
	go Crc32Job(md5, md5chan)
	src := <-crc32chan
	mdc := <-md5chan
	out <- src.(string) + "~" + mdc.(string)
}

func Crc32Job(s string, out chan interface{}) {
	out <- DataSignerCrc32(s)
}

func MultiHash(in, out chan interface{}) {
	var wg sync.WaitGroup
	for val := range in {
		s, _ := val.(string)
		wg.Add(1)
		go MultiHashLaunch(&wg, s, out)
	}
	wg.Wait()
}

func MultiHashLaunch(wg *sync.WaitGroup, s string, out chan interface{}) {
	defer wg.Done()

	var wgMod sync.WaitGroup
	arr := make([]string, 6)

	crc32chan := make(chan interface{})
	for j := 0; j < 6; j++ {
		wgMod.Add(1)
		go Crc32JobMod(&wgMod, &arr[j], strconv.Itoa(j)+s, crc32chan)
	}
	wgMod.Wait()
	ans := strings.Join(arr, "")
	out <- ans
}

func Crc32JobMod(wg *sync.WaitGroup, arr *string, s string, out chan interface{}) {
	defer wg.Done()
	*arr = DataSignerCrc32(s)
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
