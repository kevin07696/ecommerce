package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	sessionRedis "github.com/go-session/redis"
	"github.com/go-session/session"
	"github.com/joho/godotenv"
	"github.com/kevin07696/ecommerce/adapters/driven/email"
	"github.com/kevin07696/ecommerce/adapters/driven/mongodb"
	ur "github.com/kevin07696/ecommerce/adapters/driven/mongodb/user"
	cache "github.com/kevin07696/ecommerce/adapters/driven/redis"
	sess "github.com/kevin07696/ecommerce/adapters/driven/session"
	"github.com/kevin07696/ecommerce/adapters/driven/slogger"
	"github.com/kevin07696/ecommerce/adapters/driving"
	"github.com/kevin07696/ecommerce/adapters/driving/http/auth"
	"github.com/kevin07696/ecommerce/domain/auth/models"
	up "github.com/kevin07696/ecommerce/domain/auth/port"
	us "github.com/kevin07696/ecommerce/domain/auth/services"
	"github.com/redis/go-redis/v9"
	"github.com/wneessen/go-mail"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	env := map[string]string{
		"APP_ENV":    "",
		"REDIS_HOST": "",
		"REDIS_PORT": "",
		"EMAIL_PORT": "",
		"EMAIL_HOST": "",
		"EMAIL_PASS": "",
		"EMAIL_USER": "",
		"EMAIL_AUTH": "",
		"MONGO_URI":  "",
		"MONGO_POOL": "",
	}
	for k := range env {
		env[k] = os.Getenv(k)
		if env[k] == "" {
			log.Printf("Environment variable %s is missing", k)
		}
	}

	router := http.NewServeMux()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var programLevel = new(slog.LevelVar)
	h := &slogger.ContextHandler{
		Handler: slog.NewJSONHandler(os.Stderr,
			&slog.HandlerOptions{AddSource: true, Level: programLevel}),
	}
	slog.SetDefault(slog.New(h))
	if env["APP_ENV"] == "prod" {
		programLevel.Set(slog.LevelError)
	} else {
		programLevel.Set(slog.LevelDebug)
	}

	mongoPool, err := strconv.ParseUint(env["MONGO_POOL"], 10, 16)
	if err != nil {
		log.Panic("failed to parse mongo connection pool size: " + err.Error())
	}
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	mongoOpts := []*options.ClientOptions{
		options.Client().ApplyURI(env["MONGO_URI"]).SetServerAPIOptions(serverAPI),
		options.Client().SetMaxPoolSize(mongoPool),
	}

	client, err := mongodb.ConnectMongoDb(ctx, mongoOpts...)
	if err != nil {
		log.Panic(err)
	}
	defer mongodb.DisconnectDB(ctx, client)

	var userRepo up.IRepository
	userRepo, err = ur.NewUserRepository(ctx, client, "db", "user")
	if err != nil {
		log.Panic(err)
	}

	redisAddr := fmt.Sprintf("%s:%s", env["REDIS_HOST"], env["REDIS_PORT"])

	redisOpts := redis.Options{
		Addr: redisAddr,
		DB:   0,
	}
	var userCache up.ICache = cache.NewRedisAdapter(&redisOpts)

	var sessionOpts sessionRedis.Options = sessionRedis.Options{
		Addr: redisAddr,
		DB:   0,
	}
	var store = sess.NewRedisStore(&sessionOpts)
	var sessionManager up.ISessionManager = sess.NewSessionManager(session.SetStore(store),
		session.SetCookieLifeTime(30), session.SetCookieName("auth-session"),
		session.SetSecure(true), session.SetEnableSetCookie(true), session.SetExpired(30))

	emailPort, err := strconv.ParseUint(env["EMAIL_PORT"], 10, 16)
	if err != nil {
		log.Panic("failed to parse email port: " + err.Error())
	}
	emailOpts := []mail.Option{
		mail.WithPort(int(emailPort)),
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(env["EMAIL_USER"]),
		mail.WithPassword(env["EMAIL_PASS"]),
	}
	emailClient, err := email.NewEmailClient(env["EMAIL_HOST"], emailOpts...)
	if err != nil {
		log.Panic(err)
	}

	authEmailer, err := email.NewClientWrapper(emailClient, env["EMAIL_AUTH"])
	if err != nil {
		log.Panic(err)
	}

	userAPI := us.NewService(userRepo, sessionManager, userCache, authEmailer, models.Models{})
	auth.Handle(router, userAPI)

	app := driving.NewApp(8080, router)
	app.Run()
}
