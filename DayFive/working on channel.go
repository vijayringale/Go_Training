package main

import "fmt"

// func portal1(channel1 chan string) {

// 	channel1 <- " Welcome "

// }

// func main() {
// 	ch1 := make(chan string)
// 	go portal1(ch1)
// 	select{
// 	case o1 := <-ch1:
// 	fmt.Println(o1)
// 	}

// }

// func A1() {
// 	arr1:= [4]string{"Amit","Ajit","Avinash","Arvind"}

// 	for t1:= 0;t1<4;t1++ {
// 		time.Sleep(100 *time.Millisecond)
// 		fmt.Println(arr1)
// 	}
// }

// func A2(){
// 	arr2:= [5]string{"vinod","vishnu","vinita","vinit","vilas"}

// 	for t1:=0; t1<5;t1++ {
// 		time.Sleep(500 * time.Millisecond)
// 		fmt.Println(arr2)

// 	}
// }

// func main(){
// 	fmt.Println("! ... Main GO-routine Start ...!" )
// 	go A1()
// 	go A2()

// 	time.Sleep(4000 * time.Millisecond)
// 	fmt.Println("\n .. Main Go-routine End...!")
// }

// func main(){
// 	var mychannel chan int
// 	fmt.Println("Value of the channel: ", mychannel)
// 	fmt.Printf("Type of the channel : %T", mychannel)

// 	channel:= make(chan int)
// 	fmt.Println("Value Of the channel: ",channel)
// 	fmt.Printf("type of the channel1: %T",channel)
// }

// func myfun(ch chan int){
// 	fmt.Println(44 + <-ch)
// }
// func main(){
// 	fmt.Println("start Main Method")
// 	ch := make(chan int)

// 	go myfun(ch)
// 	ch<-12
// }

// func myfun(ch1 chan string){

// 	for v:=0;v<5;v++ {
// 		ch1<- "king"
// 	}
// 	close(ch1)
// }

// func main(){
// 	c := make(chan string)
// 	go myfun(c)
// 	for {
// 		res, ok:= <-c
// 		if ok == false {
// 		fmt.Println(res,ok)
// 		break
// 		}
// 		fmt.Println("Channel Open", res , ok)
// 	}
// }

// func main(){

// 	mychnl := make(chan string)

// 	go func(){
// 		mychnl<-"kgf"
// 		mychnl<-"pushpa"
// 		mychnl<-"RRR"
// 		mychnl<-"bahubali"
// 		close(mychnl)
// 	}()

// 	for res := range mychnl {
// 		fmt.Println(res)
// 	}

// }

// func main() {
// 	Bahubali := make(chan string , 4)
// 	Bahubali<-"BhalalDev"
// 	Bahubali<-"Shivgami"
// 	Bahubali<-"Katapa"
// 	Bahubali<-"Bahubali"
// 	close(Bahubali)
// 	fmt.Println("Lenth Of Channel: ",len(Bahubali))
// 	fmt.Println("Capacity Of Channel: ",cap(Bahubali))
// 	func () {

// 		for pintu := range Bahubali{

// 			fmt.Println("Charecter in Bahubali Move :-",pintu)

// 		}
// 	}()

// }

// func main() {
// 	x := 0xff
// 	y := 0x9c

// 	fmt.Printf("Type of variable x is %T\n",x)
// 	fmt.Printf("Values of X in Hexadecimal is %X \n",x)
// 	fmt.Printf("Value of X in decimal is %v \n", x)

// 	fmt.Printf("Type of variable y is %T \n", y )
// 	fmt.Printf("Value of y in hexadecimal is %X\n", y)
// 	fmt.Printf("Value of y in decimal is %v\n ",y)
// }

// func main(){

// 	var x int = 5748
// 	var p *int
// 	p = &x

// 	fmt.Println("Value stored in x = ",x)
// 	fmt.Println("Address of x =", &x)
// 	fmt.Println("Value stored in variable p =", p)
// }

// func main(){
// 	var y = 456
// 	var p = &y

// 	fmt.Println("Values stored in y = ",y)
// 	fmt.Println("Address of y = ",&y)
// 	fmt.Println("Value stored in pointer variable p = ",p)

// 	fmt.Println("Value store in y(*p) = ",*p)
// }

// func main(){

// 	var y = 458
// 	var p = &y
// 	fmt.Println("Value stored in y before changing = ", y)
// 	fmt.Println("Address of y = ",&y)
// 	fmt.Println("Value stored in pointer variable p = ",p)

// 	fmt.Println("Value stored in y(*p) Before Changing = ", *p)

// 	*p = 500

// 	fmt.Println("Value stored in y(*p) after Changing = ",y)

// }

// func fibo(n int) {
// 	a:=0
// 	b:=1
// 	c:=b
// 	for true{
// 		c=b
// 		b=a+b
// 		if b>n{
// 			break
// 		}
// 		a=c

// 	fmt.Print(b)
// 	}

// }
// func main(){
// 	fibo(5)
// }

// func main(){
// 	t1:=0
// 	t2:=1
// 	t3:=0

// 	for i:=0;i<5;i++ {
// 		t3 =t1+t2
// 		t1=t2
// 		t2=t3
// 		fmt.Print(t3)
// 	}
// }

func fact(i int) int {
	if i<=0{
		return 1
	}
	fmt.Print(i)
	return i * fact(i-1)

}
func main(){

fmt.Print(fact(4))
}