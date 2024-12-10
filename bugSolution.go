```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        var count int
        var mutex sync.Mutex

        for i := 0; i < 10; i++ {
                wg.Add(1)
                go func() {
                        defer wg.Done()
                        mutex.Lock()
                        count++
                        mutex.Unlock()
                }()
        }

        wg.Wait()
        fmt.Println(count) // Output: 10
}
```