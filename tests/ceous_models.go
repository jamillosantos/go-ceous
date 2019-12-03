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

// GetID returns the primary key for model `Group`.
func (model *Group) GetID() interface{} {
	return model.ID
}

// GetID returns the primary key for model `UserGroup`.
func (model *UserGroup) GetID() interface{} {
	return model.ID
}

// User returns the User from UserGroup.
func (model *UserGroup) User() *User {
	return model.user
}

// SetUser updates the value for the user,
// updating the user.
func (model *UserGroup) SetUser(value *User) error {
	localPkPtr, err := model.ColumnAddress("id")
	if err != nil {
		return err
	}

	localFKTypedPtr, ok := localPkPtr.(*int)
	if !ok {
		return errors.New("invalid key type") // TODO(jota): To formalize this error.
	}
	*localFKTypedPtr = value.ID
	model.user = value
	return nil
}

// Value returns the value from a field given its column name.
func (model *UserGroupPK) Value(name string) (interface{}, error) {
	switch name {
	case "user_id":
		return model.UserID, nil
	case "group_id":
		return model.GroupID, nil
	default:
		return nil, errors.Wrapf(ceous.ErrFieldNotFound, "field %s not found", name)
	}
}

// ColumnAddress returns the pointer address of a field given its column name.
func (model *UserGroupPK) ColumnAddress(name string) (interface{}, error) {
	switch name {
	case "user_id":
		return &model.UserID, nil
	case "group_id":
		return &model.GroupID, nil
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
	case "street":
		return model.Address.Street, nil
	case "number":
		return model.Address.Number, nil
	case "city":
		return model.Address.City, nil
	case "state":
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
	case "street":
		return &model.Address.Street, nil
	case "number":
		return &model.Address.Number, nil
	case "city":
		return &model.Address.City, nil
	case "state":
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
func (model *UserGroup) Value(name string) (interface{}, error) {
	switch name {
	case "user_id":
		return model.ID.UserID, nil
	case "group_id":
		return model.ID.GroupID, nil
	case "admin":
		return model.Admin, nil
	default:
		return nil, errors.Wrapf(ceous.ErrFieldNotFound, "field %s not found", name)
	}
}

// ColumnAddress returns the pointer address of a field given its column name.
func (model *UserGroup) ColumnAddress(name string) (interface{}, error) {
	switch name {
	case "user_id":
		return &model.ID.UserID, nil
	case "group_id":
		return &model.ID.GroupID, nil
	case "admin":
		return &model.Admin, nil
	default:
		return nil, errors.Wrapf(ceous.ErrFieldNotFound, "field %s not found", name)
	}
}

type schema struct {
	UserGroupPK *schemaUserGroupPK
	User        *schemaUser
	Group       *schemaGroup
	UserGroup   *schemaUserGroup
}

// schemaUserGroupPK have all fields for the model UserGroupPK.
type schemaUserGroupPK struct {
	*ceous.BaseSchema

	UserID  ceous.SchemaField
	GroupID ceous.SchemaField
}

// schemaUserAddress have all fields for the model UserAddress.
type schemaUserAddress struct {
	Street ceous.SchemaField
	Number ceous.SchemaField
	City   ceous.SchemaField
	State  ceous.SchemaField
}

// schemaUserWork have all fields for the model UserWork.
type schemaUserWork struct {
	Street ceous.SchemaField
	Number ceous.SchemaField
	City   ceous.SchemaField
	State  ceous.SchemaField
}

// schemaUser have all fields for the model User.
type schemaUser struct {
	*ceous.BaseSchema

	ID        ceous.SchemaField
	Name      ceous.SchemaField
	Password  ceous.SchemaField
	Role      ceous.SchemaField
	Address   schemaUserAddress
	Work      schemaUserWork
	CreatedAt ceous.SchemaField
	UpdatedAt ceous.SchemaField
}

// schemaGroup have all fields for the model Group.
type schemaGroup struct {
	*ceous.BaseSchema

	ID   ceous.SchemaField
	Name ceous.SchemaField
}

// schemaUserGroupID have all fields for the model UserGroupID.
type schemaUserGroupID struct {
	UserID  ceous.SchemaField
	GroupID ceous.SchemaField
}

// schemaUserGroup have all fields for the model UserGroup.
type schemaUserGroup struct {
	*ceous.BaseSchema

	ID    schemaUserGroupID
	Admin ceous.SchemaField
}

var (
	baseSchemaUserGroupPK = ceous.NewBaseSchema(
		"",
		"",
		ceous.NewSchemaField("user_id"),
		ceous.NewSchemaField("group_id"),
	)

	baseSchemaUser = ceous.NewBaseSchema(
		"users",
		"",
		ceous.NewSchemaField("id"),
		ceous.NewSchemaField("name"),
		ceous.NewSchemaField("password"),
		ceous.NewSchemaField("role"),
		ceous.NewSchemaField("street"),
		ceous.NewSchemaField("number"),
		ceous.NewSchemaField("city"),
		ceous.NewSchemaField("state"),
		ceous.NewSchemaField("work_street"),
		ceous.NewSchemaField("work_number"),
		ceous.NewSchemaField("work_city"),
		ceous.NewSchemaField("work_state"),
		ceous.NewSchemaField("created_at"),
		ceous.NewSchemaField("updated_at"),
	)

	baseSchemaGroup = ceous.NewBaseSchema(
		"groups",
		"",
		ceous.NewSchemaField("id"),
		ceous.NewSchemaField("name"),
	)

	baseSchemaUserGroup = ceous.NewBaseSchema(
		"user_groups",
		"",
		ceous.NewSchemaField("user_id"),
		ceous.NewSchemaField("group_id"),
		ceous.NewSchemaField("admin"),
	)
)

var (
	baseSchemaUserAddress = schemaUserAddress{

		Street: baseSchemaUser.ColumnsArr[4],
		Number: baseSchemaUser.ColumnsArr[5],
		City:   baseSchemaUser.ColumnsArr[6],
		State:  baseSchemaUser.ColumnsArr[7],
	}
	baseSchemaUserWork = schemaUserWork{

		Street: baseSchemaUser.ColumnsArr[8],
		Number: baseSchemaUser.ColumnsArr[9],
		City:   baseSchemaUser.ColumnsArr[10],
		State:  baseSchemaUser.ColumnsArr[11],
	}
	baseSchemaUserGroupID = schemaUserGroupID{

		UserID:  baseSchemaUserGroup.ColumnsArr[0],
		GroupID: baseSchemaUserGroup.ColumnsArr[1],
	}
)

// Schema represents the schema of the package "tests".
var Schema = schema{

	UserGroupPK: &schemaUserGroupPK{
		BaseSchema: baseSchemaUserGroupPK,

		UserID:  baseSchemaUserGroupPK.ColumnsArr[0],
		GroupID: baseSchemaUserGroupPK.ColumnsArr[1],
	},

	User: &schemaUser{
		BaseSchema: baseSchemaUser,

		ID:        baseSchemaUser.ColumnsArr[0],
		Name:      baseSchemaUser.ColumnsArr[1],
		Password:  baseSchemaUser.ColumnsArr[2],
		Role:      baseSchemaUser.ColumnsArr[3],
		Address:   baseSchemaUserAddress,
		Work:      baseSchemaUserWork,
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

		ID:    baseSchemaUserGroupID,
		Admin: baseSchemaUserGroup.ColumnsArr[2],
	},
}
