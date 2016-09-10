package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewHandler(t *testing.T) {
	require := require.New(t)

	h := New(nil)
	require.NotNil(h)
}
