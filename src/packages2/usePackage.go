package main


//you cna group imports or declare separately. best practice to group them

import(
	"fmt"
	"time"
	"math/rand"
)

import "math"




func main(){
	fmt.Println("exmaple packages")
	fmt.Println(math.Pi)
	fmt.Println(math.Sqrt(3))
	fmt.Println(rand.Intn(10000))//generate randon 
	fmt.Println(time.Now())
}