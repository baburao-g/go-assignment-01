package main

import (
	"fmt"
	"sync"
	"time"

	"problem-01/mockstream"
)

var wg sync.WaitGroup

func producer(stream mockstream.Stream, tweets chan *mockstream.Tweet) {
	defer wg.Done()
	for {
		tweet, err := stream.Next()
		if err == mockstream.ErrEOF {
			close(tweets)
			return
		}
		tweets <- tweet
	}
}

func consumer(tweets chan *mockstream.Tweet) {
	defer wg.Done()
	for t := range tweets {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}

}

func main() {
	start := time.Now()
	defer wg.Done()

	stream := mockstream.GetMockStream()
	tweets := make(chan *mockstream.Tweet)
	wg.Add(2)
	// Producer
	go producer(stream, tweets)

	// Consumer
	go consumer(tweets)
	wg.Wait()
	fmt.Printf("Process took %s\n", time.Since(start))
}
