package main

import (
	"github.com/dot123/gin-gorm-admin/internal/config"
	"github.com/dot123/gin-gorm-admin/pkg/logger"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"os"
	"path/filepath"
	"time"
)

func InitLogger() (func(), error) {
	c := config.C.Log
	logger.SetLevel(logger.Level(c.Level))
	logger.SetFormatter(c.Format)

	var rl *rotatelogs.RotateLogs
	if c.Output != "" {
		switch c.Output {
		case "stdout":
			logger.SetOutput(os.Stdout)
		case "stderr":
			logger.SetOutput(os.Stderr)
		case "file":
			if name := c.OutputFile; name != "" {
				_ = os.MkdirAll(filepath.Dir(name), 0777)

				r, err := rotatelogs.New(name+".%Y-%m-%d",
					rotatelogs.WithLinkName(name),
					rotatelogs.WithRotationTime(time.Duration(c.RotationTime)*time.Hour),
					rotatelogs.WithRotationCount(uint(c.RotationCount)))
				if err != nil {
					return nil, err
				}

				logger.SetOutput(r)
				rl = r
			}
		}
	}

	return func() {
		if rl != nil {
			rl.Close()
		}
	}, nil
}
