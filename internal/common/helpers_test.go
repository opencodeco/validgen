package common

import (
	"reflect"
	"testing"
)

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
			name: "invalid type",
			args: args{t: "<INVALID>"},
			want: []string{"invalid"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HelperFromNormalizedToBasicTypes(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromNormalizedToBasicTypes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromNormalizedToFieldTypes(t *testing.T) {
	tests := []struct {
		name           string
		normalizedType string
		want           []FieldType
	}{
		{
			"string type",
			"<STRING>",
			[]FieldType{
				{BaseType: "string", ComposedType: "", Size: ""},
			},
		},
		{
			"bool type",
			"<BOOL>",
			[]FieldType{
				{BaseType: "bool", ComposedType: "", Size: ""},
			},
		},
		{
			"int type",
			"<INT>",
			[]FieldType{
				{BaseType: "int", ComposedType: "", Size: ""},
				{BaseType: "int8", ComposedType: "", Size: ""},
				{BaseType: "int16", ComposedType: "", Size: ""},
				{BaseType: "int32", ComposedType: "", Size: ""},
				{BaseType: "int64", ComposedType: "", Size: ""},
				{BaseType: "uint", ComposedType: "", Size: ""},
				{BaseType: "uint8", ComposedType: "", Size: ""},
				{BaseType: "uint16", ComposedType: "", Size: ""},
				{BaseType: "uint32", ComposedType: "", Size: ""},
				{BaseType: "uint64", ComposedType: "", Size: ""},
			},
		},
		{
			"map string type",
			"map[<STRING>]",
			[]FieldType{
				{BaseType: "string", ComposedType: "map", Size: ""},
			},
		},
		{
			"map bool type",
			"map[<BOOL>]",
			[]FieldType{
				{BaseType: "bool", ComposedType: "map", Size: ""},
			},
		},
		{
			"map int type",
			"map[<INT>]",
			[]FieldType{
				{BaseType: "int", ComposedType: "map", Size: ""},
				{BaseType: "int8", ComposedType: "map", Size: ""},
				{BaseType: "int16", ComposedType: "map", Size: ""},
				{BaseType: "int32", ComposedType: "map", Size: ""},
				{BaseType: "int64", ComposedType: "map", Size: ""},
				{BaseType: "uint", ComposedType: "map", Size: ""},
				{BaseType: "uint8", ComposedType: "map", Size: ""},
				{BaseType: "uint16", ComposedType: "map", Size: ""},
				{BaseType: "uint32", ComposedType: "map", Size: ""},
				{BaseType: "uint64", ComposedType: "map", Size: ""},
			},
		},
		{
			"slice string type",
			"[]<STRING>",
			[]FieldType{
				{BaseType: "string", ComposedType: "[]", Size: ""},
			},
		},
		{
			"slice bool type", "[]<BOOL>",
			[]FieldType{
				{BaseType: "bool", ComposedType: "[]", Size: ""},
			},
		},
		{
			"slice int type",
			"[]<INT>",
			[]FieldType{
				{BaseType: "int", ComposedType: "[]", Size: ""},
				{BaseType: "int8", ComposedType: "[]", Size: ""},
				{BaseType: "int16", ComposedType: "[]", Size: ""},
				{BaseType: "int32", ComposedType: "[]", Size: ""},
				{BaseType: "int64", ComposedType: "[]", Size: ""},
				{BaseType: "uint", ComposedType: "[]", Size: ""},
				{BaseType: "uint8", ComposedType: "[]", Size: ""},
				{BaseType: "uint16", ComposedType: "[]", Size: ""},
				{BaseType: "uint32", ComposedType: "[]", Size: ""},
				{BaseType: "uint64", ComposedType: "[]", Size: ""},
			},
		},
		{
			"array string type",
			"[3]<STRING>",
			[]FieldType{
				{BaseType: "string", ComposedType: "[N]", Size: "3"},
			},
		},
		{
			"array bool type",
			"[3]<BOOL>",
			[]FieldType{
				{BaseType: "bool", ComposedType: "[N]", Size: "3"},
			},
		},
		{
			"array int type",
			"[3]<INT>",
			[]FieldType{
				{BaseType: "int", ComposedType: "[N]", Size: "3"},
				{BaseType: "int8", ComposedType: "[N]", Size: "3"},
				{BaseType: "int16", ComposedType: "[N]", Size: "3"},
				{BaseType: "int32", ComposedType: "[N]", Size: "3"},
				{BaseType: "int64", ComposedType: "[N]", Size: "3"},
				{BaseType: "uint", ComposedType: "[N]", Size: "3"},
				{BaseType: "uint8", ComposedType: "[N]", Size: "3"},
				{BaseType: "uint16", ComposedType: "[N]", Size: "3"},
				{BaseType: "uint32", ComposedType: "[N]", Size: "3"},
				{BaseType: "uint64", ComposedType: "[N]", Size: "3"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HelperFromNormalizedToFieldTypes(tt.normalizedType)
			if err != nil {
				t.Errorf("FromNormalizedToFieldTypes() error = %v, wantErr %v", err, nil)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromNormalizedToFieldTypes() = %v, want %v", got, tt.want)
			}
		})
	}
}
