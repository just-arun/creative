package v1

import (
	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/just-arun/creative/api/v1/hello"
	conf "github.com/just-arun/creative/config"
	"github.com/just-arun/creative/models"
)

func ApiV1(env string) func(c chi.Router) {

	envFileName := ".env." + env
	var config models.Config
	conf.GetEnv(envFileName, "yml", ".", &config)

	fmt.Println(config.App.Server.Port)

	var ctx models.HandlerCtx

	userSession := conf.GetRedisDbClient(config.App.Redis.Host, config.App.Redis.Password, conf.RedisDBTypeSession)
	otpSession := conf.GetRedisDbClient(config.App.Redis.Host, config.App.Redis.Password, conf.RedisDBTypeOTP)

	ctx.UserSessionDB = userSession
	ctx.OtpSessionDB = otpSession

	return func(r chi.Router) {
		r.Route("/hello", hello.Route(&ctx))
	}
}
