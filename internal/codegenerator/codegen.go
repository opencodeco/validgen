package codegenerator

import "github.com/opencodeco/validgen/internal/analyzer"

var structsWithValidation map[string]struct{} = map[string]struct{}{}

func GenerateCode(structs []*analyzer.Struct) error {
	for _, st := range structs {
		if st.HasValidTag {
			structsWithValidation[st.StructName] = struct{}{}
		}
	}

	for _, st := range structs {
		if !st.HasValidTag {
			continue
		}

		if err := GenerateFileValidator(st); err != nil {
			return err
		}
	}

	return nil
}
