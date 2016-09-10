package main

import (
	"flag"
	"log"

	"github.com/mattes/migrate/migrate"
	"github.com/patito/bidu/configuration"

	_ "github.com/mattes/migrate/driver/postgres"
)

func main() {
	config := flag.String("config", "configuration/config.yaml", "Configuration file")
	flag.Parse()

	conf, err := configuration.New(*config)
	if nil != err {
		log.Fatal("Failed to load configuration", err)
	}

	log.Printf("running migrations on %s", conf.StringConnection())

	if errs, ok := migrate.UpSync(conf.StringConnection(), "./migrations"); !ok {
		log.Fatal(errs)
	}

	log.Print("migrations completed successfully")
}
