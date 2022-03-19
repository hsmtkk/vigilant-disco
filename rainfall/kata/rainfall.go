package kata

import (
	"log"
	"strconv"
	"strings"
)

func Mean(town string, strng string) float64 {
	parsed := ParseData(strng)
	rains, ok := parsed[town]
	if ok {
		return calcMean(rains)
	} else {
		return -1
	}
}

func Variance(town string, strng string) float64 {
	parsed := ParseData(strng)
	rains, ok := parsed[town]
	if ok {
		return calcVariance(rains)
	} else {
		return -1
	}
}

func calcMean(fs []float64) float64 {
	return sum(fs) / float64(len(fs))
}

func sum(fs []float64) float64 {
	s := 0.0
	for _, f := range fs {
		s += f
	}
	return s
}

func calcVariance(fs []float64) float64 {
	s := 0.0
	m := calcMean(fs)
	for _, f := range fs {
		s += (f - m) * (f - m)
	}
	return s / float64(len(fs))
}

func ParseData(data string) map[string][]float64 {
	result := map[string][]float64{}
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		elems := strings.Split(line, ":")
		town := elems[0]
		monRains := strings.Split(elems[1], ",")
		rains := []float64{}
		for _, monRain := range monRains {
			elems := strings.Split(monRain, " ")
			rain, err := strconv.ParseFloat(elems[1], 64)
			if err != nil {
				log.Fatal(err)
			}
			rains = append(rains, rain)
		}
		result[town] = rains
	}
	return result
}
