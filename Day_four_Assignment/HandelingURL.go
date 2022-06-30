package main

import (
	"fmt"
	"time"
)


func display(str string) {
	for w:= 0; w<6 ; w++ {
		time.Sleep(1 * time.Second)
		fmt.Println(str)
	}
	go func(){
		fmt.Println("good")
	}()
	time.Sleep(1 * time.Second)

}

func main(){
	go display("Welcome")
	display("Go")
}