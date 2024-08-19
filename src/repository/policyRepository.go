package repository

import (
	"context"
	"github.com/graphql-iam/management-server/src/model"
	"github.com/graphql-iam/management-server/src/model/dao"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type PolicyRepository struct {
	db *mongo.Database
}

func NewPolicyRepository(db *mongo.Database) *PolicyRepository {
	return &PolicyRepository{db: db}
}

func (p *PolicyRepository) GetAllPolicies() ([]model.Policy, error) {
	var result []model.Policy
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cur, err := p.db.Collection("policies").Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	err = cur.All(ctx, &result)
	return result, err
}

func (p *PolicyRepository) GetPolicyById(id string) (model.Policy, error) {
	var result model.Policy
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := p.db.Collection("policies").FindOne(ctx, bson.D{{"id", id}}).Decode(&result)
	return result, err
}

func (p *PolicyRepository) GetPolicyByName(name string) (model.Policy, error) {
	var result model.Policy
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := p.db.Collection("policies").FindOne(ctx, bson.D{{"name", name}}).Decode(&result)
	return result, err
}

func (p *PolicyRepository) GetPoliciesByStatementAttributes(statement dao.StatementDAO) ([]model.Policy, error) {
	var result []model.Policy
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cur, err := p.db.Collection("policies").Find(ctx, statement)
	if err != nil {
		return result, err
	}
	err = cur.All(ctx, &result)
	return result, err
}

func (p *PolicyRepository) CreatePolicy(policy dao.PolicyDAO) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := p.db.Collection("policies").InsertOne(ctx, policy)
	return err
}

func (p *PolicyRepository) DeletePolicy(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := p.db.Collection("policies").DeleteOne(ctx, bson.D{{"id", id}})
	return err
}

func (p *PolicyRepository) UpdatePolicy(id string, updates interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{{"id", id}}
	_, err := p.db.Collection("policies").UpdateOne(ctx, filter, updates)
	return err
}
