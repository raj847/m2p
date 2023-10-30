package repository

import (
	"context"
	"m2p/models"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"

	"gorm.io/gorm"
)

type TrxRepository struct {
	db      *gorm.DB
	mongoDB *mongo.Database
}

func NewTrxRepository(db *gorm.DB, mongoDB *mongo.Database) *TrxRepository {
	return &TrxRepository{
		db:      db,
		mongoDB: mongoDB,
	}
}

func (t *TrxRepository) CreateTrx(ctx context.Context, trx *models.Trx) error {
	err := t.db.WithContext(ctx).Create(&trx).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *TrxRepository) GetData(start string, ctx context.Context, end string) (trxlist []*models.TrxRequest, exists bool, err error) {

	filter := bson.M{
		"docDate": bson.M{"$gte": start, "$lte": end},
	}

	data, err := t.mongoDB.Collection("trx").Find(ctx, filter)
	if err != nil {
		return nil, false, err
	}

	for data.Next(ctx) {
		var val models.TrxRequest
		if err = data.Decode(&val); err != nil {
			return nil, false, err
		}

		trxlist = append(trxlist, &val)
	}
	data.Close(ctx)

	if len(trxlist) == 0 {
		return nil, false, nil
	}

	return trxlist, true, nil
}
