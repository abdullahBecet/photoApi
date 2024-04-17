package repository

import (
	"context"
	"log"
	"photoApi/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PhotoRepositoryDB struct {
	PhotoCollection *mongo.Collection
}
type PhotoRepsitory interface {
	GetAll() ([]models.Photo, error)
}

func (t PhotoRepositoryDB) GetAll() ([](models.Photo), error) {
	var photo models.Photo
	var photos []models.Photo

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := t.PhotoCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	for result.Next(ctx) {
		if err := result.Decode(&photo); err != nil {
			log.Fatalln(err)
		}
		photos = append(photos, photo)
	}
	return photos, nil
}

func NewPhotoRepositoryDb(dbClient *mongo.Collection) PhotoRepositoryDB {
	return PhotoRepositoryDB{PhotoCollection: dbClient}
}
