package simplecipher

import (
	"fmt"
	"log"
	"strings"
)

type Cipher interface {
	Encode(text string) string
	Decode(text string) string
}

func New(map1, map2 string) Cipher {
	m1 := []string{}
	m2 := []string{}
	for _, x := range map1 {
		m1 = append(m1, string(x))
	}
	for _, x := range map2 {
		m2 = append(m2, string(x))
	}
	return &cipherImpl{map1: m1, map2: m2}
}

type cipherImpl struct {
	map1 []string
	map2 []string
}

func (c *cipherImpl) Encode(text string) string {
	encoded := []string{}
	for _, ch := range text {
		index, err := indexOf(string(ch), c.map1)
		if err != nil {
			encoded = append(encoded, string(ch))
		} else {
			encoded = append(encoded, c.map2[index])
		}
	}
	return strings.Join(encoded, "")
}

func (c *cipherImpl) Decode(text string) string {
	decoded := []string{}
	for _, ch := range text {
		index, err := indexOf(string(ch), c.map2)
		if err != nil {
			log.Print(err)
			decoded = append(decoded, string(ch))
		} else {
			decoded = append(decoded, c.map1[index])
		}
	}
	return strings.Join(decoded, "")
}

func indexOf(s string, ss []string) (int, error) {
	for i, t := range ss {
		if s == t {
			return i, nil
		}
	}
	return 0, fmt.Errorf("failed to find; %s", s)
}
