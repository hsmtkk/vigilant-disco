package kata

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func StockList(listArt []string, listCat []string) string {
	categoryStocks := []CategoryStock{}
	for _, art := range listArt {
		cs := CategoryStockFromString(art)
		categoryStocks = append(categoryStocks, cs)
	}
	result := map[string]uint{}
	for _, cat := range listCat {
		result[cat] = 0
	}
	for _, cat := range listCat {
		for _, cs := range categoryStocks {
			if cat == cs.Category {
				result[cat] += cs.Stock
			}
		}
	}
	ret := []string{}
	for _, cat := range listCat {
		ret = append(ret, fmt.Sprintf("(%s : %d)", cat, result[cat]))
	}
	return strings.Join(ret, " - ")
}

type CategoryStock struct {
	Category string
	Stock    uint
}

func CategoryStockFromString(s string) CategoryStock {
	category := s[0]
	elems := strings.Split(s, " ")
	stock, err := strconv.Atoi(elems[1])
	if err != nil {
		log.Fatalf("failed to parse as int; %s; %s", elems[1], err)
	}
	return CategoryStock{Category: string(category), Stock: uint(stock)}
}
