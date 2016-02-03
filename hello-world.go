// hello-world is a simple web server for testing connections and configurations.
package main

import "github.com/composer22/hello-world/server"

// main is the main entry point for the application or server launch.
func main() {
	server.New().Start()
}
