package hashfuncs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFnv32(t *testing.T) {
	test := assert.New(t)

	bs := []byte("abc")
	test.Equal(golangFnv32(bs), fnv132(bs))
	test.Equal(golangFnv32a(bs), fnv1a32(bs))
}
