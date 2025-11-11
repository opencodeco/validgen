# TestGen

TestGen is a tool responsible for generating comprehensive test suites for ValidGen validators and supported types.
These generated tests include unit tests, end-to-end tests, and benchmark tests.
The benchmark tests compare ValidGen and GoValidator performance.

## Why a new tool?

First, let's address the question: why create a dedicated tool for test generation?

ValidGen currently supports 21 validations across multiple data types:

| Validation      | Basic types           | Slice                 | Array                 | Map                   |
| -               | -                     | -                     | -                     | -                     |
| eq              | STRING INT FLOAT BOOL |                       |                       |                       |
| required        | STRING INT FLOAT BOOL | STRING INT FLOAT BOOL |                       | STRING INT FLOAT BOOL |
| gt              | INT FLOAT             |                       |                       |                       |
| gte             | INT FLOAT             |                       |                       |                       |
| lte             | INT FLOAT             |                       |                       |                       |
| lt              | INT FLOAT             |                       |                       |                       |
| min             | STRING                | STRING INT FLOAT BOOL |                       | STRING INT FLOAT BOOL |
| max             | STRING                | STRING INT FLOAT BOOL |                       | STRING INT FLOAT BOOL | 
| eq_ignore_case  | STRING                |                       |                       |                       |
| len             | STRING                | STRING INT FLOAT BOOL |                       | STRING INT FLOAT BOOL |
| neq             | STRING INT FLOAT BOOL |                       |                       |                       |
| neq_ignore_case | STRING                |                       |                       |                       |
| in              | STRING INT FLOAT BOOL | STRING INT FLOAT BOOL | STRING INT FLOAT BOOL | STRING INT FLOAT BOOL |
| nin             | STRING INT FLOAT BOOL | STRING INT FLOAT BOOL | STRING INT FLOAT BOOL | STRING INT FLOAT BOOL |
| email           | STRING                |                       |                       |                       |
| eqfield         | STRING INT BOOL       |                       |                       |                       |
| neqfield        | STRING INT BOOL       |                       |                       |                       |
| gtefield        | INT                   |                       |                       |                       |
| gtfield         | INT                   |                       |                       |                       |
| ltefield        | INT                   |                       |                       |                       |
| ltfield         | INT                   |                       |                       |                       |

In this table:
- **STRING** represents the `string` Go type
- **BOOL** represents the `bool` Go type
- **INT** represents all ten integer Go types: `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`
- **FLOAT** represents both float Go types: `float32`, `float64`

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

## What TestGen does

TestGen generates the following tests (without field operations):
- [x] Benchmark tests between ValidGen and GoValidator
- [x] End-to-end tests with all possible use cases (all validations vs all types vs valid and invalid inputs)
- [x] Unit tests to validate the "buildValidationCode" function

High priority generators:
- [ ] Unit tests to validate the "condition table" (get_test_elements_*_test.go)
- [ ] Benchmark tests between ValidGen and GoValidator with field operations
- [ ] End-to-end tests with all possible use cases (all validations vs all types vs valid and invalid inputs) with field operations
- [ ] Unit tests to validate the "buildValidationCode" function with field operations

Low priority generators (already exist, but could be automated):
- [ ] Unit tests to validate operations (func TestOperationsIsValid)
- [ ] Unit tests to validate operation vs type (func TestOperationsIsValidByType)
- [ ] Unit tests to validate field operations (func TestOperationsIsFieldOperation)
- [ ] Unit tests to validate argument count by operation (func TestOperationsArgsCount)
- [ ] Examples (in _examples/) could be generated

Where applicable, TestGen generates both valid and invalid test scenarios.

## Usage

To generate the tests:

```bash
# Navigate to the project root folder
cd validgen

# Run testgen
make testgen
```
