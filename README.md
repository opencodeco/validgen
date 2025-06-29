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

After that the executable will be in bin/validgen.

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

Test03 aims to be an example where the structs to be validated use gte and lte tags.

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
