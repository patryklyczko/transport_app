package db

import (
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

const (
	confOptMongoPassword = "MONGO_PASSWORD"
	confOptMongoUser     = "MONGO_USER"
	confOptMongoDatabase = "MONGO_DATABASE"
)

type DBController struct {
	log    logrus.FieldLogger
	db     *mongo.Database
	bucket *gridfs.Bucket
}

func NewDBController(log logrus.FieldLogger, db *mongo.Database, bucket *gridfs.Bucket) *DBController {
	return &DBController{
		log:    log,
		db:     db,
		bucket: bucket,
	}
}
