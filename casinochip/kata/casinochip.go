package kata

var solver *Solver

func init(){
	solver = newSolver()
}

func Solve(arr []int) int {
	return solver.Solve(arr)
	/*
		if arr[0] == 0 && arr[1] == 0 {
			return 0
		}
		if arr[0] == 0 && arr[2] == 0 {
			return 0
		}
		if arr[1] == 0 && arr[2] == 0 {
			return 0
		}
		max := 0
		if arr[0] > 0 && arr[1] > 0 {
			x := 1 + Solve([]int{arr[0] - 1, arr[1] - 1, arr[2]})
			max = x
		}
		if arr[0] > 0 && arr[2] > 0 {
			x := 1 + Solve([]int{arr[0] - 1, arr[1], arr[2] - 1})
			if x > max {
				max = x
			}
		}
		if arr[1] > 0 && arr[2] > 0 {
			x := 1 + Solve([]int{arr[0], arr[1] - 1, arr[2] - 1})
			if x > max {
				max = x
			}
		}
		return max
	*/
}

type Array struct {
	num0 int
	num1 int
	num2 int
}

type Solver struct {
	memo map[Array]int
}

func newSolver() *Solver {
	solver := &Solver{}
	solver.memo = map[Array]int{}
	return solver
}

func (s *Solver) Solve(arr []int) int {
	key := Array{arr[0], arr[1], arr[2]}
	ans, ok := s.memo[key]
	if ok {
		return ans
	}

	if arr[0] == 0 && arr[1] == 0 {
		return 0
	}
	if arr[0] == 0 && arr[2] == 0 {
		return 0
	}
	if arr[1] == 0 && arr[2] == 0 {
		return 0
	}
	max := 0
	if arr[0] > 0 && arr[1] > 0 {
		x := 1 + s.Solve([]int{arr[0] - 1, arr[1] - 1, arr[2]})
		max = x
	}
	if arr[0] > 0 && arr[2] > 0 {
		x := 1 + s.Solve([]int{arr[0] - 1, arr[1], arr[2] - 1})
		if x > max {
			max = x
		}
	}
	if arr[1] > 0 && arr[2] > 0 {
		x := 1 + s.Solve([]int{arr[0], arr[1] - 1, arr[2] - 1})
		if x > max {
			max = x
		}
	}
	s.memo[key] = max
	return max
}
