# TestGen

TestGen is a tool responsible for generating comprehensive test suites for ValidGen validators and supported types.
These generated tests include unit tests, end-to-end tests, and benchmark tests.
The benchmark tests compare ValidGen and GoValidator performance.

## Why a new tool?

First, let's address the question: why create a dedicated tool for test generation?

ValidGen currently supports 21 validations across multiple data types:

| Validation      | Basic types           | Slice                 | Array                 | Map                   |
| -               | -                     | -                     | -                     | -                     |
| eq              | STRING INT FLOAT COMPLEX BOOL |                       |                       |                       |
| required        | STRING INT FLOAT COMPLEX BOOL | STRING INT FLOAT BOOL |                       | STRING INT FLOAT BOOL |
| gt              | INT FLOAT             |                       |                       |                       |
| gte             | INT FLOAT             |                       |                       |                       |
| lte             | INT FLOAT             |                       |                       |                       |
| lt              | INT FLOAT             |                       |                       |                       |
| min             | STRING                | STRING INT FLOAT BOOL |                       | STRING INT FLOAT BOOL |
| max             | STRING                | STRING INT FLOAT BOOL |                       | STRING INT FLOAT BOOL | 
| eq_ignore_case  | STRING                |                       |                       |                       |
| len             | STRING                | STRING INT FLOAT BOOL |                       | STRING INT FLOAT BOOL |
| neq             | STRING INT FLOAT COMPLEX BOOL |                       |                       |                       |
| neq_ignore_case | STRING                |                       |                       |                       |
| in              | STRING INT FLOAT COMPLEX BOOL | STRING INT FLOAT BOOL | STRING INT FLOAT BOOL | STRING INT FLOAT BOOL |
| nin             | STRING INT FLOAT COMPLEX BOOL | STRING INT FLOAT BOOL | STRING INT FLOAT BOOL | STRING INT FLOAT BOOL |
| email           | STRING                |                       |                       |                       |
| eqfield         | STRING INT FLOAT COMPLEX BOOL |                       |                       |                       |
| neqfield        | STRING INT FLOAT COMPLEX BOOL |                       |                       |                       |
| gtefield        | INT FLOAT             |                       |                       |                       |
| gtfield         | INT FLOAT             |                       |                       |                       |
| ltefield        | INT FLOAT             |                       |                       |                       |
| ltfield         | INT FLOAT             |                       |                       |                       |

In this table:
- **STRING** represents the `string` Go type
- **BOOL** represents the `bool` Go type
- **INT** represents all ten integer Go types: `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`
- **FLOAT** represents both float Go types: `float32`, `float64`
- **COMPLEX** represents both complex Go types: `complex64`, `complex128`. Scalar `eq`, `neq`, `in`, `nin`, and `required` are generated for non-pointer values. Ordering tags are rejected. Field equality stays hand-written.

For slices, arrays, and maps, the same type expansion applies. For example, slice STRING is `[]string`, while slice INT expands to all integer Go types.

Additionally, the pointer modifier (`*`) is significant because ValidGen generates different code for pointer and non-pointer types.

For each possible combination (validation × type), the following tests are required:
- Unit test for the analyzer phase
- Unit test for the code generator phase
- Benchmark test comparing ValidGen and GoValidator
- End-to-end test

Since we need to test both valid and invalid scenarios, all validations must be tested against all types.
Currently, this means:
- Go types: 14
- Validations: 21
- Test types: 4
- Pointer variants: 2 (with/without pointer)
- Input scenarios: 2 (valid/invalid)

**14 × 21 × 4 × 2 × 2 = 4,704 distinct test cases**

Creating and maintaining these thousands of tests manually is tedious, error-prone, and impractical. Early attempts to write these tests by hand proved difficult to maintain when adding new operations and types.

Previously, ValidGen had two separate test generators:
- Benchmark tests comparing ValidGen and GoValidator
- End-to-end tests for:
    - Integer type validations
    - Float type validations
    - All possible use cases (all validations × all types)

However, these generators lacked a common configuration, didn't implement all tests for all cases, and keeping the separate configuration files in sync was difficult.

## What TestGen generates

`typesValidation` in `validations.go` is the case list. `make testgen` walks that list and writes four suites. Each suite is a pointer file and a non-pointer file. The list contains non-field operations only.

- Benchmark tests between ValidGen and GoValidator, in `tests/cmpbenchtests/generated_cmp_perf_*`.
- End-to-end tests for each validation, type class, and valid or invalid input, in `tests/endtoend/generated_endtoend_*`.
- Unit tests for `BuildValidationCode`, in `internal/codegenerator/generated_validation_code_*`.
- Unit tests for the generated validator function, in `internal/codegenerator/generated_function_code_*`.

A case with both inputs emits the valid input and the invalid input. Array `required` cases set `excludeIf` to `noPointer` because a non-pointer Go array cannot be empty. Scalar complex cases set `skipPointer` because pointer complex values are outside this generator.

`go test ./testgen` checks that list. `TestTypesValidationListsNonFieldOperations` requires one entry for each operation below, with `isFieldValidation` false, the same argument count as `operations.New()`, and a case for every type `IsValidByType` accepts. `TestTypesValidationCasesBuildValidationCode` calls `BuildValidationCode` for each concrete type, including the pointer form, and checks that the generated error text contains the catalog message. Cases marked `skipPointer` are checked only as non-pointers.

| Operation | Generated |
| - | - |
| email, required, eq, neq | yes |
| gt, gte, lt, lte | yes |
| min, max, len | yes |
| eq_ignore_case, neq_ignore_case | yes |
| in, nin | yes |
| eqfield, neqfield, gtfield, gtefield, ltfield, ltefield | hand-written |

## Suites that stay hand-written

These groups already have tests beside the code they check. A generator that reads the same list would only compare that list with itself.

- Operation checks in `internal/analyzer/operations/operations_test.go`. The functions are `TestOperationsIsValid`, `TestOperationsIsValidByType`, `TestOperationsIsFieldOperation`, and `TestOperationsArgsCount`.
- Condition-table checks in `internal/codegenerator/get_test_elements_*_test.go`. Each case stores the condition string and the error string passed to `DefineTestElements`.
- Parser checks in `internal/parser/parser_test.go`. They compare parsed structs with source text.
- Examples under `_examples/`.

Field-operation rows in the four generated suites wait on integer field operations in issue #78. Complex ordering tags stay rejected, and `dive` in issue #7 is separate work.

## Usage

To generate the tests:

```bash
# Navigate to the project root folder
cd validgen

# Run testgen
make testgen
```
