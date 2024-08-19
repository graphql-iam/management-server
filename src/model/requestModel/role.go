package requestModel

type CreateRoleRequest struct {
	Name             string   `json:"name"`
	AttachedPolicies []string `json:"attachedPolicies,omitempty"`
}

type PolicyToRoleRequest struct {
	RoleName string `json:"roleName"`
	PolicyId string `json:"policyId"`
}
