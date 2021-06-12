package database

import "go.mongodb.org/mongo-driver/mongo"

type Database struct {
	database string
}

func (c *Database) GetCollection(collection string) *mongo.Collection {
	client := GetMongoConnection()

	return client.Database(c.database).Collection(collection)
}

func New(database string) Database {
	return Database{
		database,
	}
}
