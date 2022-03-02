package kata_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/bounceball/kata"
	"github.com/stretchr/testify/assert"
)

/*
   testequal(3, 0.66, 1.5, 3)
   testequal(40, 0.4, 10, 3)
   testequal(10, 0.6, 10, -1)
   testequal(40, 1, 10, -1)
   testequal(5, -1, 1.5, -1)
*/

func Test0(t *testing.T) {
	param := kata.Parameter{
		Height: 3,
		Bounce: 0.66,
		Window: 1.5,
	}
	assert.Equal(t, kata.BouncingBall(param.Height, param.Bounce, param.Window), 3)
}

func Test1(t *testing.T) {
	param := kata.Parameter{
		Height: 40,
		Bounce: 0.4,
		Window: 10,
	}
	assert.Equal(t, kata.BouncingBall(param.Height, param.Bounce, param.Window), 3)
}

func Test2(t *testing.T) {
	param := kata.Parameter{
		Height: 10,
		Bounce: 0.6,
		Window: 10,
	}
	assert.Equal(t, kata.BouncingBall(param.Height, param.Bounce, param.Window), -1)
}

func Test3(t *testing.T) {
	param := kata.Parameter{
		Height: 40,
		Bounce: 1,
		Window: 10,
	}
	assert.Equal(t, kata.BouncingBall(param.Height, param.Bounce, param.Window), -1)
}

func Test4(t *testing.T) {
	param := kata.Parameter{
		Height: 5,
		Bounce: -1,
		Window: 1.5,
	}
	assert.Equal(t, kata.BouncingBall(param.Height, param.Bounce, param.Window), -1)
}

func Test5(t *testing.T) {
	param := kata.Parameter{
		Height: 2,
		Bounce: 0.5,
		Window: 1,
	}
	assert.Equal(t, kata.BouncingBall(param.Height, param.Bounce, param.Window), 3)
}
