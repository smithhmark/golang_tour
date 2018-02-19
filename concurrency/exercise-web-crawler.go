package main

import (
	"fmt"
	"sync"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type job struct {
	url   string
	depth int
}

type jobLog struct {
	log map[string]int
	mux sync.Mutex
}

func crawl(jq chan job, doneq chan int, worker int, fetcher Fetcher, visited *jobLog, donecb func()) {
	//fmt.Printf("(%v) starting\n", worker)
	for {
		jb := <-jq
		fmt.Printf("(%v):%v\n", worker, jb)
		if jb.url == "" || len(jb.url) == 0 {
			fmt.Println("poison!!!")
			break
		}
		if jb.depth <= 0 {
			fmt.Printf("(%v) done()\n", worker)
			donecb()
			continue
		}
		visited.mux.Lock()
		visited.log[jb.url]++
		visits := visited.log[jb.url]
		visited.mux.Unlock()

		if visits > 1 {
			// skipping repeat visit
			fmt.Printf("(%v) skip: %s\n", worker, jb.url)
			continue
		}
		body, urls, err := fetcher.Fetch(jb.url)
		if err != nil {
			fmt.Printf("(%v) %s", worker, err)
			continue
		}
		fmt.Printf("(%v) found: %s %q\n", worker, jb.url, body)
		for _, u := range urls {
			jq <- job{u, jb.depth - 1}
		}
	}
	doneq <- worker
	fmt.Printf("(%v) terminating\n", worker)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice. DONE!!!
	// This implementation doesn't do either:
	visited := new(jobLog)
	visited.log = make(map[string]int)
	jq := make(chan job, 10)
	doneq := make(chan int, 2)
	jq <- job{url, depth}
	workers := 2
	done := func() {
		fmt.Println("   taking poison pills")
		for ii := 0; ii < workers; ii++ {
			jq <- job{"", -1}
		}
		return
	}
	for ii := 0; ii < workers; ii++ {
		go crawl(jq, doneq, ii, fetcher, visited, done)
	}
	complete := 0
	for complete < workers {
		select {
		case who := <-doneq:
			fmt.Printf("(MAIN) worker %d done\n", who)
			complete++
		default:
			fmt.Printf("(MAIN) dozing qlen=%v\n", len(jq))
			time.Sleep(100 * time.Millisecond)
		}
	}
	return
}

func main() {
	Crawl("https://golang.org/", 3, fetcher)
	//Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
