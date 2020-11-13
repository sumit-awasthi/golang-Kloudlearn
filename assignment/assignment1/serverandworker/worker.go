// workers in 
package main
 
import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
 
var (
	// wg is used to force the application wait for all goroutines to finish before exiting.
	wg sync.WaitGroup
	// jobChan is a buffered channel that has the capacity of maximum 11 resource slot.
	jobChan = make(chan int, 11)
	// waiters is used to make goroutines sleep in order to simulate time it took to process the job.
	waiters = []int{0, 1, 2, 3, 4}
)
 
func main() {
	rand.Seed(time.Now().UnixNano())
 
	fmt.Println("BEGIN")
 
	// Tell how many worker you will be running.
	wg.Add(1)
 
	// Run 1 worker to handle jobs.
	go worker(jobChan, &wg)
 
	// Send 10 jobs to job channel.
	for i := 1; i <= 2; i++ {
		if !queueJob(i, jobChan) {
			fmt.Println("Channel is full... Service unavailable...")
		}
	}
 
	// Close the job channel.
	close(jobChan)
 
	// Block exiting until all the goroutines are finished.
	wg.Wait()
 
	fmt.Println("END")
}
 
// queueJob puts job into channel. If channel buffer is full, return false.
func queueJob(job int, jobChan chan<- int) bool {
	select {
	case jobChan <- job:
		return true
	default:
		return false
	}
}
 
// worker processes jobs.
func worker(jobChan <-chan int, wg *sync.WaitGroup) {
	// As soon as the current goroutine finishes (job done!), notify back WaitGroup.
	defer wg.Done()
 
	fmt.Println("Worker is waiting for jobs")
 
	for job := range jobChan {
		fmt.Println("Worker picked job", job)
 
		wait := time.Duration(rand.Intn(len(waiters)))
		time.Sleep(wait * time.Second)
 
		// Process the job ...
 
		fmt.Println("Worker completed job", job, "in", int(wait), "second(s)")
	}
}
