package packagetest31

import (
	"fmt"
)

var File31SomeVariable1 int = 100 //the variable is exported (can be used outside of package)
var file31SomeVariable2 int = 200 //the variable is NOT exported



//this function is exported (can be used outside of package)
func File31SomeMethod1(){
	fmt.Println("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
}

//this function is NOT exported
func file31SomeMethod2(){
	fmt.Println("abcdefghijklmnopqrstuvwxyz")
}