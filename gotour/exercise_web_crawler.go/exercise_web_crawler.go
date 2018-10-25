package main

import (
	"errors"
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

var fetchedMap = struct {
	m    map[string]error
	lock sync.Mutex
}{m: make(map[string]error, 0)}
var loading = errors.New("loading")

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: 并行的抓取 URL。
	// TODO: 不重复抓取页面。
	// 下面并没有实现上面两种情况：
	if depth <= 0 {
		fmt.Printf("<- Done with %v, depth 0.\n", url)
		return
	}

	fetchedMap.lock.Lock()
	if _, ok := fetchedMap.m[url]; ok {
		fmt.Printf("<- Done with %v, already fetched.\n", url)
		fetchedMap.lock.Unlock()
		return
	}
	fetchedMap.m[url] = loading
	fetchedMap.lock.Unlock()

	body, urls, err := fetcher.Fetch(url)

	fetchedMap.lock.Lock()
	fetchedMap.m[url] = err
	fetchedMap.lock.Unlock()

	if err != nil {
		fmt.Printf("<- Error on %v: %v\n", url, err)
		return
	}

	fmt.Printf("Found: %s %q\n", url, body)
	done := make(chan string)

	/*如下代码，未将for range出的u 值做为参数传给子协程， 导致子协程取到的url都是相同地。
		for i, u := range urls {
			fmt.Printf("-> Crawling child %v/%v of %v : %v.\n", i, len(urls), url, u)

			go func() {
				Crawl(u, depth-1, fetcher)
				done <- "true"
			}()
		}

	如下形式也不行，因为在子协程运行时，i值已经变为最终值
		for i, u := range urls {
			fmt.Printf("-> Crawling child %v/%v of %v : %v.\n", i, len(urls), url, u)

			go func() {
				Crawl(urls[i], depth-1, fetcher)
				done <- "true"
			}()
		}

	*/
	for i, u := range urls {
		fmt.Printf("-> Crawling child %v/%v of %v : %v.\n", i, len(urls), url, u)

		go func(url string) {
			Crawl(url, depth-1, fetcher)
			done <- "true"
		}(u)
	}

	for i, u := range urls {
		fmt.Printf("<- [%v] %v/%v Waiting for child %v.\n", url, i, len(urls), u)
		<-done
	}
	fmt.Printf("<- Done with %v\n", url)

	//return
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
	fmt.Println("Fetching stats\n--------------")
	for url, err := range fetchedMap.m {
		if err != nil {
			fmt.Printf("%v failed: %v\n", url, err)
		} else {
			fmt.Printf("%v was fetched\n", url)
		}
	}
}

// fakeFetcher 是返回若干结果的 Fetcher。
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

// fetcher 是填充后的 fakeFetcher。
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
			"https://golang.org/cmd1/",
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
