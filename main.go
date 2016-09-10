//go:generate swagger generate spec -o swagger.json
package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/patito/bidu/configuration"
	"github.com/patito/bidu/handler"
	"github.com/patito/bidu/model"

	"database/sql"

	_ "github.com/lib/pq"
)

func main() {
	config := flag.String("config", "configuration/config.yaml", "Configuration file")
	flag.Parse()

	conf, err := configuration.New(*config)
	if nil != err {
		log.Fatal("Failed to load configuration", err)
	}

	db, err := sql.Open("postgres", conf.StringConnection())
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxIdleConns(300)
	defer db.Close()

	m := model.New(db)
	h := handler.New(m)

	router := httprouter.New()

	router.POST("/api/sites", h.CreateSite)

	router.GET("/api/sites", h.GetAllSites)
	router.GET("/api/sites/:name", h.GetSite)

	router.DELETE("/api/sites/:name", h.DeleteSite)
	router.DELETE("/api/sites", h.DeleteAllSites)

	router.PUT("/api/sites/:name", h.UpdateSite)

	log.Fatal(http.ListenAndServe(":3000", router))
}
