package codegen

import (
	"testing"

	"github.com/seambiz/varcaser/varcaser"
)

func TestVarcaser(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"test", "Test"},
		{"R3Plant", "R3Plant"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			strCase, err := varcaser.Detect([]string{tc.input})
			if err != nil {
				strCase = varcaser.LowerCamelCaseKeepCaps
			}

			actual := varcaser.Caser{From: strCase, To: varcaser.UpperCamelCase}.String(tc.input)
			if actual != tc.expected {
				t.Errorf("expected %s, got %s", tc.expected, actual)
			}
		})
	}
}
