package main

import(
	"fmt"
	"encoding/xml"
)

func main(){
	c1 := Contact{}
	c1.FirstName = "Sarath"
	c1.LastName = "Ambadas"
	c1.Phone = "234-333-6699"
	c1.Address = Address{"100 airport blvd", "SFO", "CA", "92323"}
	
	fmt.Println(c1)
	
	b1,err := xml.MarshalIndent(&c1, "", "   ")
	if err != nil{
		fmt.Println(err)
	}else{
	
		fmt.Println(string(b1))
	}
}

//only upper case fieds will be serialized
type Contact struct{
	XMLName xml.Name	`xml:"contact"`
	FirstName string	`xml:"firstName"`
	LastName string		`xml:"lastName"`
	Phone string		`xml:"phone"`
	Address
}

//only upper case fieds will be serialized
type Address struct{
	Street string
	City string
	State string
	Zip string
}
