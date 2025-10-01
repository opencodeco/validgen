package common

import (
	"reflect"
	"testing"
)

func TestFieldType_ToString(t *testing.T) {
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
			if got := ft.ToString(); got != tt.want {
				t.Errorf("FieldType.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFieldType_IsGoType(t *testing.T) {
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

func TestFromNormalizedToBasicTypes(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "string type",
			args: args{t: "<STRING>"},
			want: []string{"string"},
		},
		{
			name: "bool type",
			args: args{t: "<BOOL>"},
			want: []string{"bool"},
		},
		{
			name: "int type",
			args: args{t: "<INT>"},
			want: []string{"int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64"},
		},
		{
			name: "float type",
			args: args{t: "<FLOAT>"},
			want: []string{"float32", "float64"},
		},
		{
			name: "map string type",
			args: args{t: "map[<STRING>]"},
			want: []string{"map[string]"},
		},
		{
			name: "map bool type",
			args: args{t: "map[<BOOL>]"},
			want: []string{"map[bool]"},
		},
		{
			name: "map int type",
			args: args{t: "map[<INT>]"},
			want: []string{"map[int]", "map[int8]", "map[int16]", "map[int32]", "map[int64]", "map[uint]", "map[uint8]", "map[uint16]", "map[uint32]", "map[uint64]"},
		},
		{
			name: "map float type",
			args: args{t: "map[<FLOAT>]"},
			want: []string{"map[float32]", "map[float64]"},
		},
		{
			name: "slice string type",
			args: args{t: "[]<STRING>"},
			want: []string{"[]string"},
		},
		{
			name: "slice bool type",
			args: args{t: "[]<BOOL>"},
			want: []string{"[]bool"},
		},
		{
			name: "slice int type",
			args: args{t: "[]<INT>"},
			want: []string{"[]int", "[]int8", "[]int16", "[]int32", "[]int64", "[]uint", "[]uint8", "[]uint16", "[]uint32", "[]uint64"},
		},
		{
			name: "slice float type",
			args: args{t: "[]<FLOAT>"},
			want: []string{"[]float32", "[]float64"},
		},
		{
			name: "array string type",
			args: args{t: "[3]<STRING>"},
			want: []string{"[3]string"},
		},
		{
			name: "array bool type",
			args: args{t: "[3]<BOOL>"},
			want: []string{"[3]bool"},
		},
		{
			name: "array int type",
			args: args{t: "[3]<INT>"},
			want: []string{"[3]int", "[3]int8", "[3]int16", "[3]int32", "[3]int64", "[3]uint", "[3]uint8", "[3]uint16", "[3]uint32", "[3]uint64"},
		},
		{
			name: "array float type",
			args: args{t: "[3]<FLOAT>"},
			want: []string{"[3]float32", "[3]float64"},
		},
		{
			name: "invalid type",
			args: args{t: "<INVALID>"},
			want: []string{"invalid"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromNormalizedToBasicTypes(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromNormalizedToBasicTypes() = %v, want %v", got, tt.want)
			}
		})
	}
}
