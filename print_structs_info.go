package main

func printStructsInfo(structs []StructInfo) {
	for _, structInfo := range structs {
		structInfo.PrintInfo()
	}
}
