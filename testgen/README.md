# TestGen

TestGen is the tool responsible for generating tests related to ValidGen validators and supported types.
These generated tests included unit tests, end-to-end tests, and benchmark tests.
In the case of benchmark tests, it means comparing ValidGen and GoValidator.

## Why a new tool?

First of all, it is necessary to answer the question: why a new tool to generate tests?

Currently, ValidGen has 21 possible validations with support for some types:

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

In this table, STRING represents the string Go type, and BOOL represents the bool Go type.
But INT represents the ten integer Go types: int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64.
In the same way, FLOAT represents the 2 float go types: float32, float64.
In the same way, slice STRING is just []string, but slice INT is split between all integer Go types.
For array and map, the rule is the same.
And, an important modifier is "*" (pointer) because the code generated is different with or without a pointer.

For each possible combination (validation x type) is necessary to have the following tests:
- Unit test in the analyzer phase
- Unit test in the code generator phase
- Benchmark test between ValidGen and GoValidator
- End-to-end test

As it is necessary to test all valid and invalid scenarios, it is necessary to test all validations against all types.
At this time, it means:
- Go types (14)
- Validations (21)
- Test types (4)
- Types with and without a pointer (2)
- Tests with valid and invalid inputs (2)

14 x 21 x 4 x 2 x 2 = 4.704 distinct test cases :-)

With all the tests that need to be created (valid inputs, invalid inputs) for unit tests, benchmark, and end-to-end tests, creating the tests "by hand" is a tedious and error-prone task.

Some of these necessary unit tests were created "by hand", but it is a pain to keep the code in sync when supporting new operations and types.

At this time, ValidGen already has two generators:
- To generate benchmark tests between ValidGen and GoValidator
- To generate end-to-end tests
    - to validate ValidGen with integer types
    - to validate ValidGen with float types
    - to validate all possible use cases in ValidGen (all validations x all types)

But these generators do not have a common configuration, do not implement all tests for all cases, and keeping the distinct configuration files in sync is painful.

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

Low priority generators (already exist, but they could be generated):
- [ ] Unit tests to validate operations (func TestOperationsIsValid)
- [ ] Unit tests to validate operation vs type (func TestOperationsIsValidByType)
- [ ] Unit tests to validate if is field operation (func TestOperationsIsFieldOperation)
- [ ] Unit tests to validate arguments count by operation (func TestOperationsArgsCount)
- [ ] Examples (in _examples/) could be generated

In some cases, valid scenarios and invalid scenarios must be generated.

## How TestGen works

To be possible to generate all these tests, TestGen must have a configuration with:
- All valid operations
- All valid go types
- All valid operation x type
- Valid input cases for each operation x type
- Invalid input cases for each operation x type
- Equivalent GoValidator tag

## Steps to generate the tests

The steps to generate the tests are:

```bash
# Enter in the project root folder
cd validgen

# Run testgen
make testgen
```

