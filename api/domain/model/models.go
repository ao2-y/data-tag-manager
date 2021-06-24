package model

const (
	ContextKeyRequestID = "RequestID"
)

type (
	Item struct {
		ID          uint
		Name        string
		Description *string
		Tags        []*Tag
		Metas       []*ItemMeta
	}
	// Tag タグは親子関係のみ持てる
	Tag struct {
		ID          uint
		Name        string
		Color       string
		ParentTagID uint // 親が存在しない場合は0にする
	}
	// TagWithParent 親の情報も持っている
	TagWithParent struct {
		Tag
		Parent *Tag
	}
	// ItemMeta 属性情報
	ItemMeta struct {
		ID        uint
		MetaKeyID uint
		Value     string
		MetaKey   MetaKey
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
