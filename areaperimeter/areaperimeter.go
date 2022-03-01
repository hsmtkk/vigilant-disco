package areaperimeter

type Calculator interface {
	Calculate(length, width uint) uint
}

func NewCalculator() Calculator {
	return &impl{}
}

type impl struct {
}

func (c *impl) Calculate(length, width uint) uint {
	if length == width {
		return length * length
	} else {
		return 2 * (length + width)
	}
}
