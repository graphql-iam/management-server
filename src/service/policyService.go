package service

import (
	"errors"
	"github.com/google/uuid"
	"github.com/graphql-iam/management-server/src/model"
	"github.com/graphql-iam/management-server/src/model/dao"
	"github.com/graphql-iam/management-server/src/model/requestModel"
	"github.com/graphql-iam/management-server/src/repository"
	"go.mongodb.org/mongo-driver/bson"
)

type PolicyService struct {
	policyRepository *repository.PolicyRepository
}

func NewPolicyService(policyRepository *repository.PolicyRepository) *PolicyService {
	return &PolicyService{policyRepository: policyRepository}
}

func (p *PolicyService) GetPolicies() ([]model.Policy, error) {
	return p.policyRepository.GetAllPolicies()
}

func (p *PolicyService) GetPolicy(id string, name string) (model.Policy, error) {
	if id != "" {
		return p.policyRepository.GetPolicyById(id)
	}
	if name != "" {
		return p.policyRepository.GetPolicyByName(name)
	}
	return model.Policy{}, errors.New("no query provided")
}

func (p *PolicyService) CreatePolicy(policy requestModel.CreatePolicyRequest) (model.Policy, error) {
	newUuid, err := uuid.NewUUID()
	if err != nil {
		return model.Policy{}, err
	}

	var statementsDao []dao.StatementDAO
	for _, statement := range policy.Statements {
		statementDao := dao.StatementDAO{
			Sid:       statement.Sid,
			Action:    statement.Action,
			Effect:    statement.Effect,
			Resource:  statement.Resource,
			Condition: statement.Condition,
		}
		statementsDao = append(statementsDao, statementDao)
	}

	policyDao := dao.PolicyDAO{
		ID:         newUuid.String(),
		Name:       policy.Name,
		Version:    policy.Version,
		Statements: statementsDao,
	}

	err = p.policyRepository.CreatePolicy(policyDao)
	if err != nil {
		return model.Policy{}, err
	}
	return p.GetPolicy(newUuid.String(), "")
}

func (p *PolicyService) UpdatePolicy(update requestModel.UpdatePolicyRequest) (model.Policy, error) {
	updates := bson.M{}

	if update.Name != nil {
		updates["name"] = *update.Name
	}
	if update.Version != nil {
		updates["version"] = *update.Version
	}
	if update.Statements != nil {
		updates["statements"] = update.Statements
	}

	err := p.policyRepository.UpdatePolicy(update.ID, bson.M{"$set": updates})
	if err != nil {
		return model.Policy{}, err
	}

	return p.GetPolicy(update.ID, "")
}

func (p *PolicyService) DeletePolicy(id string) error {
	return p.policyRepository.DeletePolicy(id)
}
