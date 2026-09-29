package analyzer

import (
	"strings"

	"github.com/opencodeco/validgen/internal/analyzer/operations"
	"github.com/opencodeco/validgen/internal/common"
	"github.com/opencodeco/validgen/types"
)

type Validation struct {
	Operation      string
	ExpectedValues common.CountValues
	Values         []string
}

func ParserValidation(fieldValidation string) (*Validation, error) {
	validation, values, err := parserValidationString(fieldValidation)
	if err != nil {
		return nil, err
	}

	ops := operations.New()

	valuesCount := ops.ArgsCount(validation)
	if valuesCount == common.UndefinedValue {
		return nil, types.NewValidationError("unsupported validation %s", validation)
	}

	switch valuesCount {
	case common.ZeroValue:
		return parserZeroValue(validation, valuesCount, values)
	case common.OneValue:
		return parserOneValue(validation, valuesCount, values)
	case common.ManyValues:
		return parserManyValues(validation, valuesCount, values)
	default:
		return nil, types.NewValidationError("invalid value in validation %s", validation)
	}
}

func parserValidationString(tag string) (string, string, error) {
	tokens := removeEmptyValues(strings.Split(tag, "="))
	if len(tokens) == 0 || len(tokens) > 2 {
		return "", "", types.NewValidationError("malformed validation %s", tag)
	}

	validation := canonicalOperation(strings.TrimSpace(tokens[0]))
	values := ""
	if len(tokens) == 2 {
		values = tokens[1]
	}

	return validation, values, nil
}

func canonicalOperation(validation string) string {
	switch validation {
	case "ne":
		return "neq"
	case "ne_ignore_case":
		return "neq_ignore_case"
	default:
		return validation
	}
}

func parserZeroValue(validation string, valuesCount common.CountValues, targets string) (*Validation, error) {
	if targets != "" {
		return nil, types.NewValidationError("expected zero target, but has %s", targets)
	}

	return &Validation{
		Operation:      validation,
		ExpectedValues: valuesCount,
		Values:         []string{},
	}, nil
}

func parserOneValue(validation string, valuesCount common.CountValues, targets string) (*Validation, error) {
	if targets == "" {
		return nil, types.NewValidationError("expected one target, but has nothing")
	}

	return &Validation{
		Operation:      validation,
		ExpectedValues: valuesCount,
		Values:         []string{targets},
	}, nil
}

func parserManyValues(validation string, valuesCount common.CountValues, targets string) (*Validation, error) {
	if len(targets) == 0 {
		return nil, types.NewValidationError("expected at least one target, but has 0 element(s)")
	}

	values, err := splitManyValues(targets)
	if err != nil {
		return nil, err
	}
	if len(values) == 0 {
		return nil, types.NewValidationError("expected at least one target, but has 0 element(s)")
	}

	return &Validation{
		Operation:      validation,
		ExpectedValues: valuesCount,
		Values:         values,
	}, nil
}

// splitManyValues reads a oneof/in list. Spaces and commas separate tokens.
// A single-quoted token may contain spaces, and quoted tokens may be mixed
// with bare tokens, matching go-playground/validator oneof.
func splitManyValues(targets string) ([]string, error) {
	values := []string{}
	for i := 0; i < len(targets); {
		if targets[i] == ' ' || targets[i] == ',' {
			i++
			continue
		}

		if targets[i] == '\'' {
			end := strings.IndexByte(targets[i+1:], '\'')
			if end == -1 {
				return nil, types.NewValidationError("invalid quote value in %s", targets)
			}
			values = append(values, targets[i+1:i+1+end])
			i += end + 2
			continue
		}

		start := i
		for i < len(targets) && targets[i] != ' ' && targets[i] != ',' {
			i++
		}
		values = append(values, targets[start:i])
	}

	return values, nil
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
