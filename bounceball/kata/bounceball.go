package kata

import "fmt"

type Parameter struct {
	Height float64
	Bounce float64
	Window float64
}

func BouncingBall(h, bounce, window float64) int {
	param := Parameter{Height: h, Bounce: bounce, Window: window}
	pass, err := BouncePass(param)
	if err != nil {
		return -1
	}
	return 1 + 2*pass
}

func BouncePass(param Parameter) (int, error) {
	if param.Height <= 0 {
		return 0, fmt.Errorf("height must be greater than zero")
	}
	if param.Bounce <= 0 || param.Bounce >= 1 {
		return 0, fmt.Errorf("bounce b must be 0 < b < 1")
	}
	if param.Window >= param.Height {
		return 0, fmt.Errorf("window must be less than Height")
	}
	pass := 0
	h := param.Height
	for {
		h *= param.Bounce
		if h <= param.Window {
			break
		}
		pass += 1
	}
	return pass, nil
}
