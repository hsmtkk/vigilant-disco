package muldigsum

import (
	"log"
	"strconv"
)

func Multiples(n uint) []uint {
	ms := []uint{}
	var i uint = 1
	for i = 1; ; i++ {
		m := n * i
		if m > 100 {
			break
		} else {
			ms = append(ms, m)
		}
	}
	return ms
}

func DigitSum(n uint) uint {
	var s uint = 0
	dig := strconv.Itoa(int(n))
	for _, ch := range dig {
		m, err := strconv.Atoi(string(ch))
		if err != nil {
			log.Fatalf("failed to parse %s as int; %s", string(ch), err)
		}
		s += uint(m)
	}
	return s
}

func Sum(ns []uint) uint {
	var s uint = 0
	for _, n := range ns {
		s += n
	}
	return s
}

func MulDigSum(n uint) uint {
	ms := Multiples(n)
	ns := []uint{}
	for _, m := range ms {
		ns = append(ns, DigitSum(m))
	}
	return Sum(ns)
}
