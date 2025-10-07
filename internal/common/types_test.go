package common

import (
	"testing"
)

func TestFieldTypeToString(t *testing.T) {
	type fields struct {
		ComposedType string
		BaseType     string
		Size         string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "base type",
			fields: fields{
				ComposedType: "",
				BaseType:     "string",
				Size:         "",
			},
			want: "string",
		},
		{
			name: "array type",
			fields: fields{
				ComposedType: "[N]",
				BaseType:     "string",
				Size:         "5",
			},
			want: "[N]string",
		},
		{
			name: "slice type",
			fields: fields{
				ComposedType: "[]",
				BaseType:     "string",
				Size:         "",
			},
			want: "[]string",
		},
		{
			name: "map type",
			fields: fields{
				ComposedType: "map",
				BaseType:     "string",
				Size:         "",
			},
			want: "map[string]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ft := FieldType{
				ComposedType: tt.fields.ComposedType,
				BaseType:     tt.fields.BaseType,
				Size:         tt.fields.Size,
			}
			if got := ft.ToGenericType(); got != tt.want {
				t.Errorf("FieldType.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFieldTypeIsGoType(t *testing.T) {
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
		{
			name: "struct type",
			args: args{
				fieldType: FieldType{BaseType: "MyStruct", ComposedType: "", Size: ""},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.fieldType.IsGoType(); got != tt.want {
				t.Errorf("IsGoType() = %v, want %v", got, tt.want)
			}
		})
	}
}
