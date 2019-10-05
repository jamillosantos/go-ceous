package tests

import (
	"fmt"
	"time"

	"github.com/jamillosantos/go-ceous"
)

type User struct {
	ceous.Model
	ID        int
	Name      string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) ColumnAddress(name string) (interface{}, error) {
	switch name {
	case "id":
		return &u.ID, nil
	case "name":
		return &u.Name, nil
	case "password":
		return &u.Password, nil
	case "role":
		return &u.Role, nil
	case "created_at":
		return &u.CreatedAt, nil
	case "updated_at":
		return &u.UpdatedAt, nil
	default:
		return nil, fmt.Errorf("field %s not found", name)
	}
}

func (u *User) Value(name string) (interface{}, error) {
	switch name {
	case "id":
		return u.ID, nil
	case "name":
		return u.Name, nil
	case "password":
		return u.Password, nil
	case "role":
		return u.Role, nil
	case "created_at":
		return u.CreatedAt, nil
	case "updated_at":
		return u.UpdatedAt, nil
	default:
		return nil, fmt.Errorf("field %s not found", name)
	}
}
