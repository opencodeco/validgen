package common

type NormalizedBaseType int

const (
	InvalidType NormalizedBaseType = iota
	StringType
	BoolType
	IntType
	FloatType
)

func (n NormalizedBaseType) String() string {
	switch n {
	case StringType:
		return "<STRING>"
	case BoolType:
		return "<BOOL>"
	case IntType:
		return "<INT>"
	}

	return "<INVALID>"
}
