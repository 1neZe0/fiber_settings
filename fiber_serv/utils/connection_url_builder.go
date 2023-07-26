package utils

import (
	"fmt"
	"github.com/valyala/fasttemplate"
	"os"
)

// ConnectionURLBuilder func for building URL connection.
func ConnectionURLBuilder(n string) (string, error) {
	// Define URL to connection.
	var url string

	// Switch given names.
	switch n {
	case "postgres":
		template := "host={{host}} port={{port}} user={{user}} password={{password}} dbname={{dbname}} sslmode{{sslmode}}"
		t := fasttemplate.New(template, "{{", "}}")
		url = t.ExecuteString(map[string]interface{}{
			"host":     os.Getenv("DB_HOST"),
			"post":     os.Getenv("DB_PORT"),
			"user":     os.Getenv("DB_USER"),
			"password": os.Getenv("DB_PASSWORD"),
			"dbname":   os.Getenv("DB_NAME"),
			"sslmode":  os.Getenv("DB_SSL_MODE"),
		})
	case "redis":
		// URL for Redis connection.
		template := "{{host}}:{{port}}"
		t := fasttemplate.New(template, "{{", "}}")
		url = t.ExecuteString(map[string]interface{}{
			"host": os.Getenv("REDIS_HOST"),
			"port": os.Getenv("REDIS_PORT"),
		})
	case "fiber_serv":
		template := "{{host}}:{{port}}"
		t := fasttemplate.New(template, "{{", "}}")
		url = t.ExecuteString(map[string]interface{}{
			//"host": os.Getenv("SERVER_HOST"),
			"host": "0.0.0.0",
			"port": "1488",
			//"port": os.Getenv("SERVER_PORT"),
		})
	default:
		// Return error message.
		return "", fmt.Errorf("connection name '%v' is not supported", n)
	}

	// Return connection URL.
	return url, nil
}
