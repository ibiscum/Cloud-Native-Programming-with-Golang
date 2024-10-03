package mongolayer

import (
	"context"
	"log"

	"github.com/ibiscum/Cloud-Native-Programming-with-Golang/chapter02/lib/persistence"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	DB     = "myevents"
	USERS  = "users"
	EVENTS = "events"
)

type MongoDBLayer struct {
	client *mongo.Client
}

func NewMongoDBLayer(connection string) (persistence.DatabaseHandler, error) {
	c, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connection))
	if err != nil {
		log.Fatal(err)
	}

	if err := c.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	return &MongoDBLayer{
		client: c,
	}, err
}

func (mgoLayer *MongoDBLayer) AddEvent(e persistence.Event) ([]byte, error) {
	// 	s := mgoLayer.getFreshSession()
	// 	defer s.Close()

	// 	if e.ID.IsZero() {
	// 		e.ID = bson.NewObjectID()
	// 	}

	// 	if e.Location.ID.IsZero() {
	// 		e.Location.ID = bson.NewObjectID()
	// 	}

	// 	b, _ := e.ID.MarshalJSON()

	// return b, s.DB(DB).C(EVENTS).Insert(e)
	return nil, nil
}

// func (mgoLayer *MongoDBLayer) FindEvent(id []byte) (persistence.Event, error) {
// 	s := mgoLayer.getFreshSession()
// 	defer s.Close()
// 	e := persistence.Event{}

// 	err := s.DB(DB).C(EVENTS).FindId(bson.ObjectID(id)).One(&e)
// 	return e, err
// }

// func (mgoLayer *MongoDBLayer) FindEventByName(name string) (persistence.Event, error) {
// 	s := mgoLayer.getFreshSession()
// 	defer s.Close()
// 	e := persistence.Event{}
// 	err := s.DB(DB).C(EVENTS).Find(bson.M{"name": name}).One(&e)
// 	return e, err
// }

// func (mgoLayer *MongoDBLayer) FindAllAvailableEvents() ([]persistence.Event, error) {
// 	s := mgoLayer.getFreshSession()
// 	defer s.Close()
// 	events := []persistence.Event{}
// 	err := s.DB(DB).C(EVENTS).Find(nil).All(&events)
// 	return events, err
// }

// func (mgoLayer *MongoDBLayer) getFreshSession() *mgo.Session {
// 	return mgoLayer.session.Copy()
// }
