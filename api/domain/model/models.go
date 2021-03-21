package model

type (
	Item struct {
		ID          uint
		Name        string
		Description *string
		Tags        []Tag
		Metas       []Meta
	}
	Tag struct {
		ID          uint
		Name        string
		Level       int8
		ParentTagId *uint
	}
	Meta struct {
		ID    uint
		Name  string
		Value string
	}
)
