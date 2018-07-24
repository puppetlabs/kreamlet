package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractJoinToken(t *testing.T) {
	tokenCreateOutput := `ilgxgd.6z328tuq2njy0u2y
`
	token, err := extractJoinToken(tokenCreateOutput)
	assert.NoError(t, err)
	assert.Equal(t, "ilgxgd.6z328tuq2njy0u2y", token)
}
