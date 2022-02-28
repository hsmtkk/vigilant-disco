package kata

import (
	"encoding/xml"
	"fmt"
	"log"
)

func Catalog(s string, article string) string {
	products := products{}
	if err := xml.Unmarshal([]byte(s), &products); err != nil {
		log.Fatalf("failed to parse XML; %s", err)
	}
	log.Print(products)
	for _, p := range products.Products {
		if article == p.Name {
			return fmt.Sprintf("%s > prx: $%f qty: %d", article, p.Price, p.Quantity)
		}
	}
	return ""
}

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
