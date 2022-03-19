package kata

import (
	"math/big"
	"strconv"
)

func NbDig(n int, d int) int {
	count := 0
	for k := 0; k <= n; k++ {
		sq := squareString(k)
		for _, s := range sq {
			if string(s) == strconv.Itoa(d) {
				count += 1
			}
		}
	}
	return count
}

func squareString(k int) string {
	bigk := big.NewInt(int64(k))
	sq := bigk.Mul(bigk, bigk)
	return sq.String()
}
