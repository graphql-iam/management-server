package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/graphql-iam/management-server/src/model/requestModel"
	"github.com/graphql-iam/management-server/src/service"
	"log"
	"net/http"
)

type RolesHandler struct {
	rolesService *service.RolesService
}

func NewRolesHandler(rolesService *service.RolesService) RolesHandler {
	return RolesHandler{
		rolesService: rolesService,
	}
}

func (r *RolesHandler) GetRole(c *gin.Context) {
	name, found := c.GetQuery("name")
	if !found {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	role, err := r.rolesService.GetRole(name)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, role)
}

func (r *RolesHandler) GetRoles(c *gin.Context) {
	roles, err := r.rolesService.GetRoles()
	if err != nil {
		log.Printf("Error getting roles: %v\n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, roles)
}

func (r *RolesHandler) CreateRole(c *gin.Context) {
	var body requestModel.CreateRoleRequest
	err := json.NewDecoder(c.Request.Body).Decode(&body)
	if err != nil || body.Name == "" {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	role, err := r.rolesService.CreateRole(body.Name, body.AttachedPolicies)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, role)
}

func (r *RolesHandler) DeleteRole(c *gin.Context) {
	name, found := c.GetQuery("name")
	if !found {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err := r.rolesService.DeleteRole(name)
	if err != nil {
		log.Printf("Error deleting role %s: %v\n", name, err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func (r *RolesHandler) AttachPolicyToRole(c *gin.Context) {
	var body requestModel.PolicyToRoleRequest
	err := json.NewDecoder(c.Request.Body).Decode(&body)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	role, err := r.rolesService.AttachPolicyToRole(body.PolicyId, body.RoleName)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, role)
}

func (r *RolesHandler) DetachPolicyFromRole(c *gin.Context) {
	var body requestModel.PolicyToRoleRequest
	err := json.NewDecoder(c.Request.Body).Decode(&body)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	role, err := r.rolesService.DetachPolicyFromRole(body.PolicyId, body.RoleName)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, role)
}
