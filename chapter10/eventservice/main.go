package main

import (
	"flag"
	"fmt"
	"log"

	"net/http"

	"github.com/IBM/sarama"

	"github.com/ibiscum/Cloud-Native-Programming-with-Golang/chapter10/eventservice/rest"
	"github.com/ibiscum/Cloud-Native-Programming-with-Golang/chapter10/lib/configuration"
	"github.com/ibiscum/Cloud-Native-Programming-with-Golang/chapter10/lib/msgqueue"
	msgqueue_amqp "github.com/ibiscum/Cloud-Native-Programming-with-Golang/chapter10/lib/msgqueue/amqp"
	"github.com/ibiscum/Cloud-Native-Programming-with-Golang/chapter10/lib/msgqueue/kafka"
	"github.com/ibiscum/Cloud-Native-Programming-with-Golang/chapter10/lib/persistence/dblayer"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/streadway/amqp"
)

func main() {
	var eventEmitter msgqueue.EventEmitter

	confPath := flag.String("conf", `.\configuration\config.json`, "flag to set the path to the configuration json file")
	flag.Parse()
	//extract configuration
	config, _ := configuration.ExtractConfiguration(*confPath)

	switch config.MessageBrokerType {
	case "amqp":
		conn, err := amqp.Dial(config.AMQPMessageBroker)
		if err != nil {
			panic(err)
		}

		eventEmitter, err = msgqueue_amqp.NewAMQPEventEmitter(conn, "events")
		if err != nil {
			panic(err)
		}
	case "kafka":
		conf := sarama.NewConfig()
		conf.Producer.Return.Successes = true
		conn, err := sarama.NewClient(config.KafkaMessageBrokers, conf)
		if err != nil {
			panic(err)
		}

		eventEmitter, err = kafka.NewKafkaEventEmitter(conn)
		if err != nil {
			panic(err)
		}
	default:
		panic("Bad message broker type: " + config.MessageBrokerType)
	}

	fmt.Println("Connecting to database")
	dbhandler, _ := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)

	go func() {
		fmt.Println("Serving metrics API")
		h := http.NewServeMux()
		h.Handle("/metrics", promhttp.Handler())

		err := http.ListenAndServe(":9100", h)
		if err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Println("Serving API")
	//RESTful API start
	err := rest.ServeAPI(config.RestfulEndpoint, dbhandler, eventEmitter)
	if err != nil {
		panic(err)
	}
}
