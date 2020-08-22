package sequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBM(t *testing.T) {
	test := assert.New(t)
	test.Equal(1, 1)

	finder := makeStringFinder("EXAMPLE")
	idx := finder.next("HERE IS A SIMPLE EXAMPLE")
	test.Equal(idx, 17)
}
