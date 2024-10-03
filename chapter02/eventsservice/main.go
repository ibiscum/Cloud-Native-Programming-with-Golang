package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"

	"github.com/ibiscum/Cloud-Native-Programming-with-Golang/chapter02/eventsservice/rest"
	"github.com/ibiscum/Cloud-Native-Programming-with-Golang/chapter02/lib/configuration"
	"github.com/ibiscum/Cloud-Native-Programming-with-Golang/chapter02/lib/persistence/dblayer"
)

func main() {
	log.SetFlags(log.Lshortfile)

	defaultConfPath, err := filepath.Abs("../lib/configuration/config.json")
	if err != nil {
		log.Fatal(err)
	}

	confPath := flag.String("conf", defaultConfPath, "flag to set the path to the configuration json file")
	flag.Parse()

	//extract configuration
	config, err := configuration.ExtractConfiguration(*confPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connecting to database")
	dbhandler, err := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("RESTful API start")
	log.Fatal(rest.ServeAPI(config.RestfulEndpoint, dbhandler))
}
