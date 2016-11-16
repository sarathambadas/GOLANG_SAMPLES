package main

import (
	"fmt"
	"encoding/xml"
	"os"
)

func main(){
	p1 := Parent{}
	p1.Field1 = 1
	p1.Field2 = "field2Value"
	p1.Field3 = true
	p1.Field4 = float64(3.1415)
	p1.Field5 = "field5 value"
	p1.Field6 = "field6 value"
	p1.Field7 = "field7 value"
	p1.MoreParentFields = MoreParentFields{"field8 value", "field9 value", "field10 value"}
	
	fmt.Println(p1)
	
	b1, err := xml.MarshalIndent(&p1, "", "    ")
	if err!= nil{
		fmt.Println(err)
	}else{
		fmt.Println(string(b1))
		fmt.Println("*************************************************************")
		fmt.Println(os.Stdout.Write(b1))
	}
}


type Parent struct{
	XMLName xml.Name	`xml:"parent"`
	Field1 int		`xml:"field1"`
	Field2 string	`xml:"field2"`
	Field3 bool		`xml:"field3"`
	Field4 float64	`xml:"field4"`
	
	Field5 string	`xml:"child>field5"`
	Field6 string	`xml:"child>field6"`
	Field7 string	`xml:"child>field7"`
	MoreParentFields
}

type MoreParentFields struct{
	Field8 string
	Field9 string
	Field10 string
}