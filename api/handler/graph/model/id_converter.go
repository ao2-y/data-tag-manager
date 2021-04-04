package model

import (
	"fmt"
	"strconv"
	"strings"
)

type KeyType string

const (
	KeyUnknown      KeyType = "Unknown"
	KeyItemTemplate KeyType = "ItemTemplate"
	KeyMeta         KeyType = "Meta"
	KeyTag          KeyType = "Tag"
	KeyItem         KeyType = "Item"
)

var keys = []KeyType{KeyItemTemplate, KeyItem, KeyMeta, KeyTag}

func (key KeyType) ToExternalID(ID uint) string {
	return fmt.Sprintf("%s:%v", key, ID)
}

func (key KeyType) ToInternalID(ID string) uint {
	return 0 // TODO FIXME 未実装
}

// IDtoKeyNameAndInternalID External ID to Internal ID and KeyType
func IDtoKeyNameAndInternalID(ID string) (uint, KeyType, error) {
	for _, v := range keys {
		if strings.HasPrefix(ID, string(v)) {
			// Uint部分を抽出
			noStr := ID[len(v)+1:]
			no, err := strconv.Atoi(noStr)
			if err != nil {
				return 0, KeyUnknown, fmt.Errorf("ID convert fail:%w", err)
			}
			return uint(no), v, nil
		}
	}
	return 0, KeyUnknown, fmt.Errorf("ID type not found")
}
