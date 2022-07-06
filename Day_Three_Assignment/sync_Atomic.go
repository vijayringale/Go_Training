package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)
  

var waittime sync.WaitGroup
var atmvar int32
  
func movie(S string) {
  
    for i := 1; i < 7; i++ {
  
        time.Sleep(time.Duration((rand.Intn(5))) * time.Millisecond)
  
    
        atomic.AddInt32(&atmvar, 1)
  
        fmt.Println(S, i, "count ->", atmvar)
    }
  
    waittime.Done()
}
  
func main() {
  
    waittime.Add(2)
  
   
    go movie("bahubali 1 : ")
    go movie("bahubali 2: ")
  
    waittime.Wait()
  
    fmt.Println("The value of last count is :", atmvar)
}
