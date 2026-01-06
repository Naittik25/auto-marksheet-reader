package services

import (
	"auto-marksheet-reader/internal/storage"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type MarksheetService struct{}

// Create saves a new marksheet to the DB
func (s *MarksheetService) Create(m *storage.Marksheet) error {
	collection := storage.DB.Collection("marksheets")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	m.UploadDate = time.Now()
	_, err := collection.InsertOne(ctx, m)
	return err
}

// GetAll fetches all marksheets
func (s *MarksheetService) GetAll() ([]storage.Marksheet, error) {
	collection := storage.DB.Collection("marksheets")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []storage.Marksheet
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}
