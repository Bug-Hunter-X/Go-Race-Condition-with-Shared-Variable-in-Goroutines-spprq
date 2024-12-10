# Go Race Condition Example

This repository demonstrates a common race condition in Go programs involving shared variables accessed by multiple goroutines.

## The Problem
The `bug.go` file contains a program that uses 10 goroutines to increment a shared counter variable.  Because of the lack of synchronization mechanisms (mutexes or channels), the final value of the counter is not reliably 10.  This is because the increment operation (`count++`) is not atomic; multiple goroutines can try to increment the counter simultaneously, leading to data races and incorrect results. 

## The Solution
The `bugSolution.go` file provides a corrected version that uses a mutex to protect the shared variable, ensuring that only one goroutine can access and modify the variable at a time. This prevents race conditions and guarantees the expected final count.

## How to run
1. Clone this repository.
2. Navigate to the repository directory.
3. Run `go run bug.go` to observe the race condition in action.
4. Run `go run bugSolution.go` to see the corrected code in action.