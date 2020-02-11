package main

import (
	"fmt"
	"os"
	"os/signal"
	"simple-redirect-service/context"
	"syscall"
)

// app entrypoint
func main() {
	appCtx := context.Build()
	signal.Notify(appCtx.Signal, os.Interrupt, syscall.SIGTERM)
	fmt.Println(appCtx.App.Run())
}
