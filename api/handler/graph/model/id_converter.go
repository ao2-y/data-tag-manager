package model

import (
	"fmt"
	"strconv"
	"strings"
)

type IDType string

const (
	IDTypeUnknown      IDType = "Unknown"
	IDTypeItemTemplate IDType = "ItemTemplate:"
	IDTypeMeta         IDType = "Meta:"
	IDTypeMetaKey      IDType = "MetaKey:"
	IDTypeTag          IDType = "Tag:"
	IDTypeItem         IDType = "Item:"
)

var keys = []IDType{IDTypeItemTemplate, IDTypeItem, IDTypeMeta, IDTypeTag}

func (id IDType) ToExternalID(ID uint) string {
	return fmt.Sprintf("%s%v", id, ID)
}

func (id IDType) ToInternalID(ID string) (uint, error) {
	noStr := ID[len(id):]
	no, err := strconv.Atoi(noStr)
	if err != nil {
		return 0, fmt.Errorf("parse error:%w", err)
	}
	return uint(no), nil
}

// IDtoKeyNameAndInternalID External ID to Internal ID and IDType
func IDtoKeyNameAndInternalID(ID string) (uint, IDType, error) {
	for _, v := range keys {
		if strings.HasPrefix(ID, string(v)) {
			no, err := v.ToInternalID(ID)
			return no, v, err
		}
	}
	return 0, IDTypeUnknown, fmt.Errorf("ID type not found")
}
