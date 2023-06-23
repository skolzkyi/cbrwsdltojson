package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	//"time"

	"github.com/skolzkyi/cbrwsdltojson/internal/app"
	"github.com/skolzkyi/cbrwsdltojson/internal/logger"

	internalhttp "github.com/skolzkyi/cbrwsdltojson/internal/server/http"
	soapreq "github.com/skolzkyi/cbrwsdltojson/internal/soapreq"
)

var configFilePath string

func init() {
	flag.StringVar(&configFilePath, "config", "./configs/", "Path to config.env")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	config := NewConfig()
	err := config.Init(configFilePath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config: ", config)
	log, err := logger.New(config.Logger.Level)
	if err != nil {
		fmt.Println(err)
	}
	log.Info("servAddr: " + config.GetAddress())

	soapReqSender := soapreq.NewSoapRequestSender()
	soapReqSender.Init(&config)

	cbrwsdltojson := app.New(log, &config, &soapReqSender)

	server := internalhttp.NewServer(log, cbrwsdltojson, &config)

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), config.GetServerShutdownTimeout())
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			log.Fatal("failed to stop http server: " + err.Error())
		}
	}()

	if err := server.Start(ctx); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Error("failed to start http server: " + err.Error())
		cancel()
		os.Exit(1) //nolint:gocritic
	}
}
