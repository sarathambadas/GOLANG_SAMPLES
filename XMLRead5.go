package main

import(
	"fmt"
	"encoding/xml"
	"os"
)

func main(){
	c1 := `<customer>
				<cid>1</cid>
				<cname>Infosys</cname>
				<ctype>Software</ctype>
				<caddress>Bangalore</caddress>
		</customer>`
	
	fmt.Println(c1)
	
	var v1 Customer
	
	err := xml.Unmarshal([]byte(c1), &v1)
	
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	
	fmt.Println(v1)
	fmt.Println(v1.Cid)
	fmt.Println(v1.Cname)
	fmt.Println(v1.Ctype)
	fmt.Println(v1.Caddress)

}



type Customer struct{
	XMLName xml.Name	`xml:"customer"`
	Cid int		`xml:"cid"`
	Cname string	`xml:"cname"`
	Ctype string	`xml:"ctype"`
	Caddress string	`xml:"caddress"`
}