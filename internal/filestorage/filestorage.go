package filestorage

import (
	"context"
	"fmt"
	"io"
	"recipes/internal/config"
	"recipes/pkg/logger"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type filestorage struct {
	lg    logger.Logger
	login string
	pass  string
	url   string
	mc    *minio.Client
}

func New(lg logger.Logger, cfg config.Config) (*filestorage, error) {
	fs := filestorage{
		lg:    lg,
		login: cfg.MinioLogin,
		pass:  cfg.MinioPass,
		url:   cfg.MinioAddr,
	}
	if err := fs.Init(); err != nil {
		return nil, err
	}
	return &fs, nil
}

func (f *filestorage) Init() error {
	var err error
	f.mc, err = minio.New(f.url, &minio.Options{
		Creds: credentials.NewStaticV4(f.login, f.pass, ""), Secure: false,
	})
	if err != nil {
		return fmt.Errorf("minio.New: %w", err)
	}

	ctx := context.Background()
	bucketName := "store"
	err = f.mc.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: "local"})
	if err != nil {
		exists, errBucketExists := f.mc.BucketExists(ctx, bucketName)
		if !(errBucketExists == nil && exists) {
			return fmt.Errorf("f.mc.MakeBucket: %w", err)
		}
	}
	return nil
}

func (f *filestorage) Upload(ctx context.Context, fileName string, fileSize int64,
	reader io.Reader) error {

	_, err := f.mc.PutObject(ctx, "store", fileName, reader, fileSize,
		minio.PutObjectOptions{ContentType: "image/png"})
	if err != nil {
		return fmt.Errorf("f.mc.PutObject: %w", err)
	}
	return nil
}

func (f *filestorage) Download(ctx context.Context, fileName string) (io.Reader, error) {
	obj, err := f.mc.GetObject(ctx, "store", fileName, minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("f.mc.PutObject: %w", err)
	}
	return obj, nil
}
