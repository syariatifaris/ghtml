package ghtml

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getTag(t *testing.T) {
	start, end := getTag(TypeTable)
	assert.NotEmpty(t, start, "start tag not empty")
	assert.NotEmpty(t, end, "end tag not empty")
}
