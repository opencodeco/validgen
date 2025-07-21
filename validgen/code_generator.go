package validgen

func generateCode(structs []Struct) error {
	for _, structInfo := range structs {
		if !structInfo.HasValidTag {
			continue
		}

		if err := structInfo.GenerateFileValidator(); err != nil {
			return err
		}
	}

	return nil
}
