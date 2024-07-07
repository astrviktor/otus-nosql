package main

import (
	"fmt"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"project/internal/config"
	"project/internal/logger"
	"project/internal/middleware"
	"project/internal/service"
	"syscall"
)

func main() {
	if len(os.Args) > 1 {
		config.PrintUsage(config.ServicePrefix)
		return
	}

	log, err := logger.NewLogger("debug", true)
	if err != nil {
		fmt.Println("fail to create logger")
		os.Exit(1)
	}
	log = log.Named(config.ServicePrefix)

	cfg, err := config.ReadConfig(config.ServicePrefix)
	if err != nil {
		log.Error("fail to read config", zap.Error(err))
		os.Exit(1)
	}

	defer func() {
		_ = log.Sync()
		os.Exit(1)
	}()

	h, err := service.NewHandler(log, cfg)
	if err != nil {
		log.Error("fail to create handler", zap.Error(err))
		os.Exit(1)
	}

	h.Run()

	r := router.New()
	r.POST("/test/redis/string", middleware.Middleware(log, h.TestRedisString))
	r.POST("/test/redis/hset", middleware.Middleware(log, h.TestRedisHset))
	r.POST("/test/redis/zset", middleware.Middleware(log, h.TestRedisZset))
	r.POST("/test/redis/list", middleware.Middleware(log, h.TestRedisList))

	server := fasthttp.Server{
		Handler: r.Handler,
	}

	go func() {
		addr := fmt.Sprintf("%s:%s", cfg.Service.Host, cfg.Service.Port)

		log.Info("start server", zap.String("server addr", addr))
		err = server.ListenAndServe(addr)
		if err != nil {
			log.Error("fail to listen", zap.Error(err))
			os.Exit(1)
		}
	}()

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	<-exit
	log.Info("caught os signal to stop server")

	err = server.Shutdown()
	if err != nil {
		log.Error("fail to shutdown service", zap.Error(err))
	} else {
		log.Info("server was successfully stopped")
	}

	h.Stop()
}
