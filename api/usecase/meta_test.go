package usecase

import (
	"ao2-y/data-tag-manager/domain/model"
	"ao2-y/data-tag-manager/domain/repository"
	repository2 "ao2-y/data-tag-manager/domain/repository/mock"
	"ao2-y/data-tag-manager/logger"
	"context"
	"fmt"
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
			name: "正常",
			args: args{repository: repo},
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

type metaUseCasePrepareFunc func(repo *repository2.MockMeta)

func Test_metaUseCase_CreateKey(t *testing.T) {
	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name        string
		prepareFunc metaUseCasePrepareFunc
		args        args
		want        *model.MetaKey
		wantErr     bool
	}{
		{
			name:    "正常",
			wantErr: false,
			prepareFunc: func(repo *repository2.MockMeta) {
				repo.EXPECT().FetchByName(gomock.Any(), gomock.Any()).Times(1)
				repo.EXPECT().CreateKey(gomock.Any(), "test1")
			},
			args: args{
				ctx:  context.Background(),
				name: "test1",
			},
		},
		{
			name:    "異常(Name重複)",
			wantErr: true,
			prepareFunc: func(repo *repository2.MockMeta) {
				repo.EXPECT().FetchByName(gomock.Any(), gomock.Any()).Return(&model.MetaKey{ID: uint(1)}, nil).Times(1)
			},
			args: args{
				ctx:  context.Background(),
				name: "test2",
			},
		},
		{
			name:    "異常(Name重複(DBの制約で判明))",
			wantErr: true,
			prepareFunc: func(repo *repository2.MockMeta) {
				repo.EXPECT().FetchByName(gomock.Any(), gomock.Any()).Return(nil, nil).Times(1)
				repo.EXPECT().CreateKey(gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("hoge")).Times(1)
			},
			args: args{
				ctx:  context.Background(),
				name: "test3",
			},
		},
		{
			name:    "異常(DB異常)",
			wantErr: true,
			prepareFunc: func(repo *repository2.MockMeta) {
				repo.EXPECT().FetchByName(gomock.Any(), gomock.Any()).Return(nil, nil).Times(1)
				repo.EXPECT().CreateKey(gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("hoge")).Times(1)
			},
			args: args{
				ctx:  context.Background(),
				name: "test4",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			repo := repository2.NewMockMeta(ctrl)
			tt.prepareFunc(repo)
			m := &metaUseCase{
				repository: repo,
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
	logger.InitApplicationLogger()
	type args struct {
		ctx context.Context
		ID  uint
	}
	tests := []struct {
		name        string
		prepareFunc metaUseCasePrepareFunc
		args        args
		want        *model.MetaKey
		wantErr     bool
	}{
		{
			name: "正常",
			prepareFunc: func(repo *repository2.MockMeta) {
				repo.EXPECT().FetchByID(gomock.Any(), uint(1)).Return(&model.MetaKey{
					ID:   1,
					Name: "test1",
				}, nil).Times(1)
			},
			args: args{
				ctx: context.Background(),
				ID:  1,
			},
			want: &model.MetaKey{
				ID:   1,
				Name: "test1",
			},
			wantErr: false,
		},
		{
			name: "正常：存在しないID",
			prepareFunc: func(repo *repository2.MockMeta) {
				repo.EXPECT().FetchByID(gomock.Any(), uint(2)).Return(nil, repository.NewOperationError(repository.ErrNotFound, nil)).Times(1)
			},
			args: args{
				ctx: context.Background(),
				ID:  2,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "異常：DBエラー",
			prepareFunc: func(repo *repository2.MockMeta) {
				repo.EXPECT().FetchByID(gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("hoge")).Times(1)
			},
			args: args{
				ctx: context.Background(),
				ID:  3,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			metaRepo := repository2.NewMockMeta(ctrl)
			tt.prepareFunc(metaRepo)
			m := &metaUseCase{
				repository: metaRepo,
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
	logger.InitApplicationLogger()
	type args struct {
		ctx context.Context
		ID  uint
	}
	tests := []struct {
		name        string
		prepareFunc metaUseCasePrepareFunc
		args        args
		want        *model.MetaKey
		wantErr     bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			metaRepo := repository2.NewMockMeta(ctrl)
			tt.prepareFunc(metaRepo)
			m := &metaUseCase{
				repository: metaRepo,
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
