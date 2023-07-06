package app

import (
	"context"
	"fmt"
	"recipes/internal/config"
	"recipes/internal/filestorage"
	"recipes/internal/handler"
	"recipes/internal/server"
	"recipes/internal/storage"
	"recipes/internal/usecase"
	"recipes/pkg/logger"

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

	fs, err := filestorage.New(lg, cfg)
	if err != nil {
		return nil, fmt.Errorf("filestorage: %w", err)
	}

	// f, err := os.Open("/tmp/test.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
	// rd := bufio.NewReader(f)
	// fmt.Println(rd.Size())
	// // i, err := mc.PutObject(context.TODO(), "store", "test", rd, int64(rd.Size()), minio.PutObjectOptions{ContentType: "image/png"})
	// // i, err := mc.PutObject(context.TODO(), bucketName, "test.txt", rd, int64(rd.Size()), minio.PutObjectOptions{ContentType: "plain/text"})
	// i, err := mc.FPutObject(context.TODO(), bucketName, "test.txt", "/tmp/test.txt", minio.PutObjectOptions{ContentType: "plain/text"})

	// fmt.Println(i, err)
	// fmt.Println()
	// fmt.Println()
	// fmt.Println()
	// if err != nil {
	// 	panic(err)
	// }

	// s, err := rd.ReadString('\n')
	// fmt.Println(s, err)

	uc, err := usecase.New(lg, cfg, stor, fs)
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
