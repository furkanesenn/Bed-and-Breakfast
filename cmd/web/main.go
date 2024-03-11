package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/furkanesenn/Bed-and-Breakfast/pkg/config"
	"github.com/furkanesenn/Bed-and-Breakfast/pkg/handlers"
	"github.com/furkanesenn/Bed-and-Breakfast/pkg/render"

	"github.com/alexedwards/scs/v2"
)

const PORT_NUMBER = ":8080"

var appConfig config.AppConfig
var session *scs.SessionManager

func main() {

	// Change this to true when in production
	appConfig.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = appConfig.InProduction

	appConfig.Session = session

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	appConfig.TemplateCache = templateCache
	appConfig.UseCache = false

	render.NewTemplates(&appConfig)

	repo := handlers.NewRepo(&appConfig)
	handlers.NewHandlers(repo)

	// http.HandleFunc("/", repo.Home)
	// http.HandleFunc("/about", repo.About)

	fmt.Println("Starting server on port", PORT_NUMBER)
	// _ = http.ListenAndServe(PORT_NUMBER, nil)

	srv := &http.Server{
		Addr:    PORT_NUMBER,
		Handler: routes(&appConfig),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
