package main

import (
	"io"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/vagrant-technology/squad-leader/handlers"
)

func ConfigureWS(app *fiber.App) {
	log.Println("Configuring websocket route handlers...")

	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	//app.Get("/ws/:room/:user", websocket.New(handlers.LobbyConnection))
	app.Get("/ws/join/:room/:user", websocket.New(handlers.LobbyConnection))
	//app.Get("/ws/game/:session", websocket.New(handlers.GameConnection))
}

func ConfigureApi(app *fiber.App) {
	log.Println("Configuring API route handlers...")

	api := app.Group("/api")

	// ----- Web API section -----
	api.Post("/CreateRoom", handlers.CreateRoom)
	//api.Post("/JoinRoom", handlers.JoinRoom)
}

func main() {
	// log.Println("----- ----- ----- ----- -----")
	// create the file if it does not exist, otherwise append to it
	file, err := os.OpenFile(
		"sl-server.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	logger := log.Default()
	logger.SetOutput(io.MultiWriter(file, os.Stderr))
	logger.SetPrefix("SL-Server: ")
	logger.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)

	log.Println("----- ----- ----- ----- -----")
	log.Println("SL Logging Session Intialized")
	log.Println("----- ----- ----- ----- -----")

	serverConfig := fiber.Config{
		ServerHeader: "SL Application Server",
	}
	log.Println("Starting golang SL server...")
	app := fiber.New(serverConfig)

	ConfigureWS(app)
	ConfigureApi(app)
	//apiGroup := app.Group("/api")
	//Router(app)

	app.Listen(":3001")
	log.Println("SL Server now active on port 3001")
}
