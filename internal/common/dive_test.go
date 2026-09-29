package common

import "testing"

func TestDiveInto(t *testing.T) {
	tests := []struct {
		name    string
		in      FieldType
		want    FieldType
		wantErr string
	}{
		{
			name: "slice of structs",
			in:   FieldType{BaseType: "main.Address", ComposedType: "[]"},
			want: FieldType{BaseType: "main.Address"},
		},
		{
			name: "slice of pointers",
			in:   FieldType{BaseType: "main.Address", ComposedType: "*[]", ElemPointer: true},
			want: FieldType{BaseType: "main.Address", ComposedType: "*"},
		},
		{
			name: "pointer to slice",
			in:   FieldType{BaseType: "main.Address", ComposedType: "*[]"},
			want: FieldType{BaseType: "main.Address"},
		},
		{
			name: "map value",
			in: FieldType{
				BaseType:     "string",
				ComposedType: "map",
				MapValue:     &FieldType{BaseType: "string"},
			},
			want: FieldType{BaseType: "string"},
		},
		{
			name: "map of structs",
			in: FieldType{
				BaseType:     "string",
				ComposedType: "map",
				MapValue:     &FieldType{BaseType: "main.Address"},
			},
			want: FieldType{BaseType: "main.Address"},
		},
		{
			name: "nested slice",
			in:   FieldType{BaseType: "string", ComposedType: "[][]"},
			want: FieldType{BaseType: "string", ComposedType: "[]"},
		},
		{
			name:    "scalar",
			in:      FieldType{BaseType: "string"},
			wantErr: "string is not a slice, array, or map",
		},
		{
			name:    "nested array",
			in:      FieldType{BaseType: "int", ComposedType: "[N][N]", Size: "2"},
			wantErr: "nested arrays are not supported",
		},
		{
			name:    "pointer mixed into the next container",
			in:      FieldType{BaseType: "string", ComposedType: "*[][]", ElemPointer: true},
			wantErr: "unsupported pointer composition *[]string",
		},
		{
			name:    "map without a value type",
			in:      FieldType{BaseType: "string", ComposedType: "map"},
			wantErr: "map value type is not available",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.in.DiveInto()
			if tt.wantErr != "" {
				if err == nil || err.Error() != tt.wantErr {
					t.Fatalf("DiveInto() error = %v, want %s", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("DiveInto() error = %v", err)
			}
			if got.BaseType != tt.want.BaseType || got.ComposedType != tt.want.ComposedType || got.ElemPointer != tt.want.ElemPointer {
				t.Fatalf("DiveInto() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
