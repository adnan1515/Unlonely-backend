package persistence

import (
	"context"
	log "rest/logging"
	"rest/models"

	"go.mongodb.org/mongo-driver/bson"
)

func SaveNewUser(data models.User) (bool, error) {
	connection, err := Check()
	if err != nil {
		log.Error("not Able to connect to database")
		return false, err
	}
	coll := connection.Client.Database("Unlonely").Collection("users")

	cursor, err := coll.Find(context.TODO(), bson.D{{Key: "email", Value: data.Email}})
	if err != nil {
		log.Error(err)
		return false, err
	}
	if cursor.Next(context.Background()) {
		log.Info("Record ", data.Email, "already Exist")
		return false, nil
	}
	cursor.Close(context.Background())
	res, err := coll.InsertOne(context.TODO(), data)
	if err != nil {
		log.Error("Internal issue occured")
		return false, err
	}
	log.Info("Inserted successfully with ", res.InsertedID, " ID")
	return true, nil
}

func LoginUser(email string) (*models.User, error) {
	connection, err := Check()
	if err != nil {
		log.Error("not Able to connect to database")
		return nil, err
	}
	coll := connection.Client.Database("Unlonely").Collection("users")
	cursor, err := coll.Find(context.Background(), bson.D{{Key: "email", Value: email}})
	if err != nil {
		return nil, err
	}
	var result models.User
	if cursor.Next(context.Background()) {
		err = cursor.Decode(&result)
		if err != nil {
			log.Error("Error caused while decoding result")
			return nil, err
		}
	} else {
		return &models.User{Id: 0}, nil
	}

	return &result, nil
}
