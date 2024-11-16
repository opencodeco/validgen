package main

import "testing"

func TestPackageDefinitionsGenerate(t *testing.T) {
	tests := []struct {
		name    string
		pd      *PackageDefinitions
		want    string
		wantErr bool
	}{
		{
			name: "Valid package definition",
			pd: &PackageDefinitions{
				PackageName: "main",
			},
			want: `package main

import (
	"errors"
)

var ErrValidation = errors.New("validation error")
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pd := &PackageDefinitions{
				PackageName: tt.pd.PackageName,
			}
			got, err := pd.Generate()
			if (err != nil) != tt.wantErr {
				t.Errorf("PackageDefinition.Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PackageDefinition.Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}
