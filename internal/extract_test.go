package internal

import "testing"

func TestLibraryOptions_ExtractLibraryName(t *testing.T) {
	tests := []struct {
		Package  string
		expected string
		hasError bool
	}{
		{"github.com/vektra/mockery/v2@v2.20.0", "mockery", false},
		{"github.com/golangci/golangci-lint/cmd/golangci-lint@v1.51.0", "golangci-lint", false},
		{"github.com/swaggo/swag/cmd/swag@v1.16.1", "swag", false},
		{"invalid/package/format", "", true},
		{"github.com/username/invalidpackage@", "", true},
	}

	for _, test := range tests {
		t.Run(test.Package, func(t *testing.T) {
			output, err := ExtractLibName(test.Package)
			if test.hasError {
				if err == nil {
					t.Errorf("expected error for input %s, but got none", test.Package)
				}
			} else {
				if err != nil {
					t.Errorf("did not expect error for input %s, but got: %v", test.Package, err)
				}
				if output != test.expected {
					t.Errorf("expected %s, but got %s", test.expected, output)
				}
			}
		})
	}
}
