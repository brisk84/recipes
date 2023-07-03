package logger

import "go.uber.org/zap"

type Logger interface {
	Errorln(args ...interface{})
	Infoln(args ...interface{})
	Fatalln(args ...interface{})
}

func New(debug bool) (*zap.SugaredLogger, error) {
	var cfg zap.Config
	if debug {
		cfg = zap.NewDevelopmentConfig()
	} else {
		cfg = zap.NewProductionConfig()
	}

	cfg.Encoding = "json"
	lg, err := cfg.Build()
	if err != nil {
		return nil, err
	}
	return lg.Sugar(), nil
}
