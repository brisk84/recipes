package app

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"recipes/internal/config"
	"recipes/internal/handler"
	"recipes/internal/server"
	"recipes/internal/storage"
	"recipes/internal/usecase"
	"recipes/pkg/logger"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/redis/go-redis/v9"
)

type App struct {
	srv *server.Server
	lg  logger.Logger
}

func New(lg logger.Logger, cfg config.Config) (*App, error) {
	stor, err := storage.New(lg, cfg)
	if err != nil {
		return nil, fmt.Errorf("storage: %w", err)
	}

	mc, err := minio.New(cfg.MinioAddr, &minio.Options{
		Creds: credentials.NewStaticV4(cfg.MinioLogin, cfg.MinioPass, ""), Secure: false,
	})
	if err != nil {
		return nil, fmt.Errorf("minio: %w", err)
	}

	ctx := context.TODO()
	bucketName := "store"
	err = mc.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: "local"})
	if err != nil {
		exists, errBucketExists := mc.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	f, err := os.Open("/tmp/test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	fmt.Println(rd.Size())
	// i, err := mc.PutObject(context.TODO(), "store", "test", rd, int64(rd.Size()), minio.PutObjectOptions{ContentType: "image/png"})
	// i, err := mc.PutObject(context.TODO(), bucketName, "test.txt", rd, int64(rd.Size()), minio.PutObjectOptions{ContentType: "plain/text"})
	i, err := mc.FPutObject(context.TODO(), bucketName, "test.txt", "/tmp/test.txt", minio.PutObjectOptions{ContentType: "plain/text"})

	fmt.Println(i, err)
	fmt.Println()
	fmt.Println()
	fmt.Println()
	if err != nil {
		panic(err)
	}

	// s, err := rd.ReadString('\n')
	// fmt.Println(s, err)

	uc, err := usecase.New(lg, cfg, stor)
	if err != nil {
		return nil, fmt.Errorf("usecase: %w", err)
	}

	rcli := redis.NewClient(&redis.Options{
		Addr: cfg.RedisAddr,
	})

	h := handler.New(lg, uc, rcli)
	srv := server.New(lg, cfg.AppAddr, h)

	return &App{
		srv: srv,
		lg:  lg,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	return a.srv.Start(ctx)
}
