package enum

const (
	UserStatusInactive UserStatus = iota
	UserStatusActive
	UserStatusDisabled
)

type UserStatus int

func (s UserStatus) String() string {
	statuses := []string{"Inactive", "Active", "Disabled"}
	if int(s) < len(statuses) {
		return statuses[s]
	}
	return "Unknown"
}

func (s UserStatus) IsActive() bool {
	return s == UserStatusActive
}

const (
	UserRoleUser UserRole = iota
	UserRoleAdmin
	UserRoleSuperAdmin
)

type UserRole int

func (r UserRole) String() string {
	roles := []string{"ROLE_USER", "ROLE_ADMIN", "ROLE_SUPER_ADMIN"}
	if int(r) < len(roles) {
		return roles[r]
	}
	return "UNKNOWN"
}

func (r UserRole) IsAdmin() bool {
	return r == UserRoleAdmin || r.IsSuperAdmin()
}

func (r UserRole) IsSuperAdmin() bool {
	return r == UserRoleSuperAdmin
}
