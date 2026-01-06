package storage

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Marksheet represents the data we save to MongoDB
type Marksheet struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Filename    string             `bson:"filename" json:"filename"`
	FilePath    string             `bson:"file_path" json:"file_path"`
	UploadDate  time.Time          `bson:"upload_date" json:"upload_date"`
	StudentName string             `bson:"student_name" json:"student_name"`
	RollNumber  string             `bson:"roll_number" json:"roll_number"`
	RawText     string             `bson:"raw_text" json:"raw_text"` // The text from OCR
	Marks       []SubjectMark      `bson:"marks" json:"marks"`       // The extracted table
}

type SubjectMark struct {
	Subject string `bson:"subject" json:"subject"`
	Score   string `bson:"score" json:"score"`
	Grade   string `bson:"grade" json:"grade"`
}
