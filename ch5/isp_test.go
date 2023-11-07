package ch5

import "testing"

func TestIsEnabled(t *testing.T) {

	testCases := []struct {
		description string
		provider    BooleanProvider
		expected    bool
	}{
		{
			description: "ResProvider with value true",
			provider:    ResProvider{value: true},
			expected:    true,
		}, {
			description: "ResProvider with value false",
			provider:    ResProvider{value: false},
			expected:    false,
		}, {
			description: "SoundProvider with value true",
			provider:    SoundProvider{value: true},
			expected:    true,
		}, {
			description: "SoundProvider with value false",
			provider:    SoundProvider{value: false},
			expected:    false,
		}, {
			description: "VibrationProvider with value true",
			provider:    VibrationProvider{value: true},
			expected:    true,
		}, {
			description: "VibrationProvider with value false",
			provider:    VibrationProvider{value: false},
			expected:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			result := tc.provider.IsEnabled()
			if result != tc.expected {
				t.Errorf("%s: got %v, want %v", tc.description, result, tc.expected)
			}
		})
	}

}
