package repositories

import (
	"context"
	"errors"

	"github.com/crisantiz/calendar-golang/database"
	"github.com/crisantiz/calendar-golang/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	collection *mongo.Collection
}

func (r *Repository) GetOne(id string) (models.Calendar, error) {
	calendarId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return models.Calendar{}, errors.New("invalid calendar id")
	}

	filter := bson.D{primitive.E{Key: "_id", Value: calendarId}}

	var calendar models.Calendar

	r.collection.FindOne(context.Background(), filter).Decode(&calendar)

	return calendar, nil
}

func New() Repository {
	db := database.New("calendar-golang")

	return Repository{
		collection: db.GetCollection("calendars"),
	}
}
