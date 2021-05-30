package repository

import (
	"context"
	"github.com/hellodoge/courses-tg-bot/courses"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	searchMongoFieldSkip = "skip"
)

type searchMongoDB struct {
	Query string `bson:"query"`
	Skip  int64  `bson:"skip"`
}

func (r *CoursesMongoDB) SearchCourses(query string, limit, skip int64) ([]courses.Course, error) {
	collection := r.db.Collection(coursesCollectionMongoDB)
	var coursesMDB []CourseMongoDB
	cursor, err := collection.Find(context.Background(), bson.M{
		"$text": bson.M{
			"$search": query,
		},
	}, &options.FindOptions{
		Limit: &limit,
		Skip:  &skip,
	})
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.Background(), &coursesMDB)
	if err != nil {
		return nil, err
	}
	var result []courses.Course
	for _, course := range coursesMDB {
		result = append(result, courses.Course{
			ID:    course.ID.Hex(),
			Title: course.Title,
		})
	}
	return result, nil
}

func (r *CoursesMongoDB) SearchCoursesBySearchID(searchID string, limit int64) ([]courses.Course, error) {
	collection := r.db.Collection(searchesCollectionMongoDB)
	var search searchMongoDB
	id, err := primitive.ObjectIDFromHex(searchID)
	if err != nil {
		return nil, ErrorInvalidSearchID
	}
	err = collection.FindOneAndUpdate(context.Background(),
		bson.M{"_id": id},
		bson.M{"$inc": bson.M{searchMongoFieldSkip: limit}},
	).Decode(&search)
	if err == mongo.ErrNoDocuments {
		return nil, ErrorInvalidSearchID
	} else if err != nil {
		return nil, err
	}
	result, err := r.SearchCourses(search.Query, limit, search.Skip)
	return result, err
}

func (r *CoursesMongoDB) NewSearch(query string, skip int64) (string, error) {
	collection := r.db.Collection(searchesCollectionMongoDB)
	var search = searchMongoDB{
		Query: query,
		Skip:  skip,
	}
	id, err := collection.InsertOne(context.Background(), search)
	if err != nil {
		return "", err
	}
	return id.InsertedID.(primitive.ObjectID).Hex(), nil
}
