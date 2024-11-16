package main

import (
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func TestFileHeaderGenerate(t *testing.T) {
	type fields struct {
		PackageName string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "Valid import path",
			fields: fields{
				PackageName: "main",
			},
			want: `package main

import (
	"fmt"
)
`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &FileHeader{
				PackageName: tt.fields.PackageName,
			}
			got, err := p.Generate()
			if (err != nil) != tt.wantErr {
				t.Errorf("FileHeader.Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FileHeader.Generate() = %v, want %v", got, tt.want)
				dmp := diffmatchpatch.New()
				diffs := dmp.DiffMain(tt.want, got, false)
				if len(diffs) > 1 {
					t.Errorf("FileHeader.Generate() diff = \n%v", dmp.DiffPrettyText(diffs))
				}
			}
		})
	}
}
