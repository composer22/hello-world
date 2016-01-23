// coreos-artifactory-monitor is a simple server that monitors artifactory for image version changes and
// deploys those images to a cluster.
package main

import (
	"github.com/composer22/hello-world/logger"
	"github.com/composer22/hello-world/server"
)

var (
	log *logger.Logger
)

func init() {
	log = logger.New(logger.UseDefault, false)
}

// main is the main entry point for the application or server launch.
func main() {
	s := server.New()
	s.Start()
}
