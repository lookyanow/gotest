package main

import "fmt"
import "time"

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {

	jobs := make(chan int, 500)
	results := make(chan int, 500)

	for w := 1; w <= 5; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 500; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 500; a++ {
		mess:=<-results
		fmt.Println(mess)
	}
}
