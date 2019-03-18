// Package main is the entry point for the web application.
package main

import (
	"log"
	"runtime"
	"github.com/huzhaer/qianxun/lib/boot"
	"github.com/huzhaer/qianxun/lib/env"
	"github.com/huzhaer/teamlite_core/router"
	"github.com/huzhaer/qianxun/core/server"
	"fmt"
)

// init sets runtime settings.
func init() {
	// Verbose logging with file name and line number
	log.SetFlags(log.Lshortfile)
	// Use all CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// main loads the configuration file, registers the services, applies the
// middleware to the router, and then starts the HTTP and HTTPS listeners.
func main() {

	// Load the configuration file
	config, err := env.LoadConfig("env.json")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print("START\n")
	// Register the services
	// boot.RegisterServices(config)

    boot.RegisterServices(config)
	// CheckOrStartFrontendServer()
	// Retrieve the middleware
	handler := boot.SetUpMiddleware(router.Instance())

	// Start the HTTP and HTTPS listeners
	server.Run(
		handler,       // HTTP handler
		handler,       // HTTPS handler
		config.Server, // Server settings
	)
}
