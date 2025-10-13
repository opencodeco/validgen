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
			name: "string base type",
			fields: fields{
				ComposedType: "",
				BaseType:     "string",
				Size:         "",
			},
			want: "string",
		},
		{
			name: "int base type",
			fields: fields{
				ComposedType: "",
				BaseType:     "int",
				Size:         "",
			},
			want: "int",
		},
		{
			name: "float base type",
			fields: fields{
				ComposedType: "",
				BaseType:     "float64",
				Size:         "",
			},
			want: "float64",
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
			name: "string type",
			args: args{
				fieldType: FieldType{BaseType: "string", ComposedType: "", Size: ""},
			},
			want: true,
		},
		{
			name: "int type",
			args: args{
				fieldType: FieldType{BaseType: "int", ComposedType: "", Size: ""},
			},
			want: true,
		},
		{
			name: "float type",
			args: args{
				fieldType: FieldType{BaseType: "float64", ComposedType: "", Size: ""},
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

func TestFieldType_NormalizeBaseType(t *testing.T) {
	type fields struct {
		BaseType string
	}
	tests := []struct {
		name   string
		fields fields
		want   NormalizedBaseType
	}{
		{
			"string type",
			fields{
				BaseType: "string",
			},
			StringType,
		},
		{
			"bool type",
			fields{
				BaseType: "bool",
			},
			BoolType,
		},
		{
			"int type",
			fields{
				BaseType: "int",
			},
			IntType,
		},
		{
			"int8 type",
			fields{
				BaseType: "int8",
			},
			IntType,
		},
		{
			"int16 type",
			fields{
				BaseType: "int16",
			},
			IntType,
		},
		{
			"int32 type",
			fields{
				BaseType: "int32",
			},
			IntType,
		},
		{
			"int64 type",
			fields{
				BaseType: "int64",
			},
			IntType,
		},
		{
			"uint type",
			fields{
				BaseType: "uint",
			},
			IntType,
		},
		{
			"uint8 type",
			fields{
				BaseType: "uint8",
			},
			IntType,
		},
		{
			"uint16 type",
			fields{
				BaseType: "uint16",
			},
			IntType,
		},
		{
			"uint32 type",
			fields{
				BaseType: "uint32",
			},
			IntType,
		},
		{
			"uint64 type",
			fields{
				BaseType: "uint64",
			},
			IntType,
		},
		{
			"float32 type",
			fields{
				BaseType: "float32",
			},
			FloatType,
		},
		{
			"float64 type",
			fields{
				BaseType: "float64",
			},
			FloatType,
		},
		{
			name: "custom type",
			fields: fields{
				BaseType: "MyType",
			},
			want: InvalidType,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ft := FieldType{
				ComposedType: "",
				BaseType:     tt.fields.BaseType,
				Size:         "",
			}
			if got := ft.NormalizeBaseType(); got != tt.want {
				t.Errorf("FieldType.NormalizeBaseType() = %v, want %v", got, tt.want)
			}
		})
	}
}
