package model

import "ao2-y/data-tag-manager/domain/model"

func NewItemTemplate(template *model.ItemTemplate) *ItemTemplate {
	metaKeys := make([]*MetaKey, len(template.MetaKeys), len(template.MetaKeys))
	for i, v := range template.MetaKeys {
		metaKey := &MetaKey{
			ID:   KeyMeta.ToExternalID(v.ID),
			Name: v.Name,
		}
		metaKeys[i] = metaKey
	}
	return &ItemTemplate{
		ID:       KeyItemTemplate.ToExternalID(template.ID),
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
