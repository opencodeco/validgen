# ValidGen internals

`validgen [-unmarshal-json] <path>` reads Go files under that path and writes `validator__.go` beside structs that carry a `valid` tag. The command is `main` in the repository root. This note follows that path through the current packages. The user-facing validation matrix stays in the [README](../README.md).

`main` parses flags with `flag.ExitOnError`. `-unmarshal-json` defaults to false. The command then requires one path argument. Any other argument count prints usage and exits with status 1. An error from the calls below ends in `log.Fatal`.

## Execution pipeline

`main` calls the stages in this order.

1. `parser.ExtractStructs` walks the path and returns `[]*parser.Struct`.
2. `analyzer.AnalyzeStructs` checks tags and returns `[]*analyzer.Struct`.
3. `analyzer.Struct.PrintInfo` prints every analyzed struct before any file is written.
4. `codegenerator.GenerateCode` takes the analyzed structs and a `codegenerator.Options` value. `Options.UnmarshalJSON` is the flag. It returns `map[string]*codegenerator.Pkg` with function source for structs that have a `valid` tag.
5. `pkgwriter.Writer` formats each package and writes `validator__.go`.

TestGen is a separate program under `testgen/`. It is described at the end.

## Packages

`internal/parser` loads `.go` files and records structs, fields, tags, and imports.

`internal/analyzer` reads `valid` tags, checks them against `internal/analyzer/operations`, and resolves field-to-field comparisons.

`internal/codegenerator` turns an analyzed struct into the text of a `Validate` function, and into an `UnmarshalJSON` method when `-unmarshal-json` is set. Operation expressions live in `condition_table.go`.

`internal/pkgwriter` wraps those functions in a file, runs `go/format`, and writes `validator__.go`.

`internal/common` holds the shared `FieldType`, the normalized type names, `KeyPath`, and `CountValues`.

`types` is the package generated validators import, `github.com/opencodeco/validgen/types`. It defines `ValidationError` and the helpers the condition table calls, including `IsValidEmail`, `SliceOnlyContains`, `SliceNotContains`, `MapOnlyContains`, and `MapNotContains`. Case-insensitive string checks call `strings.EqualFold` directly.

## Parser

`ExtractStructs` uses `filepath.WalkDir`. Every file whose name ends in `.go` is read, including files in subdirectories and `*_test.go` files. `go/parser.ParseFile` builds an AST with a new `token.FileSet` per file. `ast.Inspect` then handles three node kinds.

An `*ast.File` node supplies the package name and the import map. Each import is stored as `parser.Import` with the local name and the import path. If the import has no explicit name, the local name is the last path element.

An `*ast.TypeSpec` node starts a new `parser.Struct`. `StructName` is the type name. `PackageName` is the file package. `Path` is `./` plus the file's directory. `Imports` is the file import map.

An `*ast.StructType` node fills that struct and appends it. The field loop keeps names only, so an embedded field is skipped. A field with several names becomes one `parser.Field` per name. The stored tag is the full struct tag after `strconv.Unquote`.

`extractCompleteType` fills `common.FieldType`.

`BaseType` is the type name. A builtin Go name stays as written. Any other identifier is stored as `package.Name` through `common.KeyPath`. A selector such as `otherpkg.Address` is stored as `otherpkg.Address`.

`ComposedType` records shape during that walk.

A pointer appends `*` and then walks the inner expression. A slice walks the element and then appends `[]`. An array walks the element, then sets `Size` from the length literal and appends `[N]`. A map appends `map` and then walks `ast.MapType.Key`.

The map value expression is ignored. Parser tests record `map[string]uint8` with `BaseType` `string`, and `map[uint8]string` with `BaseType` `uint8`. Later stages classify the map from that key type. `FieldType.ToType` prints a map as `map[BaseType]BaseType`.

The same walk collapses a pointer element into the container marker. The pointer parser test records `[]*int64` as `BaseType` `int64` and `ComposedType` `*[]`, and `*map[string]bool` as `BaseType` `string` and `ComposedType` `*map`.

`FieldType.IsGoType` is true when `ComposedType` is `map` or `*map`. Otherwise it is true only when `BaseType` is `string`, `bool`, `float32`, `float64`, `complex64`, `complex128`, or one of the integer names listed in the analyzer section.

`extractCompleteType` returns an empty `FieldType` for an expression it does not recognize. Fields left with an empty `BaseType` are skipped.

## Analyzer

