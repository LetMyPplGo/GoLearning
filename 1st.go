package main

import (
	"fmt"
	"time"
	"sync"
	"math/rand"
	"strconv"
)


func main1(){
	c := []chan string {
		make(chan string, 10),
		make(chan string, 10),
		make(chan string, 10),
	}

	fmt.Println(c)

  arr := []string{"sheep", "fish", "crawl"}

	var wg sync.WaitGroup

	for i := range arr {
		wg.Add(1)
		go func(i int){
			count(arr[i], c[i])
			wg.Done()
		}(i)

		go func(i int){
			for msg := range c[i] {
				fmt.Println(msg)
			}
		}(i)

	}

	wg.Wait()

	fmt.Println("end")
}

func count(what string, c chan string){
	timeout := 1000 + rand.Intn(1000)
	fmt.Printf("Timeout for %s is %d\n", what, timeout)

	for i:=1; i<=5; i++ {
		c <-what + strconv.Itoa(i)
		time.Sleep(time.Millisecond * time.Duration(timeout))
	}

	fmt.Printf("closing the channel %v\n", c)
	close(c)
}
