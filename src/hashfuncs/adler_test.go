package hashfuncs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdler32V0(t *testing.T) {
	test := assert.New(t)

	ret := adler32V0([]byte("Wikipedia"))
	test.Equal(ret, uint32(0x11E60398))
}

func TestAdler32(t *testing.T) {
	test := assert.New(t)

	ret := adler32([]byte("Wikipedia"))
	test.Equal(ret, uint32(0x11E60398))
}
