package usecase

import (
	"ao2-y/data-tag-manager/domain/model"
	"ao2-y/data-tag-manager/domain/repository"
	repository2 "ao2-y/data-tag-manager/domain/repository/mock"
	"context"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestNewMetaUseCase(t *testing.T) {
	type args struct {
		repository repository.Meta
	}
	ctrl := gomock.NewController(t)
	repo := repository2.NewMockMeta(ctrl)
	tests := []struct {
		name string
		args args
		want Meta
	}{
		{
			name:"正常",
			args: args{repository: repo },
			want: &metaUseCase{repository: repo},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMetaUseCase(tt.args.repository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMetaUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TODO UseCaseに引き渡すMockの動作を生成するファクトリ関数
type createMetaUseCaseInitMockFunc func (ctrl *gomock.Controller) Meta

func Test_metaUseCase_CreateKey(t *testing.T) {
	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name    string
		init  createMetaUseCaseInitMockFunc
		args    args
		want    *model.MetaKey
		wantErr bool
	}{
		{
			name: "正常",
			wantErr: false,
			init: createMetaUseCaseInitMockFunc(ctrl *gomock.Controller){},
		},
		{
			name: "異常(Name重複)",
			wantErr: true,
		},
		{
			name: "異常(Name重複(DBの制約で判明))",
			wantErr: true,
		},
		{
			name: "異常(DB異常)",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &metaUseCase{
				repository: tt.fields.repository,
			}
			got, err := m.CreateKey(tt.args.ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateKey() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_metaUseCase_FetchKeyByID(t *testing.T) {
	type fields struct {
		repository repository.Meta
	}
	type args struct {
		ctx context.Context
		ID  uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.MetaKey
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &metaUseCase{
				repository: tt.fields.repository,
			}
			got, err := m.FetchKeyByID(tt.args.ctx, tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchKeyByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchKeyByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_metaUseCase_RemoveKey(t *testing.T) {
	type fields struct {
		repository repository.Meta
	}
	type args struct {
		ctx context.Context
		ID  uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.MetaKey
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &metaUseCase{
				repository: tt.fields.repository,
			}
			got, err := m.RemoveKey(tt.args.ctx, tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveKey() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_metaUseCase_UpdateKey(t *testing.T) {
	type fields struct {
		repository repository.Meta
	}
	type args struct {
		ctx  context.Context
		ID   uint
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.MetaKey
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &metaUseCase{
				repository: tt.fields.repository,
			}
			got, err := m.UpdateKey(tt.args.ctx, tt.args.ID, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateKey() got = %v, want %v", got, tt.want)
			}
		})
	}
}
