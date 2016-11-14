package goutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrimeSieve(t *testing.T) {
	x := PrimeSieve(1000)
	for _, e := range x {
		assert.True(t, IsPrime(e), "Not prime element")
	}
}
