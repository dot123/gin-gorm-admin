package config

import (
	"encoding/json"
	"fmt"
	"github.com/koding/multiconfig"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	C    = new(Config)
	once sync.Once
)

// Load config file (toml/json/yaml)
func MustLoad(fpaths ...string) {
	once.Do(func() {
		loaders := []multiconfig.Loader{
			&multiconfig.TagLoader{},
			&multiconfig.EnvironmentLoader{},
		}

		for _, fpath := range fpaths {
			if strings.HasSuffix(fpath, "toml") {
				loaders = append(loaders, &multiconfig.TOMLLoader{Path: fpath})
			}
			if strings.HasSuffix(fpath, "json") {
				loaders = append(loaders, &multiconfig.JSONLoader{Path: fpath})
			}
			if strings.HasSuffix(fpath, "yaml") {
				loaders = append(loaders, &multiconfig.YAMLLoader{Path: fpath})
			}
		}

		m := multiconfig.DefaultLoader{
			Loader:    multiconfig.MultiLoader(loaders...),
			Validator: multiconfig.MultiValidator(&multiconfig.RequiredValidator{}),
		}
		m.MustLoad(C)
	})
}

func PrintWithJSON() {
	if C.PrintConfig {
		b, err := json.MarshalIndent(C, "", " ")
		if err != nil {
			os.Stdout.WriteString("[CONFIG] JSON marshal error: " + err.Error())
			return
		}
		os.Stdout.WriteString(string(b) + "\n")
	}
}

type Config struct {
	RunMode     string
	LocalPath   string
	Swagger     bool
	PrintConfig bool
	HTTP        HTTP
	Log         Log
	LogGormHook LogGormHook
	JWTAuth     JWTAuth
	Monitor     Monitor
	RateLimiter RateLimiter
	CORS        CORS
	GZIP        GZIP
	Redis       Redis
	Gorm        Gorm
	MySQL       MySQL
	Sqlite3     Sqlite3
}

func (c *Config) IsDebugMode() bool {
	return c.RunMode == "debug"
}

type LogHook string

func (h LogHook) IsGorm() bool {
	return h == "gorm"
}

type Log struct {
	Level         int
	Format        string
	Output        string
	OutputFile    string
	EnableHook    bool
	HookLevels    []string
	Hook          LogHook
	HookMaxThread int
	HookMaxBuffer int
	RotationCount int
	RotationTime  int
}

type LogGormHook struct {
	DBType       string
	MaxLifetime  time.Duration
	MaxOpenConns int
	MaxIdleConns int
	Table        string
}

type JWTAuth struct {
	Realm       string
	Key         string
	Expired     time.Duration
	IdentityKey string
}

type HTTP struct {
	Host               string
	Port               int
	CertFile           string
	KeyFile            string
	ShutdownTimeout    int
	MaxContentLength   int64
	MaxReqLoggerLength int `default:"1024"`
	MaxResLoggerLength int
}

type Monitor struct {
	Enable    bool
	Addr      string
	ConfigDir string
}

type RateLimiter struct {
	Enable  bool
	Count   int64
	RedisDB int
}

type CORS struct {
	Enable           bool
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	AllowCredentials bool
	MaxAge           time.Duration
}

type GZIP struct {
	Enable             bool
	ExcludedExtentions []string
	ExcludedPaths      []string
}

type Redis struct {
	Addr     string
	Password string
}

type Gorm struct {
	Debug             bool
	DBType            string
	MaxLifetime       time.Duration
	MaxOpenConns      int
	MaxIdleConns      int
	TablePrefix       string
	EnableAutoMigrate bool
}

type MySQL struct {
	Host       string
	Port       int
	User       string
	Password   string
	DBName     string
	Parameters string
}

func (a MySQL) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		a.User, a.Password, a.Host, a.Port, a.DBName, a.Parameters)
}

type Sqlite3 struct {
	Path string
}

func (a Sqlite3) DSN() string {
	return a.Path
}
