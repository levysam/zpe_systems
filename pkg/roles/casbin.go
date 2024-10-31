package roles

import (
	"github.com/casbin/casbin/v2"
	"strings"
)

type CasbinRule struct {
	enforcer *casbin.Enforcer
}

func NewCasbinRule(enforcer *casbin.Enforcer) *CasbinRule {
	return &CasbinRule{
		enforcer: enforcer,
	}
}

func (c *CasbinRule) CheckPermission(user string, resource string, action string) (allowed bool, err error) {
	return c.enforcer.Enforce(user, resource, action)
}

func (c *CasbinRule) SetRoleToUser(user string, role string) (changed bool, err error) {
	return c.enforcer.AddRoleForUser(user, role)
}

func (c *CasbinRule) SetRoleToUserBatch(user string, roles []string) (changed bool, err error) {
	return c.enforcer.AddRolesForUser(user, roles)
}

func (c *CasbinRule) DeleteRoleFromUser(user string, role string) (changed bool, err error) {
	return c.enforcer.DeleteRoleForUser(user, role)
}

func (c *CasbinRule) ListUserRoles(user string) (roles []string, err error) {
	rolesList, err := c.enforcer.GetRolesForUser(user)
	if err != nil {
		return nil, err
	}
	for _, role := range rolesList {
		roleSplit := strings.Split(role, "_")
		roles = append(roles, roleSplit[1])
	}
	return roles, nil
}

func (c *CasbinRule) AddPermissionToRole(role, resource, action string) (changed bool, err error) {
	return c.enforcer.AddPolicy(role, resource, action)
}

func (c *CasbinRule) DeletePermissionFromRole(role, resource string) (changed bool, err error) {
	return c.enforcer.RemovePolicy(role, resource)
}

func (c *CasbinRule) DeleteRole(role string) (changed bool, err error) {
	return c.enforcer.DeleteRole(role)
}

func (c *CasbinRule) ListRoles() (roles []string, err error) {
	rolesList, err := c.enforcer.GetAllSubjects()
	if err != nil {
		return nil, err
	}
	for _, role := range rolesList {
		roleSplit := strings.Split(role, "_")
		roles = append(roles, roleSplit[1])
	}
	return roles, nil
}
