// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Node interface {
	IsNode()
}

type AddItemInput struct {
	ClientMutationID *string `json:"clientMutationId"`
	Name             string  `json:"name"`
	Description      *string `json:"description"`
}

type AddItemPayload struct {
	ClientMutationID *string `json:"clientMutationId"`
	Item             *Item   `json:"item"`
}

type AddItemTemplateInput struct {
	ClientMutationID *string  `json:"clientMutationId"`
	Name             string   `json:"name"`
	MetaKeyIds       []string `json:"metaKeyIds"`
}

type AddItemTemplatePayload struct {
	ClientMutationID *string       `json:"clientMutationId"`
	ItemTemplate     *ItemTemplate `json:"itemTemplate"`
}

type AddItemWithMetaAndTagInput struct {
	ClientMutationID *string                 `json:"clientMutationId"`
	Name             string                  `json:"name"`
	Description      *string                 `json:"description"`
	Metas            []*AddItemWithMetaInput `json:"metas"`
	Tags             []*AddItemWithTagInput  `json:"tags"`
}

type AddItemWithMetaInput struct {
	MetaKeyID string `json:"metaKeyId"`
	Value     string `json:"value"`
}

type AddItemWithTagInput struct {
	TagID string `json:"tagId"`
}

type AddMetaKeyInput struct {
	ClientMutationID *string `json:"clientMutationId"`
	Name             string  `json:"name"`
}

type AddMetaKeyPayload struct {
	ClientMutationID *string  `json:"clientMutationId"`
	MetaKey          *MetaKey `json:"metaKey"`
}

type AddMetaToItemInput struct {
	ClientMutationID *string `json:"clientMutationId"`
	MetaKeyID        string  `json:"metaKeyId"`
	Value            string  `json:"value"`
	ItemID           string  `json:"itemId"`
}

type AddMetaToItemPayload struct {
	ClientMutationID *string `json:"clientMutationId"`
	Item             *Item   `json:"item"`
}

type AddTagInput struct {
	ClientMutationID *string `json:"clientMutationId"`
	ParentID         *string `json:"parentId"`
	Name             string  `json:"name"`
}

type AddTagPaylod struct {
	ClientMutationID *string `json:"clientMutationId"`
	Tag              *Tag    `json:"tag"`
}

type AddTagToItemInput struct {
	ClientMutationID *string `json:"clientMutationId"`
	ID               string  `json:"id"`
	ItemID           string  `json:"itemId"`
}

type AddTagToItemPayload struct {
	ClientMutationID *string `json:"clientMutationId"`
	Item             *Item   `json:"item"`
}

type Item struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Metas       []*Meta `json:"metas"`
	Tags        []*Tag  `json:"tags"`
}

func (Item) IsNode() {}

type ItemConnection struct {
	PageInfo *PageInfo   `json:"pageInfo"`
	Edges    []*ItemEdge `json:"edges"`
}

type ItemEdge struct {
	Cursor string `json:"cursor"`
	Node   *Item  `json:"node"`
}

type ItemTemplate struct {
	ID       string     `json:"id"`
	Name     string     `json:"name"`
	MetaKeys []*MetaKey `json:"metaKeys"`
}

func (ItemTemplate) IsNode() {}

type Meta struct {
	ID      string   `json:"id"`
	MetaKey *MetaKey `json:"metaKey"`
	Value   string   `json:"value"`
}

func (Meta) IsNode() {}

type MetaKey struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (MetaKey) IsNode() {}

type NoopInput struct {
	ClientMutationID *string `json:"clientMutationId"`
}

type NoopPayload struct {
	ClientMutationID *string `json:"clientMutationId"`
}

type PageInfo struct {
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *string `json:"startCursor"`
	HasNextPage     bool    `json:"hasNextPage"`
	EndCursor       *string `json:"endCursor"`
}

type RemoveItemInput struct {
	ClientMutationID *string `json:"clientMutationId"`
	ID               string  `json:"id"`
}

