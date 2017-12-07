package main

//best practice to group the imports
import (
	"fmt"
	"time"
	"math/rand" //generates psuedo random numbers
	"packagetest31"
	"packagetest32"
)

import "math"


//packages can be clubbed together and declared in one import statement or they can be declared in separate import statements
//best practice is to group them and declare in one import statement

//any variables and functions which are starting in upper case are exported by default so that they can be used in OTHER PACKAGES

func main(){
	fmt.Println("import statement test")
	
	fmt.Println(math.Pi)
	fmt.Println(math.Sqrt(3))
	fmt.Println(rand.Int())
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Float32())
	fmt.Println(rand.Float64())
	fmt.Println(time.Now())
	
	fmt.Println("******************************")
	
	fmt.Println(packagetest31.File31SomeVariable1)
	fmt.Println(packagetest32.File32SomeVariable3)
	
	//fmt.Println(packagetest31.file31SomeVariable2)//unexported variables
	//fmt.Println(packagetest32.file32SomeVariable4)//unexported variables
	
	packagetest31.File31SomeMethod1()
	packagetest32.File32SomeMethod3()

}