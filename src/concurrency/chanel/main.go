package main

import "fmt"

type Job struct {
	Data int
}

func worker(jobChan chan *Job) {
	for j := range jobChan {
		if j == nil {
			fmt.Println("j is nil")
		}
		fmt.Println(j)
	}
	fmt.Println("worker finish")
}

func main() {

	jobChan := make(chan *Job)
	defer close(jobChan)
	go worker(jobChan)

	for _, j := range []*Job{{1}, {2}, {3}, nil, {4}, {5}, {6}, {7}} {
		jobChan <- j
	}

	fmt.Println("main goroutine finish")
}
