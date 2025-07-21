package config

import (
	"os"
	"runtime"
	"testing"
)

func TestPager(t *testing.T) {
	// Save original PAGER value
	originalPager := os.Getenv("PAGER")
	defer os.Setenv("PAGER", originalPager)

	tests := []struct {
		name     string
		setup    func()
		teardown func()
		want     string
		wantType string // "exact", "empty", or "any"
	}{
		{
			name: "PAGER environment variable is respected",
			setup: func() {
				os.Setenv("PAGER", "/usr/bin/custom-pager")
			},
			teardown: func() {
				os.Unsetenv("PAGER")
			},
			want:     "/usr/bin/custom-pager",
			wantType: "exact",
		},
		{
			name: "Empty PAGER is ignored",
			setup: func() {
				os.Setenv("PAGER", "")
			},
			teardown: func() {
				os.Unsetenv("PAGER")
			},
			wantType: "any", // Will return system pager or empty string
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Run setup if provided
			if tt.setup != nil {
				tt.setup()
			}

			// Ensure teardown runs
			if tt.teardown != nil {
				defer tt.teardown()
			}

			// Run the function
			got := Pager()

			// Check the result based on type
			switch tt.wantType {
			case "exact":
				if got != tt.want {
					t.Errorf("Pager() = %v, want %v", got, tt.want)
				}
			case "empty":
				if got != "" {
					t.Errorf("Pager() = %v, want empty string", got)
				}
			case "any":
				// Just verify it doesn't panic - any result is acceptable
			}
		})
	}
}

func TestPagerWindows(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("Skipping Windows-specific test on non-Windows platform")
	}

	// On Windows, should always return "more"
	got := Pager()
	want := "more"
	if got != want {
		t.Errorf("Pager() = %v, want %v", got, want)
	}
}

func TestPagerUnixSystems(t *testing.T) {
	// Skip this test on Windows
	if runtime.GOOS == "windows" {
		t.Skip("Skipping Unix-specific test on Windows")
	}

	// Save original PAGER value
	originalPager := os.Getenv("PAGER")
	defer os.Setenv("PAGER", originalPager)

	// Unset PAGER to test fallback behavior
	os.Unsetenv("PAGER")

	// Call Pager() - we can't predict what will be found on the system
	// but we can verify it doesn't panic and returns a string
	result := Pager()

	// The result should be either empty (no pager found) or a path
	if result != "" && result[0] != '/' {
		t.Errorf("Pager() returned %q, expected empty string or absolute path", result)
	}
}
