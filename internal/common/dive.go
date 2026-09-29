package common

import (
	"fmt"
	"strings"
)

// IsLenOperation reports tags that only check the length of a string, slice, or map.
func IsLenOperation(op string) bool {
	switch op {
	case "required", "min", "max", "len":
		return true
	default:
		return false
	}
}

// IsNestedStruct reports a struct or pointer to a struct.
// Go scalars and containers are not nested structs.
func (ft FieldType) IsNestedStruct() bool {
	if ft.ComposedType != "" && ft.ComposedType != "*" {
		return false
	}

	return ft.BaseType != "" && !ft.IsGoType()
}

// PointerToContainer reports *[]T, *[N]T, and *map[K]V.
// []*T keeps the star on the element and returns false.
func (ft FieldType) PointerToContainer() bool {
	if ft.ElemPointer || !strings.HasPrefix(ft.ComposedType, "*") {
		return false
	}

	ct := strings.TrimPrefix(ft.ComposedType, "*")
	return isContainerBody(ct)
}

// ForCatalog drops an element pointer so []*T is checked as a slice of T.
// *[]T is left unchanged.
func (ft FieldType) ForCatalog() FieldType {
	if !ft.ElemPointer {
		return ft
	}

	ft.ElemPointer = false
	ft.ComposedType = strings.Replace(ft.ComposedType, "*", "", 1)
	return ft
}

// LenStandIn is a slice or map of string used when a length tag does not
// depend on the element type, such as []Address or [][]string.
func (ft FieldType) LenStandIn() (FieldType, bool) {
	ct := ft.ComposedType
	ptr := ""
	if strings.HasPrefix(ct, "*") && !ft.ElemPointer {
		ptr = "*"
		ct = strings.TrimPrefix(ct, "*")
	} else if ft.ElemPointer {
		ct = strings.Replace(ct, "*", "", 1)
	}

	switch {
	case strings.HasSuffix(ct, "[]"):
		return FieldType{ComposedType: ptr + "[]", BaseType: "string"}, true
	case ct == "map":
		return FieldType{ComposedType: ptr + "map", BaseType: "string"}, true
	default:
		return FieldType{}, false
	}
}

// OperationType selects the field type used to look up op.
// accept reports whether that type supports the operation.
// in and nin on []*T are rejected: those helpers expect []T, not pointers.
// Length tags only call len, so they still apply to []*T.
func (ft FieldType) OperationType(op string, accept func(FieldType) bool) (FieldType, bool) {
	if ft.ElemPointer && !IsLenOperation(op) {
		return FieldType{}, false
	}

	lookup := ft.ForCatalog()
	if knownComposed(lookup.ComposedType) && lookup.IsGoType() && accept(lookup) {
		return lookup, true
	}

	if IsLenOperation(op) {
		if standIn, ok := ft.LenStandIn(); ok && accept(standIn) {
			return standIn, true
		}
	}

	return FieldType{}, false
}

// DiveInto returns the type of one slice element, array element, or map value.
func (ft FieldType) DiveInto() (FieldType, error) {
	ct := ft.ComposedType
	if strings.HasPrefix(ct, "*") && !ft.ElemPointer {
		ct = strings.TrimPrefix(ct, "*")
	}

	switch {
	case strings.HasSuffix(ct, "[]"):
		return finishDiveElement(FieldType{
			BaseType:     ft.BaseType,
			ComposedType: strings.TrimSuffix(ct, "[]"),
			MapValue:     ft.MapValue,
		})
	case strings.HasSuffix(ct, "[N]"):
		elem := FieldType{
			BaseType:     ft.BaseType,
			ComposedType: strings.TrimSuffix(ct, "[N]"),
			MapValue:     ft.MapValue,
		}
		if strings.Contains(elem.ComposedType, "[N]") {
			return FieldType{}, fmt.Errorf("nested arrays are not supported")
		}
		return finishDiveElement(elem)
	case ct == "map":
		if ft.MapValue == nil {
			return FieldType{}, fmt.Errorf("map value type is not available")
		}
		return *ft.MapValue, nil
	default:
		return FieldType{}, fmt.Errorf("%s is not a slice, array, or map", ft.ToType())
	}
}

func finishDiveElement(elem FieldType) (FieldType, error) {
	// A star glued to another container (*[]T versus []*T) cannot be recovered
	// once both levels share one ComposedType string.
	if strings.Contains(elem.ComposedType, "*") && elem.ComposedType != "*" {
		return FieldType{}, fmt.Errorf("unsupported pointer composition %s%s", elem.ComposedType, elem.BaseType)
	}

	return elem, nil
}

func knownComposed(ct string) bool {
	switch ct {
	case "", "*", "[]", "[N]", "map", "*[]", "*[N]", "*map":
		return true
	default:
		return false
	}
}

func isContainerBody(ct string) bool {
	return strings.HasSuffix(ct, "[]") || strings.HasSuffix(ct, "[N]") || ct == "map" || strings.HasSuffix(ct, "map")
}
