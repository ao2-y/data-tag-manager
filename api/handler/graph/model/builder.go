package model

import "ao2-y/data-tag-manager/domain/model"

func NewItemTemplate(template *model.ItemTemplate) *ItemTemplate {
	metaKeys := make([]*MetaKey, len(template.MetaKeys), len(template.MetaKeys))
	for i, v := range template.MetaKeys {
		metaKey := &MetaKey{
			ID:   IDTypeMeta.ToExternalID(v.ID),
			Name: v.Name,
		}
		metaKeys[i] = metaKey
	}
	return &ItemTemplate{
		ID:       IDTypeItemTemplate.ToExternalID(template.ID),
		Name:     template.Name,
		MetaKeys: metaKeys,
	}
}

func NewItemTemplates(templates []*model.ItemTemplate) []*ItemTemplate {
	ret := make([]*ItemTemplate, len(templates), len(templates))
	for i, v := range templates {
		ret[i] = NewItemTemplate(v)
	}
	return ret
}

func NewItem(item *model.Item) *Item {
	return &Item{
		ID:          IDTypeItem.ToExternalID(item.ID),
		Name:        item.Name,
		Description: item.Description,
		Metas:       NewMetas(item.Metas),
		Tags:        nil,
	}
}

func NewTag(tag *model.Tag) *Tag {
	if tag == nil {
		return nil
	}
	return &Tag{
		ID:   IDTypeTag.ToExternalID(tag.ID),
		Name: tag.Name,
	}
}

func NewTagWithParent(tag *model.TagWithParent) *Tag {
	return &Tag{
		ID:     IDTypeTag.ToExternalID(tag.ID),
		Parent: NewTag(tag.Parent),
		Name:   tag.Name,
	}
}

func NewMetas(metas []*model.Meta) []*Meta {
	ret := make([]*Meta, len(metas), len(metas))
	for i, v := range metas {
		ret[i] = NewMeta(v)
	}
	return ret
}

func NewMeta(meta *model.Meta) *Meta {
	return &Meta{
		ID:      IDTypeMeta.ToExternalID(meta.ID),
		MetaKey: NewMetaKey(&meta.MetaKey),
		Value:   meta.Value,
	}
}

func NewMetaKey(metaKey *model.MetaKey) *MetaKey {
	return &MetaKey{
		ID:   IDTypeMetaKey.ToExternalID(metaKey.ID),
		Name: metaKey.Name,
	}
}
