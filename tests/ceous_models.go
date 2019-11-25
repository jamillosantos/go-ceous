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

// User returns the user from UserGroup.
func (model *UserGroup) User() *User {
	return model.user
}

// SetUser updates the value for the user,
// updating the user.
func (model *UserGroup) SetUser(value *User) error {
	c, err := model.ColumnAddress("user_id")
	if err != nil {
		return err
	}

	v, ok := c.(*int)
	if !ok {
		return errors.New("invalid key value") // TODO(jota): To formalize this error.
	}
	*v = value.ID
	model.user = value
	return nil
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

type schemaUserGroupPK struct {
	UserID  ceous.SchemaField
	GroupID ceous.SchemaField
}

var UserGroupPKSchema = schemaUserGroupPK{
	UserID:  ceous.NewSchemaField("user_id"),
	GroupID: ceous.NewSchemaField("group_id"),
}

// ColumnAddress returns the pointer address of a field given its column name.
func (model *Address) ColumnAddress(name string) (interface{}, error) {
	switch name {
	case "street":
		return &model.Street, nil
	case "number":
		return &model.Number, nil
	case "city":
		return &model.City, nil
	case "state":
		return &model.State, nil
	default:
		return nil, errors.Wrapf(ceous.ErrFieldNotFound, "field %s not found", name)
	}
}

// Value returns the value from a field given its column name.
func (model *Address) Value(name string) (interface{}, error) {
	switch name {
	case "street":
		return model.Street, nil
	case "number":
		return model.Number, nil
	case "city":
		return model.City, nil
	case "state":
		return model.State, nil
	default:
		return nil, errors.Wrapf(ceous.ErrFieldNotFound, "field %s not found", name)
	}
}

type schemaAddress struct {
	Street ceous.SchemaField
	Number ceous.SchemaField
	City   ceous.SchemaField
	State  ceous.SchemaField
}

var AddressSchema = schemaAddress{
	Street: ceous.NewSchemaField("street"),
	Number: ceous.NewSchemaField("number"),
	City:   ceous.NewSchemaField("city"),
	State:  ceous.NewSchemaField("state"),
}

type schema struct {
	User      *schemaUser
	Group     *schemaGroup
	UserGroup *schemaUserGroup
}

// Schema represents the schema of the package "tests".
var Schema = schema{
	User: &schemaUser{
		BaseSchema: baseSchemaUser,

		ID: baseSchemaUser.ColumnsArr[0],

		Name: baseSchemaUser.ColumnsArr[1],

		Password: baseSchemaUser.ColumnsArr[2],

		Role: baseSchemaUser.ColumnsArr[3],

		Address: AddressSchema,

		Work: AddressSchema,

		CreatedAt: baseSchemaUser.ColumnsArr[12],

		UpdatedAt: baseSchemaUser.ColumnsArr[13],
	},
	Group: &schemaGroup{
		BaseSchema: baseSchemaGroup,

		ID: baseSchemaGroup.ColumnsArr[0],

		Name: baseSchemaGroup.ColumnsArr[1],
	},
	UserGroup: &schemaUserGroup{
		BaseSchema: baseSchemaUserGroup,

		ID: UserGroupPKSchema,

		Admin: baseSchemaUserGroup.ColumnsArr[2],
	},
}
var baseSchemaUser = ceous.NewBaseSchema(
	"users",
	"",
	ceous.NewSchemaField("id", ceous.FieldPK, ceous.FieldAutoIncrement),
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

type schemaUser struct {
	*ceous.BaseSchema
	ID        ceous.SchemaField
	Name      ceous.SchemaField
	Password  ceous.SchemaField
	Role      ceous.SchemaField
	Address   schemaAddress
	Work      schemaAddress
	CreatedAt ceous.SchemaField
	UpdatedAt ceous.SchemaField
}

var baseSchemaGroup = ceous.NewBaseSchema(
	"groups",
	"",
	ceous.NewSchemaField("id", ceous.FieldPK, ceous.FieldAutoIncrement),
	ceous.NewSchemaField("name"),
)

type schemaGroup struct {
	*ceous.BaseSchema
	ID   ceous.SchemaField
	Name ceous.SchemaField
}

var baseSchemaUserGroup = ceous.NewBaseSchema(
	"user_groups",
	"",
	ceous.NewSchemaField("user_id", ceous.FieldPK),
	ceous.NewSchemaField("group_id", ceous.FieldPK),
	ceous.NewSchemaField("admin"),
)

type schemaUserGroup struct {
	*ceous.BaseSchema
	ID    schemaUserGroupPK
	Admin ceous.SchemaField
}
