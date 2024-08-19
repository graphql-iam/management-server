package database

import (
	"context"
	"github.com/graphql-iam/management-server/src/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/fx"
)

func NewDatabase(lc fx.Lifecycle, cfg config.Config) *mongo.Database {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.MongoUrl))
	if err != nil {
		panic(err)
	}
	db := client.Database("graphql-iam")

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return prepareDb(db, ctx)
		},
		OnStop: func(ctx context.Context) error {
			return client.Disconnect(ctx)
		},
	})

	return db
}

func prepareDb(db *mongo.Database, ctx context.Context) error {
	// TODO create Collections with jsonSchema validation

	pipeline := mongo.Pipeline{
		{
			{"$lookup", bson.D{
				{"from", "policies"},
				{"localField", "policyIds"},
				{"foreignField", "id"},
				{"as", "policies"},
			}},
		},
		{
			{"$project", bson.D{
				{"name", 1},
				{"policies", 1},
				{"_id", 0},
			}},
		},
	}

	err := db.CreateView(context.TODO(), "rolesWithPolicies", "roles", pipeline)
	if err != nil {
		return err
	}

	rolesNameIndexModel := mongo.IndexModel{
		Keys:    bson.D{{"name", 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err = db.Collection("roles").Indexes().CreateOne(ctx, rolesNameIndexModel)
	if err != nil {
		return err
	}

	policiesIdIndexModel := mongo.IndexModel{
		Keys:    bson.D{{"id", 1}},
		Options: options.Index().SetUnique(true),
	}
	_, err = db.Collection("policies").Indexes().CreateOne(ctx, policiesIdIndexModel)
	if err != nil {
		return err
	}

	policiesNameIndexModel := mongo.IndexModel{
		Keys:    bson.D{{"name", 1}},
		Options: options.Index().SetUnique(true),
	}
	_, err = db.Collection("policies").Indexes().CreateOne(ctx, policiesNameIndexModel)
	if err != nil {
		return err
	}

	return nil
}
