package poetpendulum_test

import (
	"testing"

	"github.com/hsmtkk/vigilant-disco/poetpendulum"
	"github.com/stretchr/testify/assert"
)

/*
   assert_eq!(pendulum(&[49,40,41,41,39,43,40,46,30,47,46,40]), [47,46,41,40,40,30,39,40,41,43,46,49]);
   assert_eq!(pendulum(&[48,41,38,35,50,46,38,42,37,49,44,32,37]), [50,48,44,41,38,37,32,35,37,38,42,46,49]);
   assert_eq!(pendulum(&[49,30,39,30,40,44,43,48,47,50,42,48,33]), [50,48,47,43,40,33,30,30,39,42,44,48,49]);
   assert_eq!(pendulum(&[48,43,35,47,39,38,38,46,49,32,42]), [49,47,43,39,38,32,35,38,42,46,48]);

*/

func TestArrange(t *testing.T) {
	assert.Equal(t, []int{50, 47, 40, 38, 33, 32, 33, 36, 38, 43, 48}, poetpendulum.Arrange([]int{33, 38, 38, 36, 43, 48, 32, 40, 47, 50, 33}))
}
