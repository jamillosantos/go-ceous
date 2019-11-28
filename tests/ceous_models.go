// Code generated by https://github.com/jamillosantos/go-ceous DO NOT EDIT

package tests

import (
	"github.com/jamillosantos/go-ceous"
	"github.com/pkg/errors"
)

// GetID returns the primary key for model `User`.
func (model *User) GetID() interface{} {
	return model.ID
}

// ColumnAddress returns the pointer address of a field given its column name.
func (model *User) ColumnAddress(name string) (interface{}, error) {
	switch name {
	case "id":
		return &model.ID, nil
	case "name":
		return &model.Name, nil
	case "password":
		return &model.Password, nil
	case "role":
		return &model.Role, nil
	case "Address_street":
		return &model.Address.Street, nil
	case "Address_number":
		return &model.Address.Number, nil
	case "Address_city":
		return &model.Address.City, nil
	case "Address_state":
		return &model.Address.State, nil
	case "work_street":
		return &model.Work.Street, nil
	case "work_number":
		return &model.Work.Number, nil
	case "work_city":
		return &model.Work.City, nil
	case "work_state":
		return &model.Work.State, nil
	case "created_at":
		return &model.CreatedAt, nil
	case "updated_at":
		return &model.UpdatedAt, nil
	default:
		return nil, errors.Wrapf(ceous.ErrFieldNotFound, "field %s not found", name)
	}
}

// Value returns the value from a field given its column name.
func (model *User) Value(name string) (interface{}, error) {
	switch name {
	case "id":
		return model.ID, nil
	case "name":
		return model.Name, nil
	case "password":
		return model.Password, nil
	case "role":
		return model.Role, nil
	case "Address_street":
		return model.Address.Street, nil
	case "Address_number":
		return model.Address.Number, nil
	case "Address_city":
		return model.Address.City, nil
	case "Address_state":
		return model.Address.State, nil
	case "work_street":
		return model.Work.Street, nil
	case "work_number":
		return model.Work.Number, nil
	case "work_city":
		return model.Work.City, nil
	case "work_state":
		return model.Work.State, nil
	case "created_at":
		return model.CreatedAt, nil
	case "updated_at":
		return model.UpdatedAt, nil
	default:
		return nil, errors.Wrapf(ceous.ErrFieldNotFound, "field %s not found", name)
	}
}

// GetID returns the primary key for model `Group`.
func (model *Group) GetID() interface{} {
	return model.ID
}

// ColumnAddress returns the pointer address of a field given its column name.
func (model *Group) ColumnAddress(name string) (interface{}, error) {
	switch name {
	case "id":
		return &model.ID, nil
	case "name":
		return &model.Name, nil
	default:
		return nil, errors.Wrapf(ceous.ErrFieldNotFound, "field %s not found", name)
	}
}

// Value returns the value from a field given its column name.
func (model *Group) Value(name string) (interface{}, error) {
	switch name {
	case "id":
		return model.ID, nil
	case "name":
		return model.Name, nil
	default:
		return nil, errors.Wrapf(ceous.ErrFieldNotFound, "field %s not found", name)
	}
}

// GetID returns the primary key for model `UserGroup`.
func (model *UserGroup) GetID() interface{} {
	return model.ID
}

// ColumnAddress returns the pointer address of a field given its column name.
func (model *UserGroup) ColumnAddress(name string) (interface{}, error) {
	switch name {
	case "admin":
		return &model.Admin, nil
	default:
		return nil, errors.Wrapf(ceous.ErrFieldNotFound, "field %s not found", name)
	}
}

// Value returns the value from a field given its column name.
func (model *UserGroup) Value(name string) (interface{}, error) {
	switch name {
	case "admin":
		return model.Admin, nil
	default:
		return nil, errors.Wrapf(ceous.ErrFieldNotFound, "field %s not found", name)
	}
}

type schema struct {
	User      *schemaUser
	Group     *schemaGroup
	UserGroup *schemaUserGroup
}

