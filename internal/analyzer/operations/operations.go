package operations

import (
	"slices"

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

func (o *Operations) IsValidType(op, fieldType string) bool {
	operation, ok := o.operations[op]
	if !ok {
		return false
	}

	return slices.Contains(operation.ValidTypes, fieldType)
}

func (o *Operations) IsFieldOperation(op string) bool {
	operation, ok := o.operations[op]
	if !ok {
		return false
	}

	return operation.IsFieldOperation
}

func (o *Operations) ArgsCount(op string) common.CountValues {
	operation, ok := o.operations[op]
	if !ok {
		return common.UndefinedValue
	}

	return operation.CountValues
}
