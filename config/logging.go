package config

import (
	"github.com/goravel/framework/support/facades"
)

func init() {
	config := facades.Config
	config.Add("logging", map[string]interface{}{
		//Default Log Channel
		//This option defines the default log channel that gets used when writing
		//messages to the logs. The name specified in this option should match
		//one of the channels defined in the "channels" configuration array.
		"default": config.Env("LOG_CHANNEL", "stack"),

		//Log Channels
		//Here you may configure the log channels for your application. Out of
		//the box, Laravel uses the Monolog PHP logging library. This gives
		//you a variety of powerful log handlers / formatters to utilize.
		//Available Drivers: "single", "daily", "custom", "stack"
		//Available Level: "debug", "info", "warn", "error"
		"channels": map[string]interface{}{
			"stack": map[string]interface{}{
				"driver":   "stack",
				"channels": []string{"daily"},
			},
			"single": map[string]interface{}{
				"driver": "single",
				"path":   "storage/logs/goravel.log",
				"level":  config.Env("LOG_LEVEL", "debug"),
			},
			"daily": map[string]interface{}{
				"driver": "daily",
				"path":   "storage/logs/goravel.log",
				"level":  config.Env("LOG_LEVEL", "debug"),
				"days":   7,
			},
		},
	})
}
