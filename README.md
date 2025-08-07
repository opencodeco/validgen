# ValidGen

[Validator](https://github.com/go-playground/validator) is an amazing project. But in applications with high frequency use, the way that validator works (i.e. reflection) can cause performance penalty.

ValidGen born to solve that gap. Instead of use reflection, ValidGen uses the code generating approach.

At this time it is an unstable project and should not be used in production environments.

# How to build ValidGen

The following requirements are needed to build the project:
- Git
- Go >= 1.24
- Make

The steps to build are:
```
# Clone the project repository
git clone git@github.com:opencodeco/validgen.git

# Enter in the project root folder
cd validgen

# Build the binary
make build
```

After that the executable will be in `bin/validgen`.

# Validations

The following validations will be implemented:

- eq (equal): must be equal to the specified value
- eq_ignore_case (equal ignoring case): must be equal to the specified value (ignoring case)
- gt (greater than): must be > to the specified value
- gte (greater than or equal): must be >= to the specified value
- lt (less than): must be < to the specified value
- lte (less than or equal): must be <= to the specified value
- neq (not equal): must not be equal to the specified value
- neq_ignore_case (not equal ignoring case): must not be equal to the specified value (ignoring case)
- len (length): must have the following length
- max (max): must have no more than max characters
- min (min): must have no less than min characters
- in (in): must be one of the following values
- nin (not in): must not be one of the following values
- required (required): is required
- email (email): must be a valid email format (empty is valid for optional fields)

The following table shows the validations and possible types, where "I" means "Implemented", "W" means "Will be implemented" and "-" means "Will not be implemented":

| Validation/Type | String | Numeric types | Boolean | Slice | Array | Map | Time | Duration |
| -               | -      | -             | -       | -     | -     | -   | -    | -        |
| eq              | I      | W             | W       | -     | -     | -   | W    | W        |
| eq_ignore_case  | I      | -             | -       | -     | -     | -   | -    | -        |
| gt              | -      | W             | -       | -     | -     | -   | W    | W        |
| gte             | -      | W             | -       | -     | -     | -   | W    | W        |
| lt              | -      | W             | -       | -     | -     | -   | W    | W        |
| lte             | -      | W             | -       | -     | -     | -   | W    | W        |
| neq             | I      | W             | W       | -     | -     | -   | W    | W        |
| neq_ignore_case | I      | -             | -       | -     | -     | -   | -    | -        |
| len             | I      | -             | -       | I     | W     | W   | -    | -        |
| max             | I      | -             | -       | I     | W     | W   | W    | W        |
| min             | I      | -             | -       | I     | W     | W   | W    | W        |
| in              | I      | W             | W       | W     | W     | W   | -    | W        |
| nin             | I      | W             | W       | W     | W     | W   | -    | W        |
| required        | I      | W             | W       | I     | W     | W   | W    | W        |
| email           | I      | -             | -       | -     | -     | -   | -    | -        |

# Steps to run the unit tests

The steps to run the unit tests are:

```
# Enter in the project root folder
cd validgen

# Run the unit tests
make unittests
```

# Steps to run the benchmark tests

The steps to run the benchmark tests are:

```
# Enter in the project root folder
cd validgen

# Run the benchmark tests
make benchtests
```

# Steps to run the end-to-end tests

The steps to run the end-to-end tests are:

```
# Enter in the project root folder
cd validgen

# Run the end-to-end tests
make endtoendtests
```

# Steps to run the examples

All examples are in the `_examples` folder.

## Steps to run test01

Test01 aims to be a case where all the files are in the same package (in this case, the main package).

```
# Runs validgen to generate structs validator code
./bin/validgen _examples/test01
```

After that the file `user_validator.go` will be generated. This file contains UserValidate function that is responsible to check if User object has a valid content.

```
# Execute the test
cd _examples/test01
go run .
```

## Steps to run test02

Test02 aims to be an example where the structs to be validated are in another package (structsinpkg in this test).

```
# Runs validgen to generate structs validator code
./bin/validgen _examples/test02
```

After that the file `user_validator.go` will be generated. This file contains UserValidate function that is responsible to check if User object has a valid content.

```
# Execute the test
cd _examples/test02
go run .
```

## Steps to run test03

Test03 aims to be an example where the structs to be validated use min and max tags.

```
# Runs validgen to generate structs validator code
./bin/validgen _examples/test03
```

After that the file `user_validator.go` will be generated. This file contains UserValidate function that is responsible to check if User object has a valid content.

```
# Execute the test
cd _examples/test03
go run .
```


# Steps to run the benchmark tests comparing ValidGen and Validator

The steps to run the benchmark tests are:

```
# Enter in the project root folder
cd validgen

# Run the benchmark tests
make cmpbenchtests
```


The command `make cmpbenchtests` invoke the following command:

`go test -bench=. -v -benchmem -benchtime=5s ./tests/cmpbenchtests/generated_tests`

The setup used was:

```
goos: darwin
goarch: arm64
pkg: github.com/opencodeco/validgen/tests/cmpbenchtests/generated_tests
cpu: Apple M4 Pro (12 Cores used)
```

The following table as the performance results:

| Test name      | ValidGen    | GoValidator | Performance |
| -              | -:          | -:          | -:          |
| StringRequired | 5.000 ns/op | 40.12 ns/op | 8.02x       |
| StringEq       | 5.000 ns/op | 39.90 ns/op | 7.98x       |
| StringEqIC     | 14.96 ns/op | 40.63 ns/op | 2.71x       |
| StringNeq      | 5.000 ns/op | 40.65 ns/op | 8.13x       |
| StringNeqIC    | 5.000 ns/op | 41.05 ns/op | 8.21x       |
| StringLen      | 5.000 ns/op | 44.35 ns/op | 8.87x       |
| StringMax      | 5.000 ns/op | 45.15 ns/op | 9.03x       |
| StringMin      | 5.000 ns/op | 45.10 ns/op | 9.02x       |
| StringIn       | 5.000 ns/op | 49.75 ns/op | 9.95x       |
| StringEmail    | 167.4 ns/op | 436.3 ns/op | 2.60x       |


The following table as the raw results:

| Test name                           | Iterations    | Nanoseconds per operation | Number os bytes allocated per operation | Number of allocations per operation |
| -                                   | -:            | -:                        | -:                                      | -:                                  | 
| BenchmarkValidGenStringRequired-12  |	1.000.000.000 |          5.000 ns/op      |        0 B/op	                        |       0 allocs/op                   |
| BenchmarkValidatorStringRequired-12 | 149.656.641	  |        40.12 ns/op	      |        0 B/op	                        |       0 allocs/op                   |
| BenchmarkValidGenStringEq-12        | 1.000.000.000 |          5.000 ns/op	  |        0 B/op	                        |       0 allocs/op                   |
| BenchmarkValidatorStringEq-12       | 150.419.842	  |         39.90 ns/op	      |        0 B/op	                        |       0 allocs/op                   |
| BenchmarkValidGenStringEqIC-12      | 401.652.991	  |         14.96 ns/op	      |        8 B/op	                        |       1 allocs/op                   |
| BenchmarkValidatorStringEqIC-12     | 147.517.011	  |         40.63 ns/op	      |        0 B/op	                        |       0 allocs/op                   |
| BenchmarkValidGenStringNeq-12       | 1.000.000.000 |          5.000 ns/op	  |        0 B/op	                        |       0 allocs/op                   |
| BenchmarkValidatorStringNeq-12      | 147.375.966   |         40.65 ns/op	      |        0 B/op	                        |       0 allocs/op                   |
| BenchmarkValidGenStringNeqIC-12     | 1.000.000.000 |          5.000 ns/op	  |        0 B/op	                        |       0 allocs/op                   |
| BenchmarkValidatorStringNeqIC-12    | 146.089.116	  |         41.05 ns/op	      |        0 B/op	                        |       0 allocs/op                   |
| BenchmarkValidGenStringLen-12       | 1.000.000.000 |          5.000 ns/op	  |        0 B/op	                        |       0 allocs/op                   |
| BenchmarkValidatorStringLen-12      | 135.221.859	  |         44.35 ns/op	      |        0 B/op	                        |       0 allocs/op                   |
| BenchmarkValidGenStringMax-12       | 1.000.000.000 |          5.000 ns/op	  |        0 B/op	                        |       0 allocs/op                   |
| BenchmarkValidatorStringMax-12      | 133.947.687	  |         45.15 ns/op	      |        0 B/op	                        |       0 allocs/op                   |
| BenchmarkValidGenStringMin-12       | 1.000.000.000 |          5.000 ns/op	  |        0 B/op	                        |       0 allocs/op                   |
| BenchmarkValidatorStringMin-12      | 133.335.433	  |         45.10 ns/op	      |        0 B/op	                        |       0 allocs/op                   |
| BenchmarkValidGenStringIn-12        | 1.000.000.000 |          5.000 ns/op	  |        0 B/op	                        |       0 allocs/op                   |
| BenchmarkValidatorStringIn-12       | 120.405.889	  |         49.75 ns/op	      |        0 B/op	                        |       0 allocs/op                   |
| BenchmarkValidGenStringEmail-12     | 35.513.684	  |        167.4 ns/op	      |        0 B/op	                        |       0 allocs/op                   |
| BenchmarkValidatorStringEmail-12    | 13.740.566	  |        436.3 ns/op	      |       88 B/op	                        |       5 allocs/op                   |



# License

ValidGen uses [MIT License](LICENSE). 
