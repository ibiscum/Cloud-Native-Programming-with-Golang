package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ibiscum/Cloud-Native-Programming-with-Golang/chapter02/eventsservice/rest"
	"github.com/ibiscum/Cloud-Native-Programming-with-Golang/chapter02/lib/configuration"
	"github.com/ibiscum/Cloud-Native-Programming-with-Golang/chapter02/lib/persistence/dblayer"
)

func main() {

	confPath := flag.String("conf", `.\configuration\config.json`, "flag to set the path to the configuration json file")
	flag.Parse()
	//extract configuration
	config, _ := configuration.ExtractConfiguration(*confPath)

	fmt.Println("Connecting to database")
	dbhandler, _ := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)
	//RESTful API start
	log.Fatal(rest.ServeAPI(config.RestfulEndpoint, dbhandler))
}
