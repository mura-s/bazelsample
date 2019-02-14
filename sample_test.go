package bazelsample

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSampleSuccess(t *testing.T) {
	str, err := Sample(true)
	assert.NoError(t, err)
	assert.Equal(t, "Success", str)
}

func TestSampleFail(t *testing.T) {
	_, err := Sample(false)
	assert.Error(t, err)
}
