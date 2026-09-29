package codegenerator

import (
	"testing"

	"github.com/opencodeco/validgen/internal/analyzer"
	"github.com/opencodeco/validgen/internal/common"
	"github.com/opencodeco/validgen/internal/parser"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func TestBuildUnmarshalJSONCode(t *testing.T) {
	gv := GenValidations{
		Struct: &analyzer.Struct{
			Struct: parser.Struct{
				StructName: "User",
			},
		},
	}

	want := `func (obj *User) UnmarshalJSON(b []byte) error {
	type alias User
	if err := json.Unmarshal(b, (*alias)(obj)); err != nil {
		return err
	}
	if errs := UserValidate(obj); len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}
`
	got := gv.BuildUnmarshalJSONCode()
	if got != want {
		t.Errorf("BuildUnmarshalJSONCode() mismatch")
		dmp := diffmatchpatch.New()
		t.Errorf("diff =\n%s", dmp.DiffPrettyText(dmp.DiffMain(want, got, false)))
	}
}

func TestGenerateCodeUnmarshalJSONOptIn(t *testing.T) {
	structs := []*analyzer.Struct{
		{
			Struct: parser.Struct{
				StructName:  "User",
				PackageName: "main",
				Path:        "/tmp/user",
				Fields: []parser.Field{
					{
						FieldName: "FirstName",
						Type:      common.FieldType{BaseType: "string"},
						Tag:       `valid:"required"`,
					},
				},
			},
			HasValidTag: true,
			FieldsValidations: []analyzer.FieldValidations{
				{Validations: []*analyzer.Validation{AssertParserValidation(t, "required")}},
			},
		},
	}

	pkgs, err := GenerateCode(structs, Options{})
	if err != nil {
		t.Fatalf("GenerateCode() error = %v", err)
	}
	pkg := pkgs[common.KeyPath("/tmp/user", "main")]
	if pkg == nil {
		t.Fatal("expected package")
	}
	if pkg.Structs["User"].UnmarshalJSONCode != "" {
		t.Fatalf("expected no UnmarshalJSON without opt-in, got %q", pkg.Structs["User"].UnmarshalJSONCode)
	}

	pkgs, err = GenerateCode(structs, Options{UnmarshalJSON: true})
	if err != nil {
		t.Fatalf("GenerateCode() error = %v", err)
	}
	pkg = pkgs[common.KeyPath("/tmp/user", "main")]
	if pkg.Structs["User"].UnmarshalJSONCode == "" {
		t.Fatal("expected UnmarshalJSON with opt-in")
	}
	if got := pkg.Structs["User"].UnmarshalJSONCode; got != (&GenValidations{Struct: structs[0]}).BuildUnmarshalJSONCode() {
		t.Fatalf("unexpected UnmarshalJSON code:\n%s", got)
	}
}
