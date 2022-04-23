package sqdigseq

import (
	"log"
	"strconv"
)

type Solver interface {
	Solve(int) int
}

func New() Solver {
	return &solverImpl{}
}

type solverImpl struct{}

func (s *solverImpl) Solve(n int) int {
	numbers := []int{n}
	current := n
	for {
		current = next(current)
		if contains(numbers, current) {
			break
		}
		numbers = append(numbers, current)
	}
	return len(numbers) + 1
}

func contains(numbers []int, number int) bool {
	for _, n := range numbers {
		if n == number {
			return true
		}
	}
	return false
}

func next(n int) int {
	sum := 0
	for _, d := range strconv.Itoa(n) {
		m, err := strconv.Atoi(string(d))
		if err != nil {
			log.Fatal(err)
		}
		sum += m * m
	}
	return sum
}
