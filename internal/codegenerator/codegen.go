package codegenerator

import (
	"github.com/opencodeco/validgen/internal/analyzer"
	"github.com/opencodeco/validgen/internal/common"
	"github.com/opencodeco/validgen/internal/parser"
)

type genValidations struct {
	StructsWithValidation map[string]struct{}
	Struct                *analyzer.Struct
}

func GenerateCode(structs []*analyzer.Struct) (map[string]*Pkg, error) {
	structsWithValidation := map[string]struct{}{}
	usedPkgs := map[string]struct{}{}

	for _, st := range structs {
		structsWithValidation[common.KeyPath(st.PackageName, st.StructName)] = struct{}{}
		usedPkgs[st.PackageName] = struct{}{}
	}

	pkgs := make(map[string]*Pkg)
	for _, st := range structs {
		if !st.HasValidTag {
			continue
		}

		codeInfo := &genValidations{
			StructsWithValidation: structsWithValidation,
			Struct:                st,
		}

		funcCode, err := codeInfo.BuildFuncValidatorCode()
		if err != nil {
			return nil, err
		}

		pkdId := common.KeyPath(st.Path, st.PackageName)
		pkg, ok := pkgs[pkdId]
		if !ok {
			pkg = &Pkg{
				Name:    st.PackageName,
				Path:    st.Path,
				Imports: map[string]parser.Import{},
				Structs: map[string]*Struct{},
			}
			pkgs[pkdId] = pkg
		}

		cgSt := &Struct{
			Struct:            st,
			ValidatorFuncCode: funcCode,
		}

		pkg.Structs[st.StructName] = cgSt

		for _, imp := range st.Imports {
			if _, ok := usedPkgs[imp.Name]; ok {
				pkg.Imports[imp.Name] = imp
			}
		}
	}

	return pkgs, nil
}
