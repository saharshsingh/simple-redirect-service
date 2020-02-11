package context

import (
	"os"
	"simple-redirect-service/app"
	"simple-redirect-service/test"
	"syscall"
	"testing"
)

func TestBuild(t *testing.T) {

	// make sure context can be built
	os.Setenv("SRS_REDIRECT_TARGET", "http://example.org")
	os.Setenv("PORT", "0")
	ctx := Build()

	// make sure app can be started
	go ctx.App.Run()
	appStatus := <-ctx.Status
	test.AssertEquals("", app.InitializingStatus, appStatus.Status, t)
	appStatus = <-ctx.Status
	test.AssertEquals("", app.ReadyStatus, appStatus.Status, t)

	// make sure app can be terminated
	ctx.Signal <- syscall.SIGTERM
	appStatus = <-ctx.Status
	test.AssertEquals("", app.TerminatedStatus, appStatus.Status, t)
}
