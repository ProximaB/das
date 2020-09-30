package util_test

import (
	"github.com/ProximaB/das/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInterfaceSliceToIntSlice(t *testing.T) {
	input := []interface{}{31, 5, 82, 40, 67}
	output := util.InterfaceSliceToIntSlice(input)
	assert.Equal(t, 31, output[0])
}
