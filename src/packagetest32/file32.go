package packagetest32

import (
	"fmt"
)

var File32SomeVariable3 int = 300 //the variable is exported  (can be used outside of package)
var file32SomeVariable4 int = 400 //the variable is NOT exported

//this function is exported (can be used outside of package)
func File32SomeMethod3(){
	fmt.Println("ABCDEFGHIJKLMNOPQRSTUVWXYZ 3")
}

//this function is NOT exported
func file32SomeMethod4(){
	fmt.Println("abcdefghijklmnopqrstuvwxyz 4")
}