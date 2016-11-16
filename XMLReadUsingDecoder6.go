package main

import (
	"fmt"
	"encoding/xml"
	by "bytes"
	"os"
)

func main(){
	c1 := `<customer>
		<cid>2</cid>
		<cname>Sarath Ambadas</cname>
		<ctype>person</ctype>
		<caddress>Bangalore</caddress>
	</customer>`
	
	xmlDecoder := xml.NewDecoder(by.NewReader([]byte(c1)))
	
	var c Customer
	err := xmlDecoder.Decode(&c)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	
	fmt.Println(c.Cid)
	fmt.Println(c.Cname)
	fmt.Println(c.Ctype)
	fmt.Println(c.Caddress)

}


type Customer struct{
	XMLName xml.Name	`xml:"customer"`
	Cid int		`xml:"cid"`
	Cname string	`xml:"cname"`
	Ctype string	`xml:"ctype"`
	Caddress string	`xml:"caddress"`
}