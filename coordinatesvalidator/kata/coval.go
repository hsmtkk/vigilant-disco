package kata

import (
	"log"
	"strconv"
	"strings"
)

func IsValidCoordinates(coordinates string) bool {
	elems := strings.Split(coordinates, ",")
	if len(elems) != 2 {
		log.Printf("format error: %s", coordinates)
		return false
	}
	if strings.Contains(strings.ToLower(coordinates), "e") {
		log.Printf("include `e`: %s", coordinates)
		return false
	}
	lat, err := strconv.ParseFloat(strings.TrimSpace(elems[0]), 64)
	if err != nil {
		log.Printf("latitude not number: %s", elems[0])
		return false
	}
	long, err := strconv.ParseFloat(strings.TrimSpace(elems[1]), 64)
	if err != nil {
		log.Printf("longitude not number: %s", elems[1])
		return false
	}
	if lat < -90 || lat > 90 {
		log.Printf("latitude is not in range: %f", lat)
		return false
	}
	if long < -180 || long > 180 {
		log.Printf("longitude is not in range: %f", long)
		return false
	}
	return true
}
