package repository

import (
	"context"
	"github.com/hellodoge/courses-tg-bot/courses"
	"github.com/hellodoge/courses-tg-bot/pkg/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	searchMongoFieldSkip = "skip"
)

type searchMongoDB struct {
	Query   string           `bson:"query"`
	Skip    int64            `bson:"skip"`
	Results []courses.Course `bson:"results"`
}

func (r *CoursesMongoDB) SearchCourses(query string) ([]courses.Course, error) {
	collection := r.db.Collection(coursesCollectionMongoDB)
	var coursesMDB []CourseMongoDB
	cursor, err := collection.Find(context.Background(), bson.M{
		"$text": bson.M{
			"$search": query,
		},
	}, &options.FindOptions{
		Sort: bson.M{
			"score": bson.M{
				"$meta": "textScore",
			},
		},
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

func (r *CoursesMongoDB) GetMoreSearchResults(searchID string, limit int64) ([]courses.Course, error) {
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
	}
	from := util.MinInt64(search.Skip, int64(len(search.Results)-1))
	to := util.MinInt64(search.Skip+limit, int64(len(search.Results)))
	return search.Results[from:to], err
}

func (r *CoursesMongoDB) NewSearch(query string, results []courses.Course, skip int64) (string, error) {
	collection := r.db.Collection(searchesCollectionMongoDB)
	var search = searchMongoDB{
		Query:   query,
		Skip:    skip,
		Results: results,
	}
	id, err := collection.InsertOne(context.Background(), search)
	if err != nil {
		return "", err
	}
	return id.InsertedID.(primitive.ObjectID).Hex(), nil
}
