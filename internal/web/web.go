package web

import (
	"embed"
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"go.quinn.io/ccf/assets"
	"go.quinn.io/dataq/boot"
	"go.quinn.io/dataq/internal/middleware"
	"go.quinn.io/dataq/internal/router"
)

//go:embed public
var assetsFS embed.FS

func Run(b *boot.Boot) {
	// Load content before starting server
	// content.Initialize()

	e := echo.New()
	e.Use(echomiddleware.Logger())

	// Register routes from generated code
	router.RegisterRoutes(e)

	// Attach public assets
	assets.Attach(
		e,
		"public",
		"internal/web/public",
		assetsFS,
		os.Getenv("USE_EMBEDDED_ASSETS") == "true",
	)

	e.Debug = true
	e.Use(middleware.BootContext(b))
	e.Use(middleware.EchoContext(e))
	e.HTTPErrorHandler = middleware.HTTPErrorHandler
	addRoutes(e)

	fmt.Println("Server starting on http://localhost:3000")
	if err := e.Start(":3000"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
