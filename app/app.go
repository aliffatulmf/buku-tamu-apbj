package app

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

type Certificates struct {
	Path string
	Name []string
}

type App struct {
	Server *gin.Engine
	Port   string
	Cert   *Certificates
}

func New(server *gin.Engine) *App {
	return &App{
		Server: server,
		Cert: &Certificates{
			Path: ".",
			Name: []string{"certificate.crt", "certificate-key.crt"},
		},
	}
}

func (app *App) noCert() bool {
	for _, c := range app.Cert.Name {
		lc := filepath.Join(app.Cert.Path, c)
		if _, err := os.Stat(lc); errors.Is(err, os.ErrNotExist) {
			return true
		}
	}

	return false
}

func (app *App) GetCertificateName() (cert string, key string) {
	for i := 0; i < 2; i++ {
		if strings.Contains(app.Cert.Name[i], "key") {
			key = app.Cert.Name[i]
		} else {
			cert = app.Cert.Name[i]
		}
	}

	if key == "" || cert == "" {
		fmt.Println("certificate name not valid")
		os.Exit(1)
	}

	return
}

func (app *App) SetPort(port string) {
	if strings.Contains(port, ":") {
		app.Port = port
	} else {
		app.Port = fmt.Sprintf(":%s", port)
	}
}

func (app *App) SetHandler(f func(a *App)) {
	f(app)
}

func (app *App) RunTLS() {
	if app.noCert() {
		fmt.Println("ERROR: no certificate found.")
		os.Exit(1)
	}

	cert, key := app.GetCertificateName()
	if err := app.Server.RunTLS(app.Port, cert, key); err != nil {
		fmt.Println("FATAL:", err.Error())
		os.Exit(1)
	}
}

func (app *App) RunHTTP() {
	if app.noCert() {
		fmt.Println("ERROR: no certificate found.")
		os.Exit(1)
	}

	if err := app.Server.Run(app.Port); err != nil {
		fmt.Println("FATAL:", err.Error())
		os.Exit(1)
	}
}
