package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/graphql-iam/management-server/src/model/requestModel"
	"github.com/graphql-iam/management-server/src/service"
	"log"
	"net/http"
)

type PolicyHandler struct {
	policyService *service.PolicyService
}

func NewPolicyHandler(policyService *service.PolicyService) PolicyHandler {
	return PolicyHandler{policyService: policyService}
}

func (p *PolicyHandler) GetPolicy(c *gin.Context) {
	id, idFound := c.GetQuery("id")
	name, nameFound := c.GetQuery("name")
	if !idFound && !nameFound {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	role, err := p.policyService.GetPolicy(id, name)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, role)
}

func (p *PolicyHandler) GetPolicies(c *gin.Context) {
	roles, err := p.policyService.GetPolicies()
	if err != nil {
		log.Printf("Error getting policies: %v\n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, roles)
}

func (p *PolicyHandler) CreatePolicy(c *gin.Context) {
	var body requestModel.CreatePolicyRequest
	err := json.NewDecoder(c.Request.Body).Decode(&body)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	policy, err := p.policyService.CreatePolicy(body)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, policy)
}

func (p *PolicyHandler) UpdatePolicy(c *gin.Context) {
	var body requestModel.UpdatePolicyRequest
	err := json.NewDecoder(c.Request.Body).Decode(&body)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	policy, err := p.policyService.UpdatePolicy(body)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, policy)
}

func (p *PolicyHandler) DeletePolicy(c *gin.Context) {
	id, found := c.GetQuery("id")
	if !found {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err := p.policyService.DeletePolicy(id)
	if err != nil {
		log.Printf("Error delting policy with id %s: %v\n", id, err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}
