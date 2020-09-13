package geohash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeoHashEncode(t *testing.T) {
	test := assert.New(t)

	cell1 := encode(121.439601, 31.193299, 6)
	test.Equal(cell1.code, "wtw37q")

	cell2 := encode(0.119, 52.205, 7)
	test.Equal(cell2.code, "u120fxw")
}

func TestGeoHashDecode(t *testing.T) {
	test := assert.New(t)

	cell0 := decode("w")
	test.Equal(cell0.minLng, 90.)
	test.Equal(cell0.maxLng, 135.)
	test.Equal(cell0.minLat, 0.)
	test.Equal(cell0.maxLat, 45.)
}

func TestGeoHashDecode2(t *testing.T) {
	test := assert.New(t)

	cell1 := encode(0.119, 52.205, 7)
	test.Equal(cell1.code, "u120fxw")

	cell2 := decode("u120fxw")
	test.Equal(cell1, cell2)
}

func TestGeoHashAdjacent(t *testing.T) {
	test := assert.New(t)

	test.Equal(adjacent("1", north), "3")
	test.Equal(adjacent("1", south), "c")
	test.Equal(adjacent("1", east), "4")
	test.Equal(adjacent("1", west), "0")

	test.Equal(adjacent("wtmk72", north), "wtmk73")
	test.Equal(adjacent("wtmk73", south), "wtmk72")

	test.Equal(adjacent("wtmk72", south), "wtmk5r")
	test.Equal(adjacent("wtmk72", east), "wtmk78")
	test.Equal(adjacent("wtmk72", west), "wtmk70")
}

func TestGeoHashNeighbour(t *testing.T) {
	test := assert.New(t)
	test.Equal(1, 1)

	neighbours := getNeighbours("wtmk72")
	test.Equal(neighbours, []string{
		"wtmk71", "wtmk73", "wtmk79",
		"wtmk70", "wtmk78",
		"wtmk5p", "wtmk5r", "wtmk5x",
	})
}
