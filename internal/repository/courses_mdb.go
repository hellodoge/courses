package repository

import (
	"context"
	"github.com/hellodoge/courses-tg-bot/courses"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CoursesMongoDB struct {
	client *mongo.Client
	db     *mongo.Database
}

func NewCoursesMongoDB(client *mongo.Client) *CoursesMongoDB {
	return &CoursesMongoDB{
		client: client,
		db:     client.Database(databaseMongoDB),
	}
}

func (r *CoursesMongoDB) NewCourse(course *courses.Course) (string, error) {
	var id string
	err := r.client.UseSession(context.Background(), func(ctx mongo.SessionContext) error {
		err := ctx.StartTransaction()
		var lessons []LessonInfoMongoDB = nil
		if course.Lessons != nil {
			var err error
			lessons, err = r.writeLessons(ctx, course.Lessons)
			if err != nil {
				if err := ctx.AbortTransaction(ctx); err != nil {
					logrus.Error(err)
				}
				return err
			}
		}
		var courseMongo = CourseMongoDB{
			Title:       course.Title,
			Description: course.Description,
			Photo:       course.Preview,
			Lessons:     lessons,
		}
		coursesCollection := r.db.Collection(coursesCollectionMongoDB)
		result, err := coursesCollection.InsertOne(ctx, courseMongo)
		if err != nil {
			if err := ctx.AbortTransaction(ctx); err != nil {
				logrus.Error(err)
			}
			return err
		}
		id = result.InsertedID.(primitive.ObjectID).Hex()
		return ctx.CommitTransaction(ctx)
	})
	return id, err
}

func (r *CoursesMongoDB) GetCourse(idHex string) (*courses.Course, error) {
	id, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		return nil, nil
	}
	collection := r.db.Collection(coursesCollectionMongoDB)
	var courseMDB CourseMongoDB
	err = collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&courseMDB)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	var lessons = make([]courses.Lesson, 0, len(courseMDB.Lessons))
	for _, lessonMDB := range courseMDB.Lessons {
		lessons = append(lessons, courses.Lesson{ID: lessonMDB.ID.Hex()})
	}
	var course = &courses.Course{
		ID:          idHex,
		Title:       courseMDB.Title,
		Description: courseMDB.Description,
		Preview:     courseMDB.Photo,
		Lessons:     lessons,
	}
	return course, nil
}
