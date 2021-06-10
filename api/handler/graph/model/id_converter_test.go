package model

import "testing"

func TestIDtoKeyNameAndInternalID(t *testing.T) {
	type args struct {
		ID string
	}
	tests := []struct {
		name    string
		args    args
		want    uint
		want1   IDType
		wantErr bool
	}{
		{
			name:    "正常：ItemTemplate",
			args:    args{ID: "ItemTemplate:1"},
			want:    uint(1),
			want1:   IDTypeItemTemplate,
			wantErr: false,
		},
		{
			name:    "異常：ItemTemplate：数値以外のInternalID",
			args:    args{ID: "ItemTemplate:hoge"},
			want:    uint(0),
			want1:   IDTypeUnknown,
			wantErr: true,
		},
		{
			name:    "異常：存在しないKeyType",
			args:    args{ID: "Hoge:10"},
			want:    uint(0),
			want1:   IDTypeUnknown,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := IDtoKeyNameAndInternalID(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("IDtoKeyNameAndInternalID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IDtoKeyNameAndInternalID() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("IDtoKeyNameAndInternalID() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestKeyType_ToExternalID(t *testing.T) {
	type args struct {
		ID uint
	}
	tests := []struct {
		name string
		key  IDType
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.key.ToExternalID(tt.args.ID); got != tt.want {
				t.Errorf("ToExternalID() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func TestKeyType_ToInternalID(t *testing.T) {
//	type args struct {
//		ID string
//	}
//	tests := []struct {
//		name string
//		key  IDType
//		args args
//		want uint
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.key.ToInternalID(tt.args.ID); got != tt.want {
//				t.Errorf("ToInternalID() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
