package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_compareTrue(t *testing.T) {
	var testEnvs []string
	testEnvs = append(testEnvs, "ENV_MOCK")
	testEnvs = append(testEnvs, "SECRET_MOCK")
	var secret []string
	secret = append(secret, "SECRET_MOCK")
	var iacEnv []string
	iacEnv = append(iacEnv, "ENV_MOCK")
	error := compare(testEnvs, secret, iacEnv)
	assert.False(t, error)
}

func Test_compareError(t *testing.T) {
	var testEnvs []string
	testEnvs = append(testEnvs, "ENV_MOCK")
	testEnvs = append(testEnvs, "SECRET_MOCK")
	var secret []string
	var iacEnv []string
	iacEnv = append(iacEnv, "ENV_MOCK")
	error := compare(testEnvs, secret, iacEnv)
	assert.True(t, error)
}
