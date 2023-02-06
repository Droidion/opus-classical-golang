package main

import (
	"github.com/spf13/viper"
)

// config contains application configuration
type config struct {
	ServerPort               int
	SentryDsn                string
	DatabaseConnectionString string
	CoversUri                string
	Csp                      string
	UmamiUri                 string
	UmamiWebsiteId           string
	IsDev                    bool
}

// initConfig reads, parses and returns application configuration.
func (app *application) initConfig() {
	viper.SetDefault("IS_DEV", false)
	viper.SetConfigFile("app.env")
	viper.AutomaticEnv()
	_ = viper.ReadInConfig()
	config := config{
		ServerPort:               viper.GetInt("SERVER_PORT"),
		SentryDsn:                viper.GetString("SENTRY_DSN"),
		DatabaseConnectionString: viper.GetString("DATABASE_CONNECTION_STRING"),
		CoversUri:                viper.GetString("COVERS_URI"),
		Csp:                      viper.GetString("CSP"),
		UmamiUri:                 viper.GetString("UMAMI_URI"),
		UmamiWebsiteId:           viper.GetString("UMAMI_WEBSITE_ID"),
		IsDev:                    viper.GetBool("IS_DEV"),
	}
	app.config = &config
}
