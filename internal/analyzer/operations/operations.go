package operations

import (
	"slices"
	"strings"

	"github.com/opencodeco/validgen/internal/common"
)

type Operation struct {
	CountValues      common.CountValues
	IsFieldOperation bool
	ValidTypes       []string
}

type Operations struct {
	operations map[string]Operation
}

func New() *Operations {
	return &Operations{
		operations: operationsList,
	}
}

func (o *Operations) IsValid(op string) bool {
	_, ok := o.operations[op]

	return ok
}

func (o *Operations) IsValidByType(op, fieldType string) bool {
	// * is a modifier and can be ignored for type validation.
	normalized, pointer := strings.CutPrefix(fieldType, "*")

	// Scalar complex pointers have no condition-table row. Reject them here
	// so analysis does not accept a tag that generation cannot emit.
	if pointer && normalized == "<COMPLEX>" {
		return false
	}

	// Required can be used with all pointer types.
	if pointer && op == "required" {
		// Required can be used with all pointers types.
		return true
	}

	return slices.Contains(o.operations[op].ValidTypes, normalized)
}

func (o *Operations) IsFieldOperation(op string) bool {
	return o.operations[op].IsFieldOperation
}

func (o *Operations) ArgsCount(op string) common.CountValues {
	return o.operations[op].CountValues
}
