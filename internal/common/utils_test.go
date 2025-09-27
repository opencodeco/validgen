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

func TestIsGoType(t *testing.T) {
	type args struct {
		fieldType FieldType
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "basic type",
			args: args{
				fieldType: FieldType{BaseType: "string", ComposedType: "", Size: ""},
			},
			want: true,
		},
		{
			name: "array type",
			args: args{
				fieldType: FieldType{BaseType: "string", ComposedType: "[N]", Size: "5"},
			},
			want: true,
		},
		{
			name: "slice type",
			args: args{
				fieldType: FieldType{BaseType: "string", ComposedType: "[]", Size: ""},
			},
			want: true,
		},
		{
			name: "map type",
			args: args{
				fieldType: FieldType{BaseType: "string", ComposedType: "map", Size: ""},
			},
			want: true,
		},
		{
			name: "custom type",
			args: args{
				fieldType: FieldType{BaseType: "MyType", ComposedType: "", Size: ""},
			},
			want: false,
		},
		{
			name: "pointer type",
			args: args{
				fieldType: FieldType{BaseType: "*MyType", ComposedType: "", Size: ""},
			},
			want: false,
		},
		// {
		// 	name: "pointer to basic type",
		// 	args: args{
		// 		fieldType: FieldType{BaseType: "*string", ComposedType: "", Size: ""},
		// 	},
		// 	want: true,
		// },
		{
			name: "struct type",
			args: args{
				fieldType: FieldType{BaseType: "MyStruct", ComposedType: "", Size: ""},
			},
			want: false,
		},
		// {
		// 	name: "interface type",
		// 	args: args{
		// 		fieldType: FieldType{BaseType: "interface{}", ComposedType: "", Size: ""},
		// 	},
		// 	want: true,
		// },
		// {
		// 	name: "complex type",
		// 	args: args{
		// 		fieldType: FieldType{BaseType: "map[string][]*MyType", ComposedType: "", Size: ""},
		// 	},
		// 	want: true,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsGoType(tt.args.fieldType); got != tt.want {
				t.Errorf("IsGoType() = %v, want %v", got, tt.want)
			}
		})
	}
}
