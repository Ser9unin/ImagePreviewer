package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Ser9unin/ImagePreviewer/internal/app"
	"github.com/Ser9unin/ImagePreviewer/internal/cache"
	"github.com/Ser9unin/ImagePreviewer/internal/config"
	"github.com/Ser9unin/ImagePreviewer/internal/logger"
	"github.com/Ser9unin/ImagePreviewer/internal/server"
	"golang.org/x/sync/errgroup"
)

func main() {
	logger := logger.NewLogger()
	config := config.New()
	cache := cache.NewCache(config.Cache)
	app := app.New(cache, logger)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		c := make(chan os.Signal, 1) // we need to reserve to buffer size 1, so the notifier are not blocked
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		<-c
		cancel()
	}()

	server := server.NewServer(config, app, logger)

	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return server.Run()
	})
	g.Go(func() error {
		<-gCtx.Done()
		return server.Stop(context.Background())
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("exit reason: %s \n", err)
	}
}
