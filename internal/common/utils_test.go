package common

import (
	"testing"
)

func TestExtractPackage(t *testing.T) {
	type args struct {
		fieldType string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "no package",
			args: args{
				fieldType: "string",
			},
			want: "",
		},
		{
			name: "with package",
			args: args{
				fieldType: "mypkg.MyType",
			},
			want: "mypkg",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtractPackage(tt.args.fieldType); got != tt.want {
				t.Errorf("ExtractPackage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeyPath(t *testing.T) {
	type args struct {
		values []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "single value",
			args: args{
				values: []string{"field"},
			},
			want: "field",
		},
		{
			name: "multiple values",
			args: args{
				values: []string{"field1", "field2", "field3"},
			},
			want: "field1.field2.field3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KeyPath(tt.args.values...); got != tt.want {
				t.Errorf("KeyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
