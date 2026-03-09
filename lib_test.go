package lib

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHello(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{
			name:     "returns greeting with version",
			expected: "Hello, World! (v" + Version + ")",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Hello()
			require.Equal(t, tt.expected, got)
		})
	}
}

func TestVersionInfo(t *testing.T) {
	tests := []struct {
		name     string
		contains []string
	}{
		{
			name:     "contains version components",
			contains: []string{Version, BuildDate, CommitHash},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := VersionInfo()
			for _, s := range tt.contains {
				require.Contains(t, got, s)
			}
		})
	}
}
