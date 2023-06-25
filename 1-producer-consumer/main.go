//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer scenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
)

func producer(stream Stream, jobs chan *Tweet) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			return
		}
		jobs <- tweet

	}
}

func consumer(jobs chan *Tweet) {
	for t := range jobs {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
	// need to exit the goroutine
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	jobs := make(chan *Tweet, 20)

	// Consumer
	go consumer(jobs)

	// Producer
	go producer(stream, jobs)

	fmt.Printf("Process took %s\n", time.Since(start))
}