type RemoveItemPayload struct {
	ClientMutationID *string `json:"clientMutationId"`
	Item             *Item   `json:"item"`
}

type RemoveItemTemplateInput struct {
	ClientMutationID *string `json:"clientMutationId"`
	ItemTemplateID   string  `json:"itemTemplateId"`
}

type RemoveItemTemplatePayload struct {
	ClientMutationID *string       `json:"clientMutationId"`
	ItemTemplate     *ItemTemplate `json:"itemTemplate"`
}

type RemoveMetaKeyInput struct {
	ClientMutationID *string `json:"clientMutationId"`
	ID               string  `json:"id"`
}

type RemoveMetaKeyPayload struct {
	ClientMutationID *string  `json:"clientMutationId"`
	MetaKey          *MetaKey `json:"metaKey"`
}

type RemoveMetaToItemInput struct {
	ClientMutationID *string `json:"clientMutationId"`
	MetaID           string  `json:"metaId"`
}

type RemoveMetaToItemPayload struct {
	ClientMutationID *string `json:"clientMutationId"`
	Item             *Item   `json:"item"`
}

type RemoveTagInput struct {
	ClientMutaionID *string `json:"clientMutaionId"`
	ID              string  `json:"id"`
}

type RemoveTagPayload struct {
	ClientMutationID *string `json:"clientMutationId"`
	Tag              *Tag    `json:"tag"`
}

type RemoveTagToItemInput struct {
	ClientMutationID *string `json:"clientMutationId"`
	ID               string  `json:"id"`
	ItemID           string  `json:"itemId"`
}

type RemoveTagToItemPayload struct {
	ClientMutationID *string `json:"clientMutationId"`
	Item             *Item   `json:"item"`
}

type Tag struct {
	ID     string `json:"id"`
	Parent *Tag   `json:"parent"`
	Name   string `json:"name"`
}

func (Tag) IsNode() {}

type UpdateItemTemplateMetaKeysInput struct {
	ClientMutationID *string  `json:"clientMutationId"`
	ItemTemplateID   string   `json:"itemTemplateId"`
	MetaKeyIds       []string `json:"metaKeyIds"`
}

type UpdateItemTemplateNameInput struct {
	ClientMutationID *string `json:"clientMutationId"`
	ItemTemplateID   string  `json:"itemTemplateId"`
	Name             string  `json:"name"`
}

type UpdateItemTemplatePayload struct {
	ClientMutationID *string       `json:"clientMutationId"`
	ItemTemplate     *ItemTemplate `json:"itemTemplate"`
}

type UpdateMetaKeyInput struct {
	ClientMutationID *string `json:"clientMutationId"`
	ID               string  `json:"id"`
	Name             string  `json:"name"`
}

type UpdateMetaKeyPayload struct {
	ClientMutationID *string  `json:"clientMutationId"`
	MetaKey          *MetaKey `json:"metaKey"`
}

type Error string

const (
	ErrorInternalServerError Error = "INTERNAL_SERVER_ERROR"
	ErrorValidationError     Error = "VALIDATION_ERROR"
	ErrorResourceNotFound    Error = "RESOURCE_NOT_FOUND"
	ErrorOptimiseLockError   Error = "OPTIMISE_LOCK_ERROR"
	ErrorAuthError           Error = "AUTH_ERROR"
)

var AllError = []Error{
	ErrorInternalServerError,
	ErrorValidationError,
	ErrorResourceNotFound,
	ErrorOptimiseLockError,
	ErrorAuthError,
}

func (e Error) IsValid() bool {
	switch e {
	case ErrorInternalServerError, ErrorValidationError, ErrorResourceNotFound, ErrorOptimiseLockError, ErrorAuthError:
		return true
	}
	return false
}

func (e Error) String() string {
	return string(e)
}

func (e *Error) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Error(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Error", str)
	}
	return nil
}

func (e Error) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
