package dao

import (
	"context"
	mgo "coolcar/shared/mongo"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const openIDField = "open_id"

type Mongo struct {
	col      *mongo.Collection
	newObjID func() primitive.ObjectID
}

// NewMongo create a new mongo dao.
func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		//用户db
		col:      db.Collection("account"),
		newObjID: primitive.NewObjectID,
	}
}

func (m *Mongo) ResolveAccountID(ctx context.Context, openID string) (string, error) {
	insertedID := m.newObjID()

	res := m.col.FindOneAndUpdate(ctx, bson.M{
		openIDField: openID,
	}, mgo.SetOnInsert(bson.M{
		mgo.IDField: insertedID,
		openIDField: openID,
	}), options.FindOneAndUpdate().
		SetUpsert(true).
		SetReturnDocument(options.After))

	if err := res.Err(); err != nil {
		return "", fmt.Errorf("cannot findOneAndUpdate:%v", err)
	}

	var row mgo.ObjID

	err := res.Decode(&row)
	if err != nil {
		return "", fmt.Errorf("cannot decode result:%v", err)
	}

	return row.ID.Hex(), nil
}