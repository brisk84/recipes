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
	user02 := domain.User{Login: "user02", Password: "$2a$10$IjY1vdDM41oSNcoPwN8ipOlt3bg0J4cn2afz5RJ8RxEI5CglF7/V."}

	stor := new(storageMock)
	stor.On("ReadUser", context.Background(), "user01").Return(user01)

	type fields struct {
		logger  logger.Logger
		storage storage
	}
	type args struct {
		ctx  context.Context
		user domain.User
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "positive 1",
			fields: fields{
				logger:  lg,
				storage: stor,
			},
			args: args{
				ctx:  context.Background(),
				user: user01,
			},
			want:    "http://localhost/0",
			wantErr: false,
		},
		{
			name: "positive 2",
			fields: fields{
				logger:  lg,
				storage: stor,
			},
			args: args{
				ctx: context.TODO(),
				url: "http://www.very-long-url.com/2",
			},
			want:    "http://localhost/1",
			wantErr: false,
		},
		{
			name: "negative 1",
			fields: fields{
				logger:  lg,
				storage: stor,
			},
			args: args{
				ctx: context.TODO(),
				url: "",
			},
			want:    "",
			wantErr: true,
		},
	}
}
