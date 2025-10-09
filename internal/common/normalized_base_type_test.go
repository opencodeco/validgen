package common

import "testing"

func TestNormalizedBaseTypeString(t *testing.T) {
	tests := []struct {
		name string
		n    NormalizedBaseType
		want string
	}{
		{
			name: "StringType",
			n:    StringType,
			want: "<STRING>",
		},
		{
			name: "BoolType",
			n:    BoolType,
			want: "<BOOL>",
		},
		{
			name: "IntType",
			n:    IntType,
			want: "<INT>",
		},
		{
			name: "FloatType",
			n:    FloatType,
			want: "<FLOAT>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.String(); got != tt.want {
				t.Errorf("NormalizedBaseType.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
