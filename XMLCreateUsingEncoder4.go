package main

import (
	"fmt"
	"encoding/xml"
	"os"
)

func main(){
	a1 := Account{Id:1, Name:"Amazon Inc", Type:"Software", Active:true}
	
	xmlEncoder := xml.NewEncoder(os.Stdout)
	xmlEncoder.Indent("", "    ")
	err := xmlEncoder.Encode(a1)
	if err != nil{
		fmt.Println(err)
	}
	
	fmt.Println("**********************************************")
	
	//writing the xml to the file
	writer, err1 := os.Create("account.xml")
	if err1 != nil{
		fmt.Println(err1)
	}else{
		xmlEncoder = xml.NewEncoder(writer)
		xmlEncoder.Indent("", "    ")
		xmlEncoder.Encode(a1)
	}
}


type Account struct{
	XMLName xml.Name `xml:"account"`
	Id int `xml:"id"`
	Name string `xml:"name"`
	Type string `xml:"type"`
	Active bool `xml:"active"`
}