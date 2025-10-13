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

func (o *Operations) IsValidByType(op, fieldType string) bool {
	return slices.Contains(o.operations[op].ValidTypes, fieldType)
}

func (o *Operations) IsFieldOperation(op string) bool {
	return o.operations[op].IsFieldOperation
}

func (o *Operations) ArgsCount(op string) common.CountValues {
	return o.operations[op].CountValues
}