`AnalyzeStructs` runs three steps. The operations catalog is `operationsList` in `internal/analyzer/operations/operations_list.go`. Each entry has an argument count, a field-operation flag, and the normalized types it allows.

Normalized names come from `FieldType.ToNormalizedString`.

- `<STRING>` is `string`.
- `<BOOL>` is `bool`.
- `<INT>` is `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, and `uint64`.
- `<FLOAT>` is `float32` and `float64`.
- `<COMPLEX>` is `complex64` and `complex128`.

Slices, arrays, maps, and pointers keep a marker on that name, such as `[]<INT>`, `[N]<STRING>`, `map[<STRING>]`, and `*<FLOAT>`. `common.HelperFromNormalizedToFieldTypes` expands those names back to concrete `FieldType` values. TestGen and several generator tests call that helper. The CLI generator calls `ToNormalizedString` on the parsed field.

### Tag parsing

`analyzeFieldValidations` copies each `parser.Struct` into an `analyzer.Struct`. A field counts as validated only when its tag text starts with `valid:`. The analyzer sets `HasValidTag` if any field matches. It unquotes the text after `valid:` and splits on commas.

`ParserValidation` splits one validation on `=`. More than two pieces is an error. The operation's `CountValues` selects the shape.

`ZeroValue` operations `required` and `email` reject a target. `OneValue` operations require one target. `ManyValues` operations `in` and `nin` require a target list. A list that starts with a single quote is read as quoted strings. Any other list is split on commas and spaces.

The result is an `analyzer.Validation` with `Operation`, `ExpectedValues`, and `Values`.

An unknown operation has `CountValues` zero, which is `UndefinedValue`, and `ParserValidation` returns `unsupported validation`.

### Operation checks

`checkForInvalidOperations` rejects an operation that is not in the catalog.

If the field's `BaseType` is `package.Struct` for a struct parsed in this run, the type check stops there. Otherwise the field must be a Go type, and `IsValidByType` must accept the operation for `ToNormalizedString`.

`IsValidByType` strips one leading `*`. For `required` on a pointer it returns true without reading the type list. Every other operation must list the remaining normalized type. `GetConditionTable` still has to find a row for that normalized type when code is generated.

### Field comparisons

`eqfield`, `neqfield`, `gtfield`, `gtefield`, `ltfield`, and `ltefield` set `IsFieldOperation`. `eqfield` and `neqfield` allow `<STRING>`, `<INT>`, `<FLOAT>`, `<COMPLEX>`, and `<BOOL>`. Their condition-table rows compare with `==` and `!=`. `gtfield`, `gtefield`, `ltfield`, and `ltefield` allow `<INT>` and `<FLOAT>` only. `analyzeFieldOperations` checks field operations after the catalog checks.

The target is the single value from the tag. `Field2` refers to a field of the same struct. `Nested.Field2` refers to a field of the struct stored on `Nested`. The lookup key is `common.KeyPath`, which joins names with `.`. Both fields must already be in the parsed set, and their `FieldType` values must be equal. That comparison includes `ComposedType`, `BaseType`, and `Size`, so `int` and `int32` do not match.

## Code generator

`GenerateCode` indexes every parsed struct by `package.Struct` and every package name it saw. It then skips structs whose `HasValidTag` is false. Those structs produce no function. Their `package.Struct` keys stay in the index used for nested calls.

Structs that remain are grouped by `common.KeyPath(Path, PackageName)`. One group becomes one `codegenerator.Pkg` with the package name, the directory, a subset of imports, and a map of structs. Each struct stores the `Validate` function text in `ValidatorFuncCode`. When `Options.UnmarshalJSON` is true, `BuildUnmarshalJSONCode` also fills `UnmarshalJSONCode`.

`BuildUnmarshalJSONCode` returns this method for the struct name `Name`.

```go
func (obj *Name) UnmarshalJSON(b []byte) error {
	type alias Name
	if err := json.Unmarshal(b, (*alias)(obj)); err != nil {
		return err
	}
	if errs := NameValidate(obj); len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}
