package validgen

func generateCode(structs []Struct) error {
	for _, structInfo := range structs {
		if !structInfo.HasVerifyTag {
			continue
		}

		if err := structInfo.GenerateFileValidator(); err != nil {
			return err
		}
	}

	return nil
}
