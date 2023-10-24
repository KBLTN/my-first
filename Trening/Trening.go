package main

import (
	"fmt"
	"strconv"
	"time"
)

func printer(docs chan string, lotok chan int) {
	for {
		paper := <-lotok
		doc := <-docs

		fmt.Println("warming...")
		time.Sleep(1 * time.Second)

		if paper == 0 {
			fmt.Println("failed to print " + doc + " because we need more paper")
			continue
		}

		fmt.Println("printing...")
		fmt.Println(doc)
		fmt.Println("turning of...")
		time.Sleep(1 * time.Second)
	}
}

func main() {
	var lotok chan int = make(chan int)
	var docs chan string = make(chan string)

	go printer(docs, lotok)

	for i := 0; i < 10; i++ {
		paper := 1
		if i == 6 {
			paper = 0
		}

		fmt.Println("put paper")
		lotok <- paper

		time.Sleep(5 * time.Second)

		fmt.Println("send document to print")
		docs <- "document #" + strconv.Itoa(i)
	}

	time.Sleep(5 * time.Second)
}
