package codegenerator

import "github.com/opencodeco/validgen/internal/analyzer"

type genValidations struct {
	StructsWithValidation map[string]struct{}
	Struct                *analyzer.Struct
}

func GenerateCode(structs []*analyzer.Struct) error {
	structsWithValidation := map[string]struct{}{}

	for _, st := range structs {
		structsWithValidation[st.PackageName+"."+st.StructName] = struct{}{}
	}

	for _, st := range structs {
		if !st.HasValidTag {
			continue
		}

		codeInfo := &genValidations{
			StructsWithValidation: structsWithValidation,
			Struct:                st,
		}

		if err := codeInfo.GenerateFileValidator(); err != nil {
			return err
		}
	}

	return nil
}
