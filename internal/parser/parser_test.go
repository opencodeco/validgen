package parser

import (
	"reflect"
	"testing"

	"github.com/opencodeco/validgen/internal/common"
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
							Type:      common.FieldType{BaseType: "string", ComposedType: "", Size: ""},
							Tag:       "valid:\"required\"",
						},
						{
							FieldName: "LastName",
							Type:      common.FieldType{BaseType: "string", ComposedType: "", Size: ""},
							Tag:       "valid:\"required\"",
						},
						{
							FieldName: "Age",
							Type:      common.FieldType{BaseType: "uint8", ComposedType: "", Size: ""},
							Tag:       "valid:\"gte=18,lte=130\"",
						},
						{
							FieldName: "UserName",
							Type:      common.FieldType{BaseType: "string", ComposedType: "", Size: ""},
							Tag:       "valid:\"min=5,max=10\"",
						},
						{
							FieldName: "Optional",
							Type:      common.FieldType{BaseType: "string", ComposedType: "", Size: ""},
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
							Type:      common.FieldType{BaseType: "string", ComposedType: "", Size: ""},
							Tag:       "valid:\"required\"",
						},
						{
							FieldName: "LastName",
							Type:      common.FieldType{BaseType: "string", ComposedType: "", Size: ""},
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
							Type:      common.FieldType{BaseType: "string", ComposedType: "", Size: ""},
							Tag:       "",
						},
						{
							FieldName: "Field2",
							Type:      common.FieldType{BaseType: "uint8", ComposedType: "", Size: ""},
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
							Type:      common.FieldType{BaseType: "string", ComposedType: "", Size: ""},
							Tag:       "valid:\"required\"",
						},
						{
							FieldName: "LastName",
							Type:      common.FieldType{BaseType: "string", ComposedType: "", Size: ""},
							Tag:       "valid:\"required\"",
						},
						{
							FieldName: "Address",
							Type:      common.FieldType{BaseType: "inpkg.TypeAddress", ComposedType: "", Size: ""},
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
							Type:      common.FieldType{BaseType: "string", ComposedType: "", Size: ""},
							Tag:       "valid:\"required\"",
						},
						{
							FieldName: "LastName",
							Type:      common.FieldType{BaseType: "string", ComposedType: "", Size: ""},
							Tag:       "valid:\"required\"",
						},
						{
							FieldName: "Address",
							Type:      common.FieldType{BaseType: "main.TypeAddress", ComposedType: "", Size: ""},
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
							Type:      common.FieldType{BaseType: "string", ComposedType: "", Size: ""},
							Tag:       "valid:\"required\"",
						},
						{
							FieldName: "City",
							Type:      common.FieldType{BaseType: "string", ComposedType: "", Size: ""},
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
							Type:      common.FieldType{BaseType: "string", ComposedType: "", Size: ""},
							Tag:       "valid:\"required\"",
						},
						{
							FieldName: "Types",
							Type:      common.FieldType{BaseType: "string", ComposedType: "[]", Size: ""},
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
							Type:      common.FieldType{BaseType: "string", ComposedType: "", Size: ""},
							Tag:       "valid:\"required\"",
						},
						{
							FieldName: "Types",
							Type:      common.FieldType{BaseType: "string", ComposedType: "[N]", Size: "5"},
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
							Type:      common.FieldType{BaseType: "bool", ComposedType: "", Size: ""},
							Tag:       "valid:\"eq=true\"",
						},
						{
							FieldName: "IsNotCorrect",
							Type:      common.FieldType{BaseType: "bool", ComposedType: "", Size: ""},
							Tag:       "valid:\"eq=false\"",
						},
					},
					Imports: map[string]Import{},
				},
			},
		},

		{
			name: "Map type",
			args: args{
				fullpath: "example/main.go",
				src: "package main\n" +
					"type AllTypes struct {\n" +
					"	MapField1    map[string]string `valid:\"required\"`\n" +
					"	MapField2    map[string]uint8 `valid:\"len=3\"`\n" +
					"	MapField3    map[uint8]string `valid:\"max=5\"`\n" +
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
							FieldName: "MapField1",
							Type:      common.FieldType{BaseType: "string", ComposedType: "map", Size: ""},
							Tag:       "valid:\"required\"",
						},
						{
							FieldName: "MapField2",
							Type:      common.FieldType{BaseType: "string", ComposedType: "map", Size: ""},
							Tag:       "valid:\"len=3\"",
						},
						{
							FieldName: "MapField3",
							Type:      common.FieldType{BaseType: "uint8", ComposedType: "map", Size: ""},
							Tag:       "valid:\"max=5\"",
						},
					},
					Imports: map[string]Import{},
				},
			},
		},

		{
			name: "Float32 type",
			args: args{
				fullpath: "example/main.go",
				src: "package main\n" +
					"type AllTypes struct {\n" +
					"	IsEqual    float32 `valid:\"eq=1.2\"`\n" +
					"	IsNotEqual float32 `valid:\"eq=3.4\"`\n" +
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
							FieldName: "IsEqual",
							Type:      common.FieldType{BaseType: "float32", ComposedType: "", Size: ""},
							Tag:       "valid:\"eq=1.2\"",
						},
						{
							FieldName: "IsNotEqual",
							Type:      common.FieldType{BaseType: "float32", ComposedType: "", Size: ""},
							Tag:       "valid:\"eq=3.4\"",
						},
					},
					Imports: map[string]Import{},
				},
			},
		},

		{
			name: "Slice type",
			args: args{
				fullpath: "example/main.go",
				src: "package main\n" +
					"type AllTypes struct {\n" +
					"	SliceField1 []string `valid:\"required\"`\n" +
					"	SliceField2 []uint8 `valid:\"len=3\"`\n" +
					"	SliceField3 []string `valid:\"max=5\"`\n" +
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
							FieldName: "SliceField1",
							Type:      common.FieldType{BaseType: "string", ComposedType: "[]", Size: ""},
							Tag:       "valid:\"required\"",
						},
						{
							FieldName: "SliceField2",
							Type:      common.FieldType{BaseType: "uint8", ComposedType: "[]", Size: ""},
							Tag:       "valid:\"len=3\"",
						},
						{
							FieldName: "SliceField3",
							Type:      common.FieldType{BaseType: "string", ComposedType: "[]", Size: ""},
							Tag:       "valid:\"max=5\"",
						},
					},
					Imports: map[string]Import{},
				},
			},
		},

		{
			name: "Array type",
			args: args{
				fullpath: "example/main.go",
				src: "package main\n" +
					"type AllTypes struct {\n" +
					"	ArrayField1 [3]string `valid:\"in=1 2 3\"`\n" +
					"	ArrayField2 [3]uint8 `valid:\"nin= 4 5 6\"`\n" +
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
							FieldName: "ArrayField1",
							Type:      common.FieldType{BaseType: "string", ComposedType: "[N]", Size: "3"},
							Tag:       "valid:\"in=1 2 3\"",
						},
						{
							FieldName: "ArrayField2",
							Type:      common.FieldType{BaseType: "uint8", ComposedType: "[N]", Size: "3"},
							Tag:       "valid:\"nin= 4 5 6\"",
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
			result += "  Field: " + f.FieldName + " Type: " + f.Type.ToString() + " Tag: " + f.Tag + "\n"
		}
		for _, v := range s.Imports {
			result += "  Import: " + v.Name + " Path: " + v.Path + "\n"
		}
	}

	return result
}
