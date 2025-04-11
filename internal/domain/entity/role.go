package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Permission string

const (
	PermissionUserRead       Permission = "user:read"
	PermissionUserWrite      Permission = "user:write"
	PermissionUserDelete     Permission = "user:delete"
	PermissionUserBlock      Permission = "user:block"
	PermissionUserUnblock    Permission = "user:unblock"
	PermissionRoleManage     Permission = "role:manage"
	PermissionAdminPanel     Permission = "admin:panel"
	PermissionAuthMethodEdit Permission = "auth:method:edit"
)

type Role struct {
	ID          primitive.ObjectID `bson:"_id, omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	Permissions []Permission       `bson:"permissions"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

// Predefined roles
var PredefinedRoles = map[string][]Permission{
	"admin": {
		PermissionUserRead,
		PermissionUserWrite,
		PermissionUserDelete,
		PermissionUserBlock,
		PermissionUserUnblock,
		PermissionRoleManage,
		PermissionAdminPanel,
		PermissionAuthMethodEdit,
	},
	"manager": {
		PermissionUserRead,
		PermissionUserWrite,
		PermissionUserBlock,
		PermissionUserUnblock,
		PermissionAdminPanel,
	},
	"user": {
		PermissionUserRead,
	},
}

// new roles if needed
func NewRole(name, description string, permissions []Permission) *Role {
	return &Role{
		Name:        name,
		Description: description,
		Permissions: permissions,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (r *Role) HasPermission(permission Permission) bool {
	for _, p := range r.Permissions {
		if p == permission {
			return true
		}
	}
	return false
}

// AddPermission добавляет разрешение к роли
func (r *Role) AddPermission(permission Permission) {
	// Проверяем, существует ли уже такое разрешение
	for _, p := range r.Permissions {
		if p == permission {
			return
		}
	}
	r.Permissions = append(r.Permissions, permission)
	r.UpdatedAt = time.Now()
}

// RemovePermission удаляет разрешение из роли
func (r *Role) RemovePermission(permission Permission) {
	var newPermissions []Permission
	for _, p := range r.Permissions {
		if p != permission {
			newPermissions = append(newPermissions, p)
		}
	}
	r.Permissions = newPermissions
	r.UpdatedAt = time.Now()
}
