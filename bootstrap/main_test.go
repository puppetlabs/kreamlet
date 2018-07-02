package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	initKubeAdminOutput = `
blah blah blah
You can now join any number of machines by running the following on each node
as root:
 
  kubeadm join 192.168.65.6:6443 --token ilgxgd.6z328tuq2njy0u2y --discovery-token-ca-cert-hash sha256:e617f08daec4c0f34a651bd709dc8168b1ffa0acdc7d22c9d4bb191778f4ab4a
  Valid certificates and keys now exist
Applying 50-weave.yaml
blah blah blah`
)

func TestExtractJoinTokenFromInitOutput(t *testing.T) {
	token, err := extractJoinTokenFromInitOutput(initKubeAdminOutput)
	assert.NoError(t, err)
	assert.Equal(t, "ilgxgd.6z328tuq2njy0u2y", token)
}

func TestExtractJoinTokenFromTokenCreate(t *testing.T) {
	tokenCreateOutput := `ilgxgd.6z328tuq2njy0u2y
`
	token, err := extractJoinTokenFromTokenCreate(tokenCreateOutput)
	assert.NoError(t, err)
	assert.Equal(t, "ilgxgd.6z328tuq2njy0u2y", token)
}
