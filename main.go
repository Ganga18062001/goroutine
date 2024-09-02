package main

import (
	"fmt"
	"strconv"
	"sync"
	//cls
	//"time"
	//"time"
	//"time"
)

// func Demo(){
// 	fmt.Println("Go rutine is light weight thred of go programming language")
// 	// time.Sleep(100)
// }
func Greet(name string){
	fmt.Println("My name is " + strconv.Itoa(10))
	//defer tg.Done()
}
func Demo(count int, wg *sync.WaitGroup){
	fmt.Println("welcome " + strconv.Itoa(count))
	wg.Done()
}

func main(){
	var wg sync.WaitGroup

	for _,i := range []int{1,2,3,4,5}{

		go Demo(i,&wg)
		wg.Add(1)


	}
	wg.Wait()
	fmt.Println("Program was executed")




	// var tg sync.WaitGroup
	// var tr sync.WaitGroup
	// tg.Add(2)
	// tr.Add(1)

	// func(tr *sync.WaitGroup){
	// 	fmt.Println("for example")
	// 	defer tr.Done()
	// }(&tr)

	// // go Greet("Raju")
	// // time.Sleep(1 * time.Second)

	//  go func(tg *sync.WaitGroup){
	// 	fmt.Println("welcome  in anonymous function")
	// 	defer tg.Done()

	// }(&tg)
	// go func(name string,tg *sync.WaitGroup){
	// 	fmt.Println("welcome " + name)
	// 	tg.Done()
	// }("Ganga",&tg)
	// tg.Wait()
	// fmt.Println("In main function")








	//go Greet("Ran")

	// go func(){
	// 	fmt.Println("Anonymous funnction")
	// }()
	// time.Sleep()

	// time.Sleep(1 * time.Second)


	// go Demo()
	// time.Sleep(1 * time.Second)
	// go Greet("Ganga")
	// time.Sleep(1 * time.Second)

	// fmt.Println("call in main function")
	// time.Sleep(1 * time.Second)
	// time.Sleep(100)


}