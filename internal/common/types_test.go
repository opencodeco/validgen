package common

import "testing"

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
