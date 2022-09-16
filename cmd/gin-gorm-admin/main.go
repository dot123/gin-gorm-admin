package main

import (
	"context"
	"github.com/dot123/gin-gorm-admin/pkg/logger"
	"github.com/urfave/cli/v2"
	"os"
)

// https://github.com/swaggo/swag/blob/master/README_zh-CN.md

// @title          gin-gorm-admin API
// @version        1.0
// @description    This is a game management background. you can use the api key `ApiKeyAuth` to test the authorization filters.
// @termsOfService https://github.com

// @contact.name  conjurer
// @contact.url   https:/github.com/dot123
// @contact.email conjurer888888@gmail.com

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @host     127.0.0.1:8000
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in                         header
// @name                       Authorization

// Usage: go build -ldflags "-X main.VERSION=x.x.x"
var VERSION = "1.0.0"

func main() {
	ctx := logger.NewTagContext(context.Background(), "__main__")

	app := cli.NewApp()
	app.Name = "gin-gorm-admin"
	app.Version = VERSION
	app.Usage = "Web scaffolding based on GIN + GORM + WIRE."
	app.Commands = []*cli.Command{
		newWebCmd(ctx),
	}
	err := app.Run(os.Args)
	if err != nil {
		logger.WithContext(ctx).Errorf(err.Error())
	}
}

func newWebCmd(ctx context.Context) *cli.Command {
	return &cli.Command{
		Name:  "web",
		Usage: "Run http server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "conf",
				Aliases:  []string{"c"},
				Usage:    "App configuration file(.json,.yaml,.toml)",
				Required: true,
			},
			&cli.StringFlag{
				Name:    "localPath",
				Aliases: []string{"p"},
				Usage:   "Local file directory",
			},
		},
		Action: func(c *cli.Context) error {
			return Run(ctx,
				SetConfigFile(c.String("conf")),
				SetLocalPathDir(c.String("localPath")),
				SetVersion(VERSION))
		},
	}
}
