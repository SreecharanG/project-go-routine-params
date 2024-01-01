# project-go-routine-params
project-go-routine, going through parameters. 

* This project helps to understand concurrent code allows a program to execute multiple tasks at the same time.
* Go accomplishes this by using GoRoutines and Channels.

**Go Routines**: 
* These are light weight threads managed by the Go runtime. Goroutines allow you to run multiple functions concurrently. They can be created using the 'go' keyword.


Example: 

  ```
  go testRun("Run task completed") 
  ```

**Channels**: 
* These are the pipes that connect concurrent Goroutines. Can send values from one Goroutine to another using a channel.


Example: 

  ```
  ch := make(chan string) 
  go printString(ch)
  ch <- "Hello, World!" 
  ```

* These are several types of concurrency that Go supports

  1. **Pipelines**: Pipelines are a way to chain together functions that transform input data. They can be thought  of as a linear chain of data transformation.

Example: 

  ```
  ch := make(chan int)
  go square(ch)
  go printIntch)
  ch <- 5
  ```

  2. **FanOut / Fan In**: This is a way to parallelize tasks. One function produces data that is sent to multiple Goroutines, which process it and send it back to another function that consumes the data.

Example: 

  ```
  ch := make(chan int)
  for i := 0; i < 10; i++ {
    go read(ch)
  }
  for i := 0; i < 100; i++ {
    go write(ch, i)
  }
  ```
  3. **Worker Pools**: This is a way to manage a pool of Goroutines that are used to perform a task. This is useful for controlling resource usage:

Example: 

  ```
  type worker struct {
    in chan int
    done chan bool
  }

  func newWorker(in chan int) *worker {
    return &worker {
      in: in,
      done: make(chan bool)
    }
  }

  func (w *worker) start() {
    go func() {
        for n:= range w.in {
          fmt.Printf("Received: %d\n", n)
        }
        w.done <- true
    }()
  }

  func (w *worker) stop() {
    close(w.in)
  <- w.done
  }



  
  
