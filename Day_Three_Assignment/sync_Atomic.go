package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)
  

var waittime sync.WaitGroup
  
var atmvar int32
  
func hike(S string) {
  
    for i := 1; i < 7; i++ {
  
        
        time.Sleep(time.Duration( time.Millisecond))
  
        
        atomic.AddInt32(&atmvar, 1)
  
        
        fmt.Println(S, i, "count ->", atmvar)
    }
  
    
    waittime.Done()
}
  

func main() {
  
    
    waittime.Add(2)
  
   
    go hike("Pushpa: ")
    go hike("Bahubali: ")
  
    
    waittime.Wait()
  
    fmt.Println("The value of last count is :", atmvar)
}