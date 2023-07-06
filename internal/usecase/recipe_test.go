package usecase

import (
	"context"
	"recipes/domain"
	"recipes/pkg/logger"
	"testing"
)

func Test_usecase_SignIn(t *testing.T) {
	lg, err := logger.New(true)
	if err != nil {
		t.Errorf("Error creating logger %v", err)
	}

	user01 := domain.User{Login: "user01", Password: "$2a$10$IjY1vdDM41oSNcoPwN8ipOlt3bg0J4cn2afz5RJ8RxEI5CglF7/V."}
	user02 := domain.User{Login: "user02", Password: "$2a$10$Lb4GeaBxy96Di2XYi00DaOhWHJQCniecwglJY5WZFzHGPyYykNctG"}
	user03 := domain.User{Login: "user03", Password: "$2a$10$oSI.Pisk4JMCE./uj/ndfOFGPp94rOcCCArXHiigxz53T6CtvSS.C"}

	stor := new(storageMock)
	stor.On("ReadUser", context.Background(), "user01").Return(user01, nil)
	stor.On("ReadUser", context.Background(), "user02").Return(user02, nil)
	stor.On("ReadUser", context.Background(), "user03").Return(user03, nil)
	fs := new(filestorageMock)

	type fields struct {
		logger      logger.Logger
		storage     storage
		filestorage filestorage
	}
	type args struct {
		ctx  context.Context
		user domain.User
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "positive 1",
			fields: fields{
				logger:      lg,
				storage:     stor,
				filestorage: fs,
			},
			args: args{
				ctx:  context.Background(),
				user: domain.User{Login: "user01", Password: "pass01"},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "positive 2",
			fields: fields{
				logger:      lg,
				storage:     stor,
				filestorage: fs,
			},
			args: args{
				ctx:  context.Background(),
				user: domain.User{Login: "user02", Password: "pass02"},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "negative 1",
			fields: fields{
				logger:      lg,
				storage:     stor,
				filestorage: fs,
			},
			args: args{
				ctx:  context.Background(),
				user: domain.User{Login: "user03", Password: "pass04"},
			},
			want:    false,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &usecase{
				lg:   tt.fields.logger,
				stor: tt.fields.storage,
			}
			_, got2, err := u.SignIn(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.SignIn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got2 != tt.want {
				t.Errorf("usecase.SignIn() = %v, want %v", got2, tt.want)
			}
		})
	}
}
