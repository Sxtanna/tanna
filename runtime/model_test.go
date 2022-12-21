package runtime

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuiltInPlatformTypes(t *testing.T) {
	assert.Equal(t, reflect.TypeOf(true),
		Bit.(*ModelPlatform).platformType)

	assert.Equal(t, reflect.TypeOf(int64(10)),
		Int.(*ModelPlatform).platformType)

	assert.Equal(t, reflect.TypeOf(float64(1.0)),
		Dec.(*ModelPlatform).platformType)

	assert.Equal(t, reflect.TypeOf('T'),
		Let.(*ModelPlatform).platformType)

	assert.Equal(t, reflect.TypeOf("Hi"),
		Txt.(*ModelPlatform).platformType)
}
