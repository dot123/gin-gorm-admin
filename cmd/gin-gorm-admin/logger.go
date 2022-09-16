package main

import (
	"github.com/dot123/gin-gorm-admin/internal/config"
	"github.com/dot123/gin-gorm-admin/pkg/logger"
	loggerhook "github.com/dot123/gin-gorm-admin/pkg/logger/hook"
	loggergormhook "github.com/dot123/gin-gorm-admin/pkg/logger/hook/gorm"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"os"
	"path/filepath"
	"time"
)

func InitLogger() (func(), error) {
	c := config.C.Log
	logger.SetLevel(logger.Level(c.Level))
	logger.SetFormatter(c.Format)

	var file *rotatelogs.RotateLogs
	if c.Output != "" {
		switch c.Output {
		case "stdout":
			logger.SetOutput(os.Stdout)
		case "stderr":
			logger.SetOutput(os.Stderr)
		case "file":
			if name := c.OutputFile; name != "" {
				_ = os.MkdirAll(filepath.Dir(name), 0777)

				f, err := rotatelogs.New(name+".%Y-%m-%d",
					rotatelogs.WithLinkName(name),
					rotatelogs.WithRotationTime(time.Duration(c.RotationTime)*time.Hour),
					rotatelogs.WithRotationCount(uint(c.RotationCount)))
				if err != nil {
					return nil, err
				}

				logger.SetOutput(f)
				file = f
			}
		}
	}

	var hook *loggerhook.Hook
	if c.EnableHook {
		var hookLevels []logger.Level
		for _, lvl := range c.HookLevels {
			plvl, err := logger.ParseLevel(lvl)
			if err != nil {
				return nil, err
			}
			hookLevels = append(hookLevels, plvl)
		}

		switch {
		case c.Hook.IsGorm():
			db, err := NewGormDB()
			if err != nil {
				return nil, err
			}

			h := loggerhook.New(loggergormhook.New(db),
				loggerhook.SetMaxWorkers(c.HookMaxThread),
				loggerhook.SetMaxQueues(c.HookMaxBuffer),
				loggerhook.SetLevels(hookLevels...),
			)
			logger.AddHook(h)
			hook = h
		}
	}

	return func() {
		if file != nil {
			file.Close()
		}

		if hook != nil {
			hook.Flush()
		}
	}, nil
}
