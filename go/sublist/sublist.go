package sublist

import (
	"bytes"
	"reflect"
)

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {
	bytesL1 := ConvertSliceToByte(l1)
	bytesL2 := ConvertSliceToByte(l2)
	switch {
	case reflect.DeepEqual(bytesL1, bytesL2):
		return RelationEqual
	case bytes.Contains(bytesL1, bytesL2):
		return RelationSuperlist
	case bytes.Contains(bytesL2, bytesL1):
		return RelationSublist
	default:
		return RelationUnequal
	}

}

func ConvertSliceToByte(slice []int) []byte {
	bt := []byte{}

	for _, vl := range slice {
		bt = append(bt, byte(vl))
	}

	return bt
}
