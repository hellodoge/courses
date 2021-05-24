package repository

import (
	"context"
	"github.com/hellodoge/courses-tg-bot/courses"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func (r *CoursesMongoDB) GetLesson(idHex string) (*courses.Lesson, error) {
	id, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		return nil, nil
	}
	collection := r.db.Collection(lessonsCollectionMongoDB)
	var lessonMDB LessonMongoDB
	err = collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&lessonMDB)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	lesson := &courses.Lesson{
		ID:          idHex,
		Title:       lessonMDB.Title,
		Description: lessonMDB.Description,
		Photos:      lessonMDB.Photos,
		Videos:      lessonMDB.Videos,
		Documents:   lessonMDB.Documents,
	}
	if !lessonMDB.NextLessonID.IsZero() {
		lesson.NextLessonID = lessonMDB.NextLessonID.Hex()
	}
	return lesson, nil
}
