package main

import "fmt"
import "time"

// These workers will receive work on the `jobs` channel and send the corresponding
// results on `results`. We'll sleep a second per job to simulate an expensive task.
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("worker", id, "started  job", j)
        time.Sleep(time.Second)
        fmt.Println("worker", id, "finished job", j)
        results <- j * 1
    }
}

func main() {

	// Two channels - to send them work and to collect their results. 
    jobs := make(chan int, 700)
    results := make(chan int, 700)

    // This starts up 2 workers, initially blocked because there are no jobs yet.
    for w := 1; w <= 2; w++ {
        go worker(w, jobs, results)
    }

    // Here we send 6 `jobs` and then `close` that channel to indicate that's all the work we have.
    for j := 1; j <= 6; j++ {
        jobs <- j
    }
    close(jobs)

    // collect all the results of the work.
    for a := 1; a <= 6; a++ {
        <-results
    }
}