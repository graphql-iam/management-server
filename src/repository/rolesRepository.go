package repository

import (
	"context"
	"fmt"
	"github.com/graphql-iam/management-server/src/model"
	"github.com/graphql-iam/management-server/src/model/dao"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type RolesRepository struct {
	db *mongo.Database
}

func NewRolesRepository(db *mongo.Database) *RolesRepository {
	return &RolesRepository{db: db}
}

func (r *RolesRepository) GetRoleByName(name string) (model.Role, error) {
	var result model.Role
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := r.db.Collection("rolesWithPolicies").FindOne(ctx, bson.D{{"name", name}}).Decode(&result)
	return result, err
}

func (r *RolesRepository) GetAllRoles() ([]model.Role, error) {
	var result []model.Role
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cur, err := r.db.Collection("rolesWithPolicies").Find(ctx, bson.D{})
	if err != nil {
		fmt.Println(err.Error())
		return result, err
	}
	err = cur.All(ctx, &result)
	return result, err
}

func (r *RolesRepository) CreateRole(role dao.RoleDAO) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := r.db.Collection("roles").InsertOne(ctx, role)
	return err
}

func (r *RolesRepository) DeleteRole(name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := r.db.Collection("roles").DeleteOne(ctx, bson.D{{"name", name}})
	return err
}

func (r *RolesRepository) AttachPolicyToRole(policyId string, roleName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{{"name", roleName}}
	update := bson.M{"$push": bson.D{{"policyIds", policyId}}}
	_, err := r.db.Collection("roles").UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (r *RolesRepository) DetachPolicyFromRole(policyId string, roleName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{{"name", roleName}}
	update := bson.D{{"$pull", bson.D{{"policyIds", policyId}}}}
	_, err := r.db.Collection("roles").UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
