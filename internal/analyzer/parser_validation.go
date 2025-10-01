package analyzer

import (
	"strings"

	"github.com/opencodeco/validgen/types"
)

type Validation struct {
	Operation      string
	ExpectedValues CountValues
	Values         []string
}

func ParserValidation(fieldValidation string) (*Validation, error) {
	validation, values, err := parserValidationString(fieldValidation)
	if err != nil {
		return nil, err
	}

	valuesCount := operations[validation].CountValues
	if valuesCount == UNDEFINED {
		return nil, types.NewValidationError("unsupported validation %s", validation)
	}

	switch valuesCount {
	case ZERO_VALUE:
		return parserZeroValue(validation, valuesCount, values)
	case ONE_VALUE:
		return parserOneValue(validation, valuesCount, values)
	case MANY_VALUES:
		return parserManyValues(validation, valuesCount, values)
	default:
		return nil, types.NewValidationError("invalid value in validation %s", validation)
	}
}

func parserValidationString(tag string) (string, string, error) {
	tokens := removeEmptyValues(strings.Split(tag, "="))
	if len(tokens) > 2 {
		return "", "", types.NewValidationError("malformed validation %s", tag)
	}

	validation := strings.TrimSpace(tokens[0])
	values := ""
	if len(tokens) == 2 {
		values = tokens[1]
	}

	return validation, values, nil
}

func parserZeroValue(validation string, valuesCount CountValues, targets string) (*Validation, error) {
	if targets != "" {
		return nil, types.NewValidationError("expected zero target, but has %s", targets)
	}

	return &Validation{
		Operation:      validation,
		ExpectedValues: valuesCount,
		Values:         []string{},
	}, nil
}

func parserOneValue(validation string, valuesCount CountValues, targets string) (*Validation, error) {
	if targets == "" {
		return nil, types.NewValidationError("expected one target, but has nothing")
	}

	return &Validation{
		Operation:      validation,
		ExpectedValues: valuesCount,
		Values:         []string{targets},
	}, nil
}

func parserManyValues(validation string, valuesCount CountValues, targets string) (*Validation, error) {
	if len(targets) == 0 {
		return nil, types.NewValidationError("expected at least one target, but has 0 element(s)")
	}

	targetValues := targets

	if targetValues[0] == '\'' {
		values := []string{}
		for {
			first := strings.IndexByte(targetValues, '\'')
			if first == -1 {
				break
			}

			if first != 0 {
				// ' must be the first chr
				return nil, types.NewValidationError("invalid quote value in %s", targets)
			}

			second := strings.IndexByte(targetValues[first+1:], '\'')
			if second == -1 {
				return nil, types.NewValidationError("invalid quote value in %s", targets)
			}
			values = append(values, targetValues[first+1:second+1])
			targetValues = strings.TrimSpace(targetValues[second+2:])
		}

		return &Validation{
			Operation:      validation,
			ExpectedValues: valuesCount,
			Values:         values,
		}, nil
	}

	// Break by commas or spaces.
	targetValues = strings.ReplaceAll(targetValues, " ", ",")
	values := strings.Split(targetValues, ",")
	values = removeEmptyValues(values)

	return &Validation{
		Operation:      validation,
		ExpectedValues: valuesCount,
		Values:         values,
	}, nil
}

func removeEmptyValues(input []string) []string {
	var output []string

	for _, element := range input {
		if strings.TrimSpace(element) != "" {
			output = append(output, strings.TrimSpace(element))
		}
	}

	return output
}
