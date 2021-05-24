package repository

import (
	"github.com/hellodoge/courses-tg-bot/courses"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	databaseMongoDB          = "courses"
	coursesCollectionMongoDB = "courses"
	lessonsCollectionMongoDB = "lessons"
)

type CourseMongoDB struct {
	ID          primitive.ObjectID     `bson:"_id,omitempty"`
	Title       string                 `bson:"title"`
	Description string                 `bson:"description,omitempty"`
	Photo       *courses.URLCollection `bson:"photo,omitempty"`
	Lessons     []LessonInfoMongoDB    `bson:"lessons"`
}

type LessonInfoMongoDB struct {
	ID    primitive.ObjectID `bson:"_id"`
	Title string             `bson:"title,omitempty"`
}

type LessonMongoDB struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	NextLessonID primitive.ObjectID `bson:"next_id,omitempty"`
	Title        string             `bson:"title,omitempty"`
	Description  string             `bson:"description,omitempty"`
	Documents    []courses.Document `bson:"documents,omitempty"`
}
