// Package role adjasdlkjadslk asd.
package role

import (
	"fmt"
)

var (
	Admin = newRole("ADMIN")
	User  = newRole("USER")
)

// =============================================================================

var roles = make(map[string]Role)

type Role struct {
	value string
}

func newRole(role string) Role {
	r := Role{role}
	roles[role] = r
	return r
}

func (r Role) String() string {
	return r.value
}

func (r Role) Equal(r2 Role) bool {
	return r.value == r2.value
}

func (r Role) MarshalText() ([]byte, error) {
	return []byte(r.value), nil
}

// =============================================================================

func Parse(value string) (Role, error) {
	role, exists := roles[value]
	if !exists {
		return Role{}, fmt.Errorf("invalid role %q", value)
	}

	return role, nil
}

func MustParse(value string) Role {
	role, err := Parse(value)
	if err != nil {
		panic(err)
	}

	return role
}
