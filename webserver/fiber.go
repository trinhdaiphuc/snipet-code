package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Embed a single file
//go:embed web/build
var embedDir embed.FS

// Embed a directory
//go:embed web/build/static/*
var embedDirStatic embed.FS

func main() {
	app := fiber.New()
	// Default middleware config
	app.Use(logger.New())
	app.Use(cors.New())

	app.Use("/", filesystem.New(filesystem.Config{
		Root: http.FS(embedDir),
	}))

	// Access file "image.png" under `static/` directory via URL: `http://<server>/static/image.png`.
	// Without `PathPrefix`, you have to access it via URL:
	// `http://<server>/static/static/image.png`.
	app.Use("/static", filesystem.New(filesystem.Config{
		Root:       http.FS(embedDirStatic),
		PathPrefix: "",
		Browse:     true,
	}))

	log.Fatal(app.Listen(":8080"))
}
