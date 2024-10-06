package persistence

import (
	"context"
	"errors"
	log "rest/logging"
	"rest/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://localhost:27017/?maxPoolSize=20&w=majority"

var Connection *models.DbInstance

func Connect() (*models.DbInstance, error) {

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		log.Error(err)
		return nil, err
	}
	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Error(err)
		return nil, err
	}
	log.Info("Pinged your deployment. You successfully connected to MongoDB!")
	Connection = &models.DbInstance{Client: client}
	return Connection, nil
}
func Check() (*models.DbInstance, error) {

	if Connection == nil {
		return nil, errors.New("connection not available")
	}
	if err := Connection.Client.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}

	return Connection, nil
}
