package model

const (
	ContextKeyRequestID = "RequestID"
)

type (
	Item struct {
		ID          uint
		Name        string
		Description *string
		Tags        []Tag
		Metas       []Meta
	}
	// Tag タグは親子関係のみ持てる
	Tag struct {
		ID          uint
		Name        string
		ParentTagID uint // 親が存在しない場合は0にする
	}
	// TagWithParent 親の情報も持っている
	TagWithParent struct {
		Tag
		parent *Tag
	}
	// Meta 属性情報
	Meta struct {
		ID        uint
		MetaKeyID uint
		Value     string
	}
	// MetaKey Meta属性名
	MetaKey struct {
		ID   uint
		Name string
	}
	ItemTemplate struct {
		ID       uint
		Name     string
		MetaKeys []*MetaKey
	}
)
