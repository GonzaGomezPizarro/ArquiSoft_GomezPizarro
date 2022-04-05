package sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	for i := -8; i < 12; i++ {
		for j := 12; j > -8; j-- {
			assert.Equal(t, float64(i+j), Sum(float64(i), float64(j)))
		}
	}

}
