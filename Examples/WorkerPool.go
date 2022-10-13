package Examples

import "fmt"

func WorkerPool() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go Worker(jobs, results)
	go Worker(jobs, results)
	go Worker(jobs, results)
	go Worker(jobs, results)

	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)
	for result := range results {
		fmt.Println(result)
	}

}

func Worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		//results <- Fib(n)
		results <- DynamicFib(n)
	}
}
