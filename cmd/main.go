package main

import (
	"context"
	"fmt"
	"github.com/brnck/jiratron/cmd/app"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

func main() {
	ctx := SignalHandlerContext()
	cmd := app.New(ctx)

	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func SignalHandlerContext() context.Context {
	ctx, cancel := context.WithCancel(context.TODO())
	ch := make(chan os.Signal, 2)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	go func() {
		sig := <-ch

		cancel()

		for i := 0; i < 3; i++ {
			logrus.Warnf("received signal %s, shutting down...", sig)
			sig = <-ch
		}

		logrus.Warnf("received signal %s, force closing", sig)

		os.Exit(1)
	}()

	return ctx
}
