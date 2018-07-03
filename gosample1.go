package main

import ("fmt"
		"math"
		"math/rand")


//function to calculate suare root
func foo(){
	fmt.Println("the square root of 4 is ", math.Sqrt(4))
}
 //function for generating the random number
func random(){
	fmt.Println("A number from 1-100",rand.Intn(100))
}

func main(){
	fmt.Println("Welcome to GO")
	foo()
	random()
}