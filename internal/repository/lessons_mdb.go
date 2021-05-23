package repository

import (
	"context"
	"github.com/hellodoge/courses-tg-bot/courses"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *CoursesMongoDB) writeLessons(ctx context.Context, lessons []courses.Lesson) ([]LessonInfoMongoDB, error) {
	collection := r.db.Collection(lessonsCollectionMongoDB)
	var (
		nextLessonID primitive.ObjectID
		result       = make([]LessonInfoMongoDB, len(lessons))
	)
	for i := len(lessons) - 1; i >= 0; i-- {
		lesson := lessons[i]
		var lessonMongo = LessonMongoDB{
			NextLessonID: nextLessonID,
			Title:        lesson.Title,
			Description:  lesson.Description,
			Videos:       lesson.Videos,
			Photos:       lesson.Photos,
			Documents:    lesson.Documents,
		}
		id, err := collection.InsertOne(ctx, lessonMongo)
		if err != nil {
			return nil, err
		}
		nextLessonID = id.InsertedID.(primitive.ObjectID)
		result[i].ID = id.InsertedID.(primitive.ObjectID)
		result[i].Title = lesson.Title
	}
	return result, nil
}
