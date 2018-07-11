package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	nlog "github.com/nuveo/log"
	"github.com/prest/adapters/postgres"
	"github.com/prest/config"
	"github.com/prest/config/router"
	"github.com/prest/controllers"
	"github.com/prest/middlewares"
)

func main() {
	config.Load()

	// Load Postgres Adapter
	postgres.Load()

	startServer()
}

func MakeHandler() http.Handler {
	n := middlewares.GetApp()
	r := router.Get()
	r.HandleFunc("/{queriesLocation}/{script}", controllers.ExecuteFromScripts)
	n.UseHandler(r)
	return n
}

func startServer() {
	http.Handle(config.PrestConf.ContextPath, MakeHandler())
	l := log.New(os.Stdout, "[prest] ", 0)

	if !config.PrestConf.AccessConf.Restrict {
		nlog.Warningln("You are running pREST in public mode.")
	}

	if config.PrestConf.Debug {
		nlog.DebugMode = config.PrestConf.Debug
		nlog.Warningln("You are running pREST in debug mode.")
	}
	addr := fmt.Sprintf("%s:%d", config.PrestConf.HTTPHost, config.PrestConf.HTTPPort)
	l.Printf("listening on %s and serving on %s", addr, config.PrestConf.ContextPath)
	if config.PrestConf.HTTPSMode {
		l.Fatal(http.ListenAndServeTLS(addr, config.PrestConf.HTTPSCert, config.PrestConf.HTTPSKey, nil))
	}
	l.Fatal(http.ListenAndServe(addr, nil))
}
