package dblayer

import (
	"github.com/ibiscum/Cloud-Native-Programming-with-Golang/chapter03/lib/persistence"
	"github.com/ibiscum/Cloud-Native-Programming-with-Golang/chapter03/lib/persistence/mongolayer"
)

type DBTYPE string

const (
	MONGODB    DBTYPE = "mongodb"
	DOCUMENTDB DBTYPE = "documentdb"
	DYNAMODB   DBTYPE = "dynamodb"
)

func NewPersistenceLayer(options DBTYPE, connection string) (persistence.DatabaseHandler, error) {

	switch options {
	case MONGODB:
		return mongolayer.NewMongoDBLayer(connection)
	}
	return nil, nil
}
