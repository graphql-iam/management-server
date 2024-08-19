package service

import (
	"github.com/graphql-iam/management-server/src/model"
	"github.com/graphql-iam/management-server/src/model/dao"
	"github.com/graphql-iam/management-server/src/repository"
)

type RolesService struct {
	rolesRepository *repository.RolesRepository
}

func NewRolesService(rolesRepository *repository.RolesRepository) *RolesService {
	return &RolesService{
		rolesRepository: rolesRepository,
	}
}

func (r *RolesService) GetRoles() ([]model.Role, error) {
	return r.rolesRepository.GetAllRoles()
}

func (r *RolesService) GetRole(name string) (model.Role, error) {
	return r.rolesRepository.GetRoleByName(name)
}

func (r *RolesService) CreateRole(name string, attachedPolicies []string) (model.Role, error) {
	if attachedPolicies == nil {
		attachedPolicies = []string{}
	}

	role := dao.RoleDAO{
		Name:      name,
		PolicyIds: attachedPolicies,
	}
	err := r.rolesRepository.CreateRole(role)
	if err != nil {
		return model.Role{}, err
	}
	return r.rolesRepository.GetRoleByName(name)
}

func (r *RolesService) DeleteRole(name string) error {
	return r.rolesRepository.DeleteRole(name)
}

func (r *RolesService) AttachPolicyToRole(policyId string, roleName string) (model.Role, error) {
	err := r.rolesRepository.AttachPolicyToRole(policyId, roleName)
	if err != nil {
		return model.Role{}, err
	}
	return r.GetRole(roleName)
}

func (r *RolesService) DetachPolicyFromRole(policyId string, roleName string) (model.Role, error) {
	err := r.rolesRepository.DetachPolicyFromRole(policyId, roleName)
	if err != nil {
		return model.Role{}, err
	}
	return r.GetRole(roleName)
}
