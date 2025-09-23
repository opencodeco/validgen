package parser

import (
	"reflect"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func TestParseStructsOk(t *testing.T) {
	type args struct {
		fullpath string
		src      string
	}
	tests := []struct {
		name string
		args args
		want []*Struct
	}{
		{
			name: "No structs definition",
			args: args{
				fullpath: "example/main.go",
				src: `package main
		func main() {
		}
		`,
			},
			want: []*Struct{},
		},

		{
			name: "One struct definition",
			args: args{
				fullpath: "example/main.go",
				src: "package main\n" +
					"type AllTypes struct {\n" +
					"	FirstName string `valid:\"required\"`\n" +
					"	LastName  string `valid:\"required\"`\n" +
					"	Age       uint8  `valid:\"gte=18,lte=130\"`\n" +
					"	UserName  string `valid:\"min=5,max=10\"`\n" +
					"	Optional  string\n" +
					"}\n" +
					"func main() {\n" +
					"}\n",
			},
			want: []*Struct{
				{
					StructName:  "AllTypes",
					Path:        "./example",
					PackageName: "main",
					Fields: []Field{
						{
							FieldName: "FirstName",
							Type:      "string",
							Tag:       "valid:\"required\"",
						},
						{
							FieldName: "LastName",
							Type:      "string",
							Tag:       "valid:\"required\"",
						},
						{
							FieldName: "Age",
							Type:      "uint8",
							Tag:       "valid:\"gte=18,lte=130\"",
						},
						{
							FieldName: "UserName",
							Type:      "string",
							Tag:       "valid:\"min=5,max=10\"",
						},
						{
							FieldName: "Optional",
							Type:      "string",
							Tag:       "",
						},
					},
					Imports: map[string]Import{},
				},
			},
		},

		{
			name: "One struct with validations and one struct without validations",
			args: args{
				fullpath: "example/main.go",
				src: "package main\n" +
					"type AllTypes struct {\n" +
					"	FirstName string `valid:\"required\"`\n" +
					"	LastName  string `valid:\"required\"`\n" +
					"}\n" +
					"\n" +
					"type NoValidations struct {\n" +
					"	Field1 string\n" +
					"	Field2 uint8\n" +
					"}\n" +
					"func main() {\n" +
					"}\n",
			},
			want: []*Struct{
				{
					StructName:  "AllTypes",
					Path:        "./example",
					PackageName: "main",
					Fields: []Field{
						{
							FieldName: "FirstName",
							Type:      "string",
							Tag:       "valid:\"required\"",
						},
						{
							FieldName: "LastName",
							Type:      "string",
							Tag:       "valid:\"required\"",
						},
					},
					Imports: map[string]Import{},
				},
				{
					StructName:  "NoValidations",
					Path:        "./example",
					PackageName: "main",
					Fields: []Field{
						{
							FieldName: "Field1",
							Type:      "string",
							Tag:       "",
						},
						{
							FieldName: "Field2",
							Type:      "uint8",
							Tag:       "",
						},
					},
					Imports: map[string]Import{},
				},
			},
		},

		{
			name: "Struct definition in pkg",
			args: args{
				fullpath: "example/main.go",
				src: "package main\n" +
					"import \"github.com/opencodeco/validgen/tests/endtoend/inpkg\"\n" +
					"type User struct {\n" +
					"	FirstName string `valid:\"required\"`\n" +
					"	LastName  string `valid:\"required\"`\n" +
					"	Address inpkg.TypeAddress `valid:\"required\"`\n" +
					"}\n" +
					"\n" +
					"func main() {\n" +
					"}\n",
			},
			want: []*Struct{
				{
					StructName:  "User",
					Path:        "./example",
					PackageName: "main",
					Fields: []Field{
						{
							FieldName: "FirstName",
							Type:      "string",
							Tag:       "valid:\"required\"",
						},
						{
							FieldName: "LastName",
							Type:      "string",
							Tag:       "valid:\"required\"",
						},
						{
							FieldName: "Address",
							Type:      "inpkg.TypeAddress",
							Tag:       "valid:\"required\"",
						},
					},
					Imports: map[string]Import{
						"inpkg": {
							Name: "inpkg",
							Path: "github.com/opencodeco/validgen/tests/endtoend/inpkg",
						},
					},
				},
			},
		},

		{
			name: "Nested structs",
			args: args{
				fullpath: "example/main.go",
				src: "package main\n" +
					"type User struct {\n" +
					"	FirstName string `valid:\"required\"`\n" +
					"	LastName  string `valid:\"required\"`\n" +
					"	Address   TypeAddress `valid:\"required\"`\n" +
					"}\n" +
					"type TypeAddress struct {\n" +
					"	Street string `valid:\"required\"`\n" +
					"	City string `valid:\"required\"`\n" +
					"}\n" +
					"\n" +
					"func main() {\n" +
					"}\n",
			},
			want: []*Struct{
				{
					StructName:  "User",
					Path:        "./example",
					PackageName: "main",
					Fields: []Field{
						{
							FieldName: "FirstName",
							Type:      "string",
							Tag:       "valid:\"required\"",
						},
						{
							FieldName: "LastName",
							Type:      "string",
							Tag:       "valid:\"required\"",
						},
						{
							FieldName: "Address",
							Type:      "main.TypeAddress",
							Tag:       "valid:\"required\"",
						},
					},
					Imports: map[string]Import{},
				},
				{
					StructName:  "TypeAddress",
					Path:        "./example",
					PackageName: "main",
					Fields: []Field{
						{
							FieldName: "Street",
							Type:      "string",
							Tag:       "valid:\"required\"",
						},
						{
							FieldName: "City",
							Type:      "string",
							Tag:       "valid:\"required\"",
						},
					},
					Imports: map[string]Import{},
				},
			},
		},

		{
			name: "Slice of strings",
			args: args{
				fullpath: "example/main.go",
				src: "package main\n" +
					"type AllTypes struct {\n" +
					"	FirstName string `valid:\"required\"`\n" +
					"	Types   []string `valid:\"required\"`\n" +
					"}\n" +

					"func main() {\n" +
					"}\n",
			},
			want: []*Struct{
				{
					StructName:  "AllTypes",
					Path:        "./example",
					PackageName: "main",
					Fields: []Field{
						{
							FieldName: "FirstName",
							Type:      "string",
							Tag:       "valid:\"required\"",
						},
						{
							FieldName: "Types",
							Type:      "[]string",
							Tag:       "valid:\"required\"",
						},
					},
					Imports: map[string]Import{},
				},
			},
		},

		{
			name: "Array of strings",
			args: args{
				fullpath: "example/main.go",
				src: "package main\n" +
					"type AllTypes struct {\n" +
					"	FirstName string `valid:\"required\"`\n" +
					"	Types   [5]string `valid:\"in=a b c\"`\n" +
					"}\n" +

					"func main() {\n" +
					"}\n",
			},
			want: []*Struct{
				{
					StructName:  "AllTypes",
					Path:        "./example",
					PackageName: "main",
					Fields: []Field{
						{
							FieldName: "FirstName",
							Type:      "string",
							Tag:       "valid:\"required\"",
						},
						{
							FieldName: "Types",
							Type:      "[N]string",
							Tag:       "valid:\"in=a b c\"",
						},
					},
					Imports: map[string]Import{},
				},
			},
		},

		{
			name: "Boolean type",
			args: args{
				fullpath: "example/main.go",
				src: "package main\n" +
					"type AllTypes struct {\n" +
					"	IsCorrect    bool `valid:\"eq=true\"`\n" +
					"	IsNotCorrect bool `valid:\"eq=false\"`\n" +
					"}\n" +

					"func main() {\n" +
					"}\n",
			},
			want: []*Struct{
				{
					StructName:  "AllTypes",
					Path:        "./example",
					PackageName: "main",
					Fields: []Field{
						{
							FieldName: "IsCorrect",
							Type:      "bool",
							Tag:       "valid:\"eq=true\"",
						},
						{
							FieldName: "IsNotCorrect",
							Type:      "bool",
							Tag:       "valid:\"eq=false\"",
						},
					},
					Imports: map[string]Import{},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseStructs(tt.args.fullpath, tt.args.src)
			var wantErr error = nil
			if err != wantErr {
				t.Errorf("parseStructs() error = %v, wantErr %v", err, wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				gotStr := structsToString(got)
				wantStr := structsToString(tt.want)
				dmp := diffmatchpatch.New()
				diffs := dmp.DiffMain(gotStr, wantStr, false)
				if len(diffs) > 1 {
					t.Errorf("parseStructs() diff = \n%v", dmp.DiffPrettyText(diffs))
				}
			}
		})
	}
}

func structsToString(structs []*Struct) string {
	var result string
	for _, s := range structs {
		result += "Struct: " + s.StructName + "\n"
		result += "Path: " + s.Path + "\n"
		result += "PackageName: " + s.PackageName + "\n"
		for _, f := range s.Fields {
			result += "  Field: " + f.FieldName + " Type: " + f.Type + " Tag: " + f.Tag + "\n"
		}
		for _, v := range s.Imports {
			result += "  Import: " + v.Name + " Path: " + v.Path + "\n"
		}
	}

	return result
}
