package main

import (
	"fmt"
	"encoding/xml"
	"os"
)

func main(){
	c1 := Customer{CustomerId:1, CustomerName:"ASK Inc", CustomerAddress:"SFO", Type:"Software", Active:true}
	
	//b1, err := xml.Marshal(c1)
	b1, err := xml.MarshalIndent(c1, "", "    ")//2n and 3rd parameters ae prefix and indent
	
	
	if err != nil{
		fmt.Println(err)
	}else{
		os.Stdout.Write(b1)
	}
}


//only upper case fieds will be serialized
type Customer struct{
	XMLName xml.Name	`xml:"person"`
	CustomerId int	`xml:"customerId"`
	CustomerName string	`xml:"customerName"`
	CustomerAddress string	`xml:"customerAddress"`
	Type string	`xml:"type"`
	Active bool `xml:"active"`
}