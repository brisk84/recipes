package usecase

import (
	"recipes/pkg/logger"
	"testing"
)

func TestHashPassword(t *testing.T) {
	lg, _ := logger.New(true)
	stor := new(storageMock)
	fs := new(filestorageMock)

	type fields struct {
		logger      logger.Logger
		storage     storage
		filestorage filestorage
	}
	type args struct {
		password string
		hash     string
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "positive 1",
			fields: fields{
				logger:      lg,
				storage:     stor,
				filestorage: fs,
			},
			args: args{
				password: "pass01",
				hash:     "$2a$10$IjY1vdDM41oSNcoPwN8ipOlt3bg0J4cn2afz5RJ8RxEI5CglF7/V.",
			},
			want: true,
		},
		{
			name: "positive 2",
			fields: fields{
				logger:  lg,
				storage: stor,
			},
			args: args{
				password: "pass02",
				hash:     "$2a$10$Lb4GeaBxy96Di2XYi00DaOhWHJQCniecwglJY5WZFzHGPyYykNctG",
			},
			want: true,
		},
		{
			name: "negative 1",
			fields: fields{
				logger:  lg,
				storage: stor,
			},
			args: args{
				password: "pass04",
				hash:     "$2a$10$oSI.Pisk4JMCE./uj/ndfOFGPp94rOcCCArXHiigxz53T6CtvSS.C_",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &usecase{
				lg:   tt.fields.logger,
				stor: tt.fields.storage,
			}
			got := u.checkPasswordHash(tt.args.password, tt.args.hash)
			if got != tt.want {
				t.Errorf("usecase.checkPasswordHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
