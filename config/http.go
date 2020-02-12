package config

// HTTPConfig contains HTTP related configuration
type HTTPConfig struct {
	RedirectTarget          string
	StartupTimeoutInSeconds int
	Port                    int
}

// InitFromEnvironment populates configuration from environment variables
func (config *HTTPConfig) InitFromEnvironment() {
	config.RedirectTarget = MustLoadEnvValueAsString("SRS_REDIRECT_TARGET")
	config.StartupTimeoutInSeconds = LoadEnvValueAsInteger("SRS_SERVER_STARTUP_TIMEOUT_SECONDS", 5)
	config.Port = LoadEnvValueAsInteger("PORT", 8080)
}