// schemaUser have all fields for the model User.
type schemaUser struct {
	*ceous.BaseSchema

	ID        ceous.SchemaField
	Name      ceous.SchemaField
	Password  ceous.SchemaField
	Role      ceous.SchemaField
	Address   schemaUserAddressAddress
	Work      schemaUserWorkAddress
	CreatedAt ceous.SchemaField
	UpdatedAt ceous.SchemaField
}

// schemaUserAddressAddress have all fields for the model UserAddressAddress.
type schemaUserAddressAddress struct {
	Street ceous.SchemaField
	Number ceous.SchemaField
	City   ceous.SchemaField
	State  ceous.SchemaField
}

// schemaUserWorkAddress have all fields for the model UserWorkAddress.
type schemaUserWorkAddress struct {
	Street ceous.SchemaField
	Number ceous.SchemaField
	City   ceous.SchemaField
	State  ceous.SchemaField
}

// schemaGroup have all fields for the model Group.
type schemaGroup struct {
	*ceous.BaseSchema

	ID   ceous.SchemaField
	Name ceous.SchemaField
}

// schemaUserGroup have all fields for the model UserGroup.
type schemaUserGroup struct {
	*ceous.BaseSchema

	Admin ceous.SchemaField
}

var (
	baseSchemaUserAddressAddress = schemaUserAddressAddress{

		Street: baseSchemaUser.ColumnsArr[4],

		Number: baseSchemaUser.ColumnsArr[5],

		City: baseSchemaUser.ColumnsArr[6],

		State: baseSchemaUser.ColumnsArr[7],
	}

	baseSchemaUserWorkAddress = schemaUserWorkAddress{

		Street: baseSchemaUser.ColumnsArr[8],

		Number: baseSchemaUser.ColumnsArr[9],

		City: baseSchemaUser.ColumnsArr[10],

		State: baseSchemaUser.ColumnsArr[11],
	}
)

// Schema represents the schema of the package "tests".
var Schema = schema{
	User: &schemaUser{
		BaseSchema: baseSchemaUser,

		ID:        baseSchemaUser.ColumnsArr[0],
		Name:      baseSchemaUser.ColumnsArr[1],
		Password:  baseSchemaUser.ColumnsArr[2],
		Role:      baseSchemaUser.ColumnsArr[3],
		Address:   baseSchemaUserAddressAddress,
		Work:      baseSchemaUserWorkAddress,
		CreatedAt: baseSchemaUser.ColumnsArr[12],
		UpdatedAt: baseSchemaUser.ColumnsArr[13],
	},
	Group: &schemaGroup{
		BaseSchema: baseSchemaGroup,

		ID:   baseSchemaGroup.ColumnsArr[0],
		Name: baseSchemaGroup.ColumnsArr[1],
	},
	UserGroup: &schemaUserGroup{
		BaseSchema: baseSchemaUserGroup,

		Admin: baseSchemaUserGroup.ColumnsArr[0],
	},
}

var (
	baseSchemaUser = ceous.NewBaseSchema(
		"",
		"",
		ceous.NewSchemaField("id"),
		ceous.NewSchemaField("name"),
		ceous.NewSchemaField("password"),
		ceous.NewSchemaField("role"),
		ceous.NewSchemaField("Address_street"),
		ceous.NewSchemaField("Address_number"),
		ceous.NewSchemaField("Address_city"),
		ceous.NewSchemaField("Address_state"),
		ceous.NewSchemaField("work_street"),
		ceous.NewSchemaField("work_number"),
		ceous.NewSchemaField("work_city"),
		ceous.NewSchemaField("work_state"),
		ceous.NewSchemaField("created_at"),
		ceous.NewSchemaField("updated_at"),
	)

	baseSchemaGroup = ceous.NewBaseSchema(
		"",
		"",
		ceous.NewSchemaField("id"),
		ceous.NewSchemaField("name"),
	)

	baseSchemaUserGroup = ceous.NewBaseSchema(
		"",
		"",
		ceous.NewSchemaField("admin"),
	)
)
