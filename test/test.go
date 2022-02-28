package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type product struct {
	XMLName  xml.Name `xml:"prod"`
	Name     string   `xml:"name"`
	Price    float64  `xml:"prx"`
	Quantity uint     `xml:"qty"`
}

type products struct {
	XMLName  xml.Name `xml:"products"`
	Products []product
}

func main() {
	ps := []product{
		{Name: "alpha", Price: 1.0, Quantity: 2},
		{Name: "bravo", Price: 3.0, Quantity: 4},
	}
	pss := products{
		Products: ps,
	}
	encoded, err := xml.Marshal(&pss)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(encoded))
}
