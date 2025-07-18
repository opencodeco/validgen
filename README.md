# ValidGen

[Validator](https://github.com/go-playground/validator) is an amazing project. But in applications with high frequency use, the way that validator works (i.e. reflection) can cause performance penalty.

ValidGen born to solve that gap. Instead of use reflection, ValidGen uses the code generating approach.

This project aims to be compatible with validator tag syntax.

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
| len             | I      | -             | -       | W     | W     | W   | -    | -        |
| max             | I      | -             | -       | W     | W     | W   | W    | W        |
| min             | I      | -             | -       | W     | W     | W   | W    | W        |
| in              | I      | W             | W       | W     | W     | W   | -    | W        |
| nin             | I      | W             | W       | W     | W     | W   | -    | W        |
| required        | I      | W             | W       | W     | W     | W   | W    | W        |

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

# License

ValidGen uses [MIT License](LICENSE). 
