package ch5

import "testing"

func TestResProviderIsEnabled(t *testing.T) {

	resProvider := ResProvider{value: true}

	result := resProvider.IsEnabled()

	expected := true
	if result != expected {
		t.Errorf("ResProvider.IsEnabled() = %v; want  %v", result, expected)
	}

}
