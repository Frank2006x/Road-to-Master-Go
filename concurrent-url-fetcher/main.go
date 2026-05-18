package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

func fetch(
	url string,
	result map[string]string,
	mu *sync.Mutex,
) error {

	ctx, cancel := context.WithTimeout(
		context.Background(),
		2*time.Second,
	)
	defer cancel()

	req, err := http.NewRequestWithContext(
		ctx,
		"GET",
		url,
		nil,
	)

	if err != nil {
		return err
	}

	client := http.Client{}

	res, err := client.Do(req)

	if err != nil {
		return err
	}

	mu.Lock()
	result[url] = res.Status
	mu.Unlock()

	res.Body.Close()

	return nil
}


func worker (id int ,
	wg *sync.WaitGroup ,
	jobs <-chan string,
	count *int32,
	result map[string]string,
	mu *sync.Mutex,
	tick *time.Ticker) {
	defer wg.Done()
	
	
	for url:=range jobs {
		<-tick.C
		fmt.Println("Worker ",id,"fetching ",url,"time",time.Now());
		err :=fetch(url,result,mu)
		if err != nil{
			success:=false
			fmt.Println("couldnt fetch the url")
			for i:=0;i<3;i++{

				time.Sleep(
					time.Duration(i+1) *
						time.Second,
				)
				fmt.Println(
					"retry attempt",
					i+1,
					"for",
					url,
				)
				err := fetch(url, result, mu)

				if err == nil {
					success = true
					break
				}
			}
			if success {
				atomic.AddInt32(count,1)
			}

			
			continue
		}

		
		atomic.AddInt32(count,1)
	}
}

func main() {

	urls := []string{
	"https://google.com",
	"https://github.com",
	"https://golang.org",
	"https://stackoverflow.com",
	"https://reddit.com",
	"https://youtube.com",
	"https://twitter.com",
	"https://linkedin.com",
	"https://facebook.com",
	"https://amazon.com",
	"https://netflix.com",
	"https://openai.com",
	"https://wikipedia.org",
	"https://microsoft.com",
	"https://apple.com",
	"https://cloudflare.com",
	"https://docker.com",
	"https://kubernetes.io",
	"https://digitalocean.com",
	"https://aws.amazon.com",
	"https://news.ycombinator.com",
	"https://medium.com",
	"https://dev.to",
	"https://bbc.com",
}
	start:=time.Now()
	var wg sync.WaitGroup
	duration :=time.Duration(5)*time.Millisecond
	ticker:=time.NewTicker(duration)
	defer ticker.Stop()

	var success int32
	jobs := make(chan string)
	result:=make(map[string]string)
	var mu sync.Mutex
	workerCount := 3
	wg.Add(workerCount)

	for i := 1; i <= workerCount; i++ {
		go worker(i,&wg,jobs,&success,result,&mu,ticker)
	}
	for _,url := range urls{
		jobs <- url
	}
	close(jobs)
	wg.Wait();
	for k,v:=range result{
		fmt.Println(k,"fetched with status ",v)
	}
	fmt.Println(time.Since(start))
	fmt.Println("success ",atomic.LoadInt32(&success),"failed ",len(urls)-int(atomic.LoadInt32(&success)) )
}