package dblayer

import (
	"github.com/ibiscum/Cloud-Native-Programming-with-Golang/chapter02/lib/persistence"
	"github.com/ibiscum/Cloud-Native-Programming-with-Golang/chapter02/lib/persistence/mongolayer"
)

type DBTYPE string

const (
	MONGODB  DBTYPE = "mongodb"
	DYNAMODB DBTYPE = "dynamodb"
)

func NewPersistenceLayer(options DBTYPE, connection string) (persistence.DatabaseHandler, error) {

	switch options {
	case MONGODB:
		return mongolayer.NewMongoDBLayer(connection)
	}
	return nil, nil
}
