package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result string
type Search func(query string) Result

var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
	Radio = fakeSearch("radio")
	Audio = fakeSearch("radio")
	Uudio = fakeSearch("uudio")
	Zudio = fakeSearch("zudio")
)

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

// 1. 순차진행
// func Google(query string) (results []Result) {
// 	results = append(results, Web(query))
// 	results = append(results, Image(query))
// 	results = append(results, Video(query))
// 	results = append(results, Radio(query))
// 	results = append(results, Audio(query))
// 	results = append(results, Uudio(query))
// 	results = append(results, Zudio(query))
// 	return
// }

// 2. 병렬처리
func Google(query string) []Result {
	crawlingList := []Search{Web, Image, Video, Radio, Audio, Uudio, Zudio}
	c := make(chan Search)

	go func() {
		for i := 0; i < 7; i++ {
			c <- crawlingList[i]
		}
	}()

	var results []Result
	for i := 0; i < 7; i++ {
		fc := <-c
		results = append(results, fc(query))
	}

	return results
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Google("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}
