package http

import (
	"fmt"
	"net/http"
	"simple-redirect-service/test"
	"simple-redirect-service/utils"
	"testing"
)

func TestPreInitState(t *testing.T) {

	// arrange
	server := Bootstrap(&ContextIn{RedirectTarget: "http://localhost:3000", Port: 0}).Server

	// act
	isReady := server.IsReady()
	_, portError := server.Port()
	shutdownError := server.Shutdown()

	// assert
	test.AssertFalse("Expected IsReady to return false", isReady, t)
	test.AssertEquals("", "Server is not running", portError.Error(), t)
	test.AssertEquals("", "Server is not running", shutdownError.Error(), t)

}

func TestRun_with_bad_port(t *testing.T) {

	// arrange
	server := Bootstrap(&ContextIn{RedirectTarget: "http://localhost:3000", Port: -1}).Server

	// assert via defer
	defer test.AssertPanic("Expected panic during Run", t)

	// act
	server.Run()

}

func TestRun_with_some_routes(t *testing.T) {

	// arrange
	server := Bootstrap(&ContextIn{RedirectTarget: "http://example.org", Port: 0}).Server

	// act
	go server.Run()
	startupError := utils.WaitTill(func() bool { return server.IsReady() }, 10)

	// assert

	// ---
	// Assert healthy startup
	// ---

	test.AssertTrue("Expected no errors starting up", startupError == nil, t)
	test.AssertTrue("Expected IsReady to return false", server.IsReady(), t)

	port, portErr := server.Port()
	test.AssertTrue("Expected no errors getting port", portErr == nil, t)
	test.AssertTrue("Expected server to pick a random port", port > 0, t)

	// ---
	// Test routes
	// ---

	url := fmt.Sprintf("http://localhost:%d", port)

	// Get
	resp, respErr := http.Get(url)
	test.AssertTrue("Expected no errors doing GET /", respErr == nil, t)
	test.AssertEquals("", "http://example.org", fmt.Sprintf("%v", resp.Request.URL), t)
	test.AssertEquals("", http.StatusOK, resp.StatusCode, t)

	// ---
	// Test shutdown
	// ---

	test.AssertTrue("Expected no errors shutting down", server.Shutdown() == nil, t)

}
