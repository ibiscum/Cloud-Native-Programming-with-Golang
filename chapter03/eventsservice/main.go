package main

import (
	"flag"
	"log"

	"github.com/ibiscum/Cloud-Native-Programming-with-Golang/chapter03/eventsservice/rest"
	"github.com/ibiscum/Cloud-Native-Programming-with-Golang/chapter03/lib/configuration"
	"github.com/ibiscum/Cloud-Native-Programming-with-Golang/chapter03/lib/persistence/dblayer"
)

func main() {

	confPath := flag.String("conf", `.\configuration\config.json`, "flag to set the path to the configuration json file")
	flag.Parse()
	//extract configuration
	config, _ := configuration.ExtractConfiguration(*confPath)

	log.Println("Connecting to database")
	dbhandler, err := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connection successful... ")
	//RESTful API start
	httpErrChan, httptlsErrChan := rest.ServeAPI(config.RestfulEndpoint, config.RestfulTLSEndPint, dbhandler)
	select {
	case err := <-httpErrChan:
		log.Fatal("HTTP Error: ", err)
	case err := <-httptlsErrChan:
		log.Fatal("HTTPS Error: ", err)
	}
}
