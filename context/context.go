package context

import (
	"fmt"
	"simple-redirect-service/app"
	"simple-redirect-service/config"
	"simple-redirect-service/http"
)

// Build bootstraps all app submodules and creates overall app context
func Build() *app.ContextOut {

	// app config
	appConfig := config.AppConfig{}
	appConfig.InitFromEnvironment()

	// http
	httpOut := http.Bootstrap(&http.ContextIn{
		RedirectTarget: appConfig.HTTPConfig.RedirectTarget,
		Port:           appConfig.HTTPConfig.Port,
	})

	// app
	return app.Bootstrap(&app.ContextIn{

		// server
		StartupTimeoutInSeconds: appConfig.HTTPConfig.StartupTimeoutInSeconds,
		HTTPServer:              httpOut.Server,

		// Add all shutdown hooks here
		ShutdownHooks: []app.ShutdownHook{
			func() {
				// make sure this is at beginning (runs last)
				fmt.Println("Graceful shutdown of application complete.")
			},
		},
	})
}
