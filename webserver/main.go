package main

import (
	"embed"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/acme/autocert"
	"io/fs"
	"net/http"
	"os"
	"strings"
)

var (
	//go:embed web/build
	embeddedFiles embed.FS
)

func getFileSystem(name string) http.FileSystem {
	fsys, err := fs.Sub(embeddedFiles, "web/build"+name)
	if err != nil {
		panic(err)
	}
	return http.FS(fsys)
}

func main() {
	var (
		assetHandler = func(name string) http.Handler {
			return http.FileServer(getFileSystem(name))
		}
		httpServer  = echo.New()
		httpsServer = echo.New()
	)

	// HTTP server
	httpServer.Pre(middleware.HTTPSRedirect())
	go httpServer.Start(":80")

	hostList := strings.Split(os.Getenv("HOST_LIST"), ",")
	// HTTPS server
	httpsServer.AutoTLSManager.HostPolicy = autocert.HostWhitelist(hostList...)
	// Cache certificates
	httpsServer.AutoTLSManager.Cache = autocert.DirCache(".cache")

	httpsServer.Use(middleware.Recover())
	httpsServer.Use(middleware.Logger())
	httpsServer.GET("/*", echo.WrapHandler(assetHandler("")))
	httpsServer.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", assetHandler("/static"))))

	httpsServer.Logger.Fatal(httpsServer.StartAutoTLS(":443"))
}
