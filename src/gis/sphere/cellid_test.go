package sphere

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCellId(t *testing.T) {
	test := assert.New(t)
	test.Equal(1, 1)

	cid := cellIDFromFaceIJ(1, 711197487, 903653800)
	test.Equal(uint64(cid), uint64(3958611028950762539))
	test.Equal(cellID2BitString(cid), "0011 0110 1110 1111 1100 1111 1100 0001 1101 1000 1000 1101 1100 0100 0010 1011")

	fmt.Printf("%d\n", cid)
	fmt.Println(cellID2BitString(cid))
}
