package models

import "go.mongodb.org/mongo-driver/mongo"

type DbInstance struct {
	Client *mongo.Client
}
