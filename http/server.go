package http

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"
)

// Server encapsulates all HTTP exposed functionality of app
type Server interface {
	Run()
	IsReady() bool
	Port() (int, error)
	Shutdown() error
}

type server struct {
	redirectTarget string
	port           int

	listener   net.Listener
	httpServer *http.Server
}

// Run initializes and starts the server.
// Once started, server listens for requests till app termination
func (server *server) Run() {

	// Create a listener for port
	var err error
	server.listener, err = net.Listen("tcp", fmt.Sprintf(":%v", server.port))
	if err != nil {
		panic(err)
	}

	// Create server
	httpServer := &http.Server{
		Handler: http.RedirectHandler(server.redirectTarget, http.StatusMovedPermanently),
	}

	// listen for requests till app termination
	server.httpServer = httpServer
	server.httpServer.Serve(server.listener)
}

func (server *server) IsReady() bool {
	return server.httpServer != nil
}

func (server *server) Port() (int, error) {

	if !server.IsReady() {
		return -1, fmt.Errorf("Server is not running")
	}

	return server.listener.Addr().(*net.TCPAddr).Port, nil
}

func (server *server) Shutdown() error {

	if !server.IsReady() {
		return fmt.Errorf("Server is not running")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := server.httpServer.Shutdown(ctx)
	server.httpServer = nil
	return err
}
