package graph

import (
	"ao2-y/data-tag-manager/handler/graph/model"
	"ao2-y/data-tag-manager/usecase"
	"errors"
	"fmt"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func newGraphqlError(msg string, err error) error {
	// 独自Error型から対応するコードに変換して返す
	var validationError *usecase.ValidationError
	var resourceNotFoundError *usecase.ResourceNotFoundError
	var internalServerError *usecase.InternalServerError
	if errors.As(err, &validationError) {
		// バリデーションエラー
		return &gqlerror.Error{
			Message: fmt.Sprintf("%s. %v", msg, err),
			Extensions: map[string]interface{}{
				"code":   model.ErrorValidationError,
				"field:": validationError.Field,
			},
		}
	}
	if errors.As(err, &resourceNotFoundError) {
		// リソースが存在しなかった
		return &gqlerror.Error{
			Message: err.Error(),
			Extensions: map[string]interface{}{
				"code": model.ErrorResourceNotFound,
			},
		}
	}
	if errors.As(err, &internalServerError) {
		return &gqlerror.Error{
			Message: fmt.Sprintf("%s. %v", msg, err),
			Extensions: map[string]interface{}{
				"code": model.ErrorInternalServerError,
			},
		}
	}
	return &gqlerror.Error{
		Message: fmt.Sprintf(msg, err),
		Extensions: map[string]interface{}{
			"code": model.ErrorInternalServerError,
		},
	}
}
