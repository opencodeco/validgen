package codegenerator

import "github.com/opencodeco/validgen/internal/parser"

func GenerateCode(structs []*parser.Struct) error {
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