```

The alias is a distinct type, so `json.Unmarshal` does not call the generated method again. A decode error is returned as-is. A non-empty validation slice is returned from `errors.Join`. The README section Optional JSON unmarshaling shows the same method from the user side. `tests/jsonunmarshal/validator__.go` is a checked-in copy of that output.

`BuildFuncValidatorCode` fills this template.

```go
func {{.StructName}}Validate(obj *{{.StructName}}) []error {
var errs []error
{{range .Fields}}{{buildValidationCode .FieldName .Type .Validations}}{{end}}return errs
}
```

`StructToTpl` lines up `FieldsValidations[i]` with field `i`. `BuildValidationCode` walks that field's validations.

When `IsGoType` is true, `DefineTestElements` loads a row from `GetConditionTable` and `buildIfCode` emits this.

```go
if !(condition) {
errs = append(errs, types.NewValidationError("message"))
}
```

`GetConditionTable` selects the row whose `AcceptedTypes` contain `ToNormalizedString`. `DefineTestElements` substitutes placeholders in that row's `operation` string. `{{.Name}}` becomes the field name. `{{.Target}}` becomes one tag value.

Scalar `in` rows set `concatOperator` to `||`, and the per-value copies are joined. That includes `*<STRING>`, `*<INT>`, `*<FLOAT>`, and `*<BOOL>`. Scalar `nin` rows set `concatOperator` to `&&`. Slice, array, and map rows leave `concatOperator` empty, so `DefineTestElements` keeps one copy. That copy lists every target through `{{.TargetsAsStringSlice}}` or `{{.TargetsAsNumericSlice}}`.

Those slice and map copies call `types.SliceOnlyContains`, `types.SliceNotContains`, `types.MapOnlyContains`, or `types.MapNotContains`. Non-pointer array rows pass `obj.Field[:]` into the slice helpers. Literal string comparisons quote the target. `email` calls `types.IsValidEmail`. `eq_ignore_case` and `neq_ignore_case` call `strings.EqualFold`. Field comparisons compile to `obj.Field` compared with `obj.Other` or `obj.Nested.Field`.

When `IsGoType` is false, each validation on that field appends a nested call instead of a condition-table test. The call is `TypeValidate(&obj.Field)`, where `Type` is `BaseType`. If `BaseType` starts with the struct's own package name and a dot, that prefix is removed. A same-package field whose `BaseType` is `main.InnerStructType` calls `InnerStructTypeValidate`. A field whose `BaseType` is `mypkg.InnerStructType` calls `mypkg.InnerStructTypeValidate`. The call is emitted when `BaseType` is in the parsed-struct index. A missing type returns `no validator found for struct type`.

Imports kept on the generated package are the struct file's imports whose local name is a package name parsed in this run. `buildImportPath` writes each of those paths as a quoted import and always adds `github.com/opencodeco/validgen/types`. When any struct in the package has `UnmarshalJSON` source, it also adds `encoding/json` and `errors`. When generated code calls `strings.EqualFold`, it also adds `strings`.

## Package writer

`Writer` renders `fileValidatorTpl` for each package. The file starts with `// Code generated by ValidGen. DO NOT EDIT.` and then `//nolint:all`, declares the source package, and prints `ValidatorFuncCode` then `UnmarshalJSONCode` for each struct in the package map. `go/format` formats the buffer. The output path is `Path + "/validator__.go"`, and the file mode is `os.ModePerm`.

The same path is rewritten on every run. Structs in one package and directory share that file. Structs with no `valid` tag are absent from it. `_examples/test01` shows the shape. `User` has `valid` tags and `UserValidate` is written. `NoValidTag` is parsed, printed, and omitted from `validator__.go`.

## TestGen

TestGen is `package main` in `testgen/`. `make testgen` runs it and moves the files it writes.

`generate_tests.go` calls four generators. Each one executes a template under `testgen/`, formats the result with `go/format`, and writes a pair of pointer and non-pointer files.

- `generated_endtoend_no_pointer_tests.go` and `generated_endtoend_pointer_tests.go` move to `tests/endtoend/`.
- `generated_validation_code_no_pointer_test.go` and `generated_validation_code_pointer_test.go` move to `internal/codegenerator/`.
- `generated_function_code_no_pointer_test.go` and `generated_function_code_pointer_test.go` move to `internal/codegenerator/`.
- `generated_cmp_perf_no_pointer_test.go` and `generated_cmp_perf_pointer_test.go` move to `tests/cmpbenchtests/`.

The case list is `typesValidation` in `testgen/validations.go`. [testgen/README.md](../testgen/README.md) records which suites that list drives and which suites stay hand-written.

Hand-written tests cover the parser, the analyzer operation checks, and condition-table cases. `make unittests` runs `go test` on `./internal/...`, `./types/...`, and `./testgen/`. `make endtoendtests` builds `bin/validgen`, deletes existing `validator__.go` files under `tests/endtoend/`, runs the generator there, and executes `go run .` in that directory. It then does the same for `tests/jsonunmarshal/`, passing `-unmarshal-json`.
