## godemo
### gotour
golang.org中gotour部分的练习。
#### exercise_web_crawler
1.如何确保主线程先不退出？
创建一个管道，用于父协程与子协程之间的通信。在子协程完成后通知父协程。
2.如何在多线程时保证不重复？
创建一个全局的map，用于保存所有url是否访问过的状态。
3.定义只使用一次的struct,单例模式的高级实现

```
var fetched = struct {
	m map[string]error
	sync.Mutex
}{m: make(map[string]error)}

```
4.错误记录

如下代码，未将for range出的u 值做为参数传给子协程， 导致子协程取到的url都是相同地。
```
for i, u := range urls {
			fmt.Printf("-> Crawling child %v/%v of %v : %v.\n", i, len(urls), url, u)

			go func() {
				Crawl(u, depth-1, fetcher)
				done <- "true"
			}()
		}

```

如下形式也不行，因为在子协程运行时，取到的i值也是相同的
    
```
for i, u := range urls {
			fmt.Printf("-> Crawling child %v/%v of %v : %v.\n", i, len(urls), url, u)

			go func() {
				Crawl(urls[i], depth-1, fetcher)
				done <- "true"
			}()
		}

```




## mytestdemo
experiment for go