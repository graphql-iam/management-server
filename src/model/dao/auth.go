package dao

import "go.mongodb.org/mongo-driver/bson/primitive"

type RoleDAO struct {
	ObjId     primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name,omitempty"`
	PolicyIds []string           `bson:"policyIds,omitempty"`
}

type PolicyDAO struct {
	ObjId      primitive.ObjectID `bson:"_id,omitempty"`
	ID         string             `bson:"id,omitempty"`
	Name       string             `bson:"name,omitempty"`
	Version    string             `bson:"version,omitempty"`
	Statements []StatementDAO     `bson:"statements,omitempty"`
}

type StatementDAO struct {
	ObjId     primitive.ObjectID           `bson:"_id,omitempty"`
	Sid       string                       `json:"sid,omitempty"`
	Action    string                       `json:"action,omitempty"`
	Effect    string                       `json:"effect,omitempty"`
	Resource  string                       `json:"resource,omitempty"`
	Condition map[string]map[string]string `json:"condition,omitempty"`
}
