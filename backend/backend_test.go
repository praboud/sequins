package backend

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBackend(t *testing.T) {
	backend := NewLocalBackend("../test_data")
	version, err := backend.LatestVersion(false)
	require.NoError(t, err)
	assert.Equal(t, version, "1")
}

func TestBackendCheckForSuccess(t *testing.T) {
	backend := NewLocalBackend("../test_data")
	version, err := backend.LatestVersion(true)
	require.NoError(t, err)
	assert.Equal(t, version, "0")
}
