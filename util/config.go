package util

import (
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/hibiken/asynq"
	"github.com/joho/godotenv"
)

type Config struct {
	DBDriver   string `mapstructure:"DB_DRIVER"`
	DBSource   string `mapstructure:"DB_SOURCE"`
	RedisUrl   string `mapstructure:"REDIS_URL"`
	SIGNINGKEY string `mapstructure:"SIGNING_KEY"`
}

func LoadEnvConfig() (config Config) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Cannot load env:", err)
	}

	config.DBDriver = os.Getenv("DB_DRIVER")
	config.DBSource = os.Getenv("DB_SOURCE")
	config.SIGNINGKEY = os.Getenv("TOKEN_SYMMETRIC_KEY")
	config.RedisUrl = os.Getenv("REDIS_URL")

	return config
}

func CorsWhiteList() string {
	return os.Getenv("CORS_WHITELIST")
}

func parseRedisURI(redisURL string) (asynq.RedisConnOpt, error) {
	o := asynq.RedisClientOpt{Network: "tcp"}

	u, err := url.Parse(redisURL)
	if err != nil {
		return nil, err
	}

	if u.Scheme != "redis" && u.Scheme != "rediss" {
		return nil, errors.New("invalid redis URL scheme: " + u.Scheme)
	}

	if u.User != nil {
		if p, ok := u.User.Password(); ok {
			o.Password = p
		}
	}

	if len(u.Query()) > 0 {
		return nil, errors.New("no options supported")
	}

	h, p, err := net.SplitHostPort(u.Host)
	if err != nil {
		h = u.Host
	}
	if h == "" {
		h = "localhost"
	}
	if p == "" {
		p = "6379"
	}
	o.Addr = net.JoinHostPort(h, p)

	f := strings.FieldsFunc(u.Path, func(r rune) bool {
		return r == '/'
	})
	switch len(f) {
	case 0:
		o.DB = 0
	case 1:
		if o.DB, err = strconv.Atoi(f[0]); err != nil {
			return nil, fmt.Errorf("invalid redis database number: %q", f[0])
		}
	default:
		return nil, errors.New("invalid redis URL path: " + u.Path)
	}

	if u.Scheme == "rediss" {
		o.TLSConfig = &tls.Config{ServerName: h}
	}

	o.Username = u.User.Username()

	return o, nil
}

func GetRedisConn() (*asynq.RedisClientOpt, error) {
	uri, err := parseRedisURI(os.Getenv("REDIS_URL"))

	config, ok := uri.(asynq.RedisClientOpt)
	if !ok {
		return nil, err
	}

	return &config, nil
}
