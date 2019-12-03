// Code generated by https://github.com/jamillosantos/go-ceous DO NOT EDIT

package db

import (
	"context"
	"database/sql"
	ceous "github.com/jamillosantos/go-ceous"
	tests "github.com/jamillosantos/go-ceous/tests"
	time "time"
)

type Creator interface {
	// UserQuery creates a new query related with the connection set.
	UserQuery(options ...ceous.CeousOption) *UserQuery
	// GroupQuery creates a new query related with the connection set.
	GroupQuery(options ...ceous.CeousOption) *GroupQuery
	// UserGroupQuery creates a new query related with the connection set.
	UserGroupQuery(options ...ceous.CeousOption) *UserGroupQuery
	// UserStore creates a new store related with the connection set.
	UserStore(options ...ceous.CeousOption) *UserStore
	// GroupStore creates a new store related with the connection set.
	GroupStore(options ...ceous.CeousOption) *GroupStore
	// UserGroupStore creates a new store related with the connection set.
	UserGroupStore(options ...ceous.CeousOption) *UserGroupStore
}

type Connection interface {
	ceous.DBRunner
	Creator
}
type defaultConnection struct {
	*ceous.BaseConnection
}

// UserQuery creates a new query related with the connection default set.
func (c *defaultConnection) UserQuery(options ...ceous.CeousOption) *UserQuery {
	return NewUserQuery(append(options, ceous.WithConn(c))...)
}

// GroupQuery creates a new query related with the connection default set.
func (c *defaultConnection) GroupQuery(options ...ceous.CeousOption) *GroupQuery {
	return NewGroupQuery(append(options, ceous.WithConn(c))...)
}

// UserGroupQuery creates a new query related with the connection default set.
func (c *defaultConnection) UserGroupQuery(options ...ceous.CeousOption) *UserGroupQuery {
	return NewUserGroupQuery(append(options, ceous.WithConn(c))...)
}

// UserStore creates a new store related with the connection default set.
func (c *defaultConnection) UserStore(options ...ceous.CeousOption) *UserStore {
	return NewUserStore(append(options, ceous.WithConn(c))...)
}

// GroupStore creates a new store related with the connection default set.
func (c *defaultConnection) GroupStore(options ...ceous.CeousOption) *GroupStore {
	return NewGroupStore(append(options, ceous.WithConn(c))...)
}

// UserGroupStore creates a new store related with the connection default set.
func (c *defaultConnection) UserGroupStore(options ...ceous.CeousOption) *UserGroupStore {
	return NewUserGroupStore(append(options, ceous.WithConn(c))...)
}

// Begin creates a new transaction with default set.
func (c *defaultConnection) Begin() (*Transaction, error) {
	tx, err := c.BaseConnection.Begin()
	if err != nil {
		return nil, err
	}
	return NewTransaction(tx), nil
}

// BeginTx creates a new transaction with extended config params with the
// connection default set.
func (c *defaultConnection) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Transaction, error) {
	tx, err := c.BaseConnection.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}
	return NewTransaction(tx), nil
}

var (
	// Default is a database connection reference.
	Default *defaultConnection
)

// InitDefault initializes the connection `Default:`.
func InitDefault(db *sql.DB) {
	Default = &defaultConnection{
		BaseConnection: ceous.NewConnection(db),
	}
}

type Transaction struct {
	*ceous.BaseTxRunner
}

func NewTransaction(tx *ceous.BaseTxRunner) *Transaction {
	return &Transaction{
		BaseTxRunner: tx,
	}
}

// UserQuery creates a new query from a transaction.
func (c *Transaction) UserQuery(options ...ceous.CeousOption) *UserQuery {
	return NewUserQuery(append(options, ceous.WithRunner(c))...)
}

// GroupQuery creates a new query from a transaction.
func (c *Transaction) GroupQuery(options ...ceous.CeousOption) *GroupQuery {
	return NewGroupQuery(append(options, ceous.WithRunner(c))...)
}

// UserGroupQuery creates a new query from a transaction.
func (c *Transaction) UserGroupQuery(options ...ceous.CeousOption) *UserGroupQuery {
	return NewUserGroupQuery(append(options, ceous.WithRunner(c))...)
}

// UserStore creates a new store from a transaction.
func (c *Transaction) UserStore(options ...ceous.CeousOption) *UserStore {
	return NewUserStore(append(options, ceous.WithRunner(c))...)
}

// GroupStore creates a new store from a transaction.
func (c *Transaction) GroupStore(options ...ceous.CeousOption) *GroupStore {
	return NewGroupStore(append(options, ceous.WithRunner(c))...)
}

// UserGroupStore creates a new store from a transaction.
func (c *Transaction) UserGroupStore(options ...ceous.CeousOption) *UserGroupStore {
	return NewUserGroupStore(append(options, ceous.WithRunner(c))...)
}

// UserQuery is the query for the model `UserQuery`.
type UserQuery struct {
	*ceous.BaseQuery
}

// NewUserQuery creates a new query for model `UserQuery`.
func NewUserQuery(options ...ceous.CeousOption) *UserQuery {
	bq := ceous.NewBaseQuery(options...)
	if bq.Schema == nil {
		bq.Schema = tests.Schema.User.BaseSchema
	}
	return &UserQuery{
		BaseQuery: bq,
	}
}

// ByID add a filter by `ID`.
func (q *UserQuery) ByID(value int) *UserQuery {
	q.BaseQuery.Where(ceous.Eq(tests.Schema.User.ID, value))
	return q
}

// ByName add a filter by `Name`.
func (q *UserQuery) ByName(value string) *UserQuery {
	q.BaseQuery.Where(ceous.Eq(tests.Schema.User.Name, value))
	return q
}

// ByPassword add a filter by `Password`.
func (q *UserQuery) ByPassword(value string) *UserQuery {
	q.BaseQuery.Where(ceous.Eq(tests.Schema.User.Password, value))
	return q
}

// ByRole add a filter by `Role`.
func (q *UserQuery) ByRole(value string) *UserQuery {
	q.BaseQuery.Where(ceous.Eq(tests.Schema.User.Role, value))
	return q
}

// ByStreet add a filter by `Street`.
func (q *UserQuery) ByAddressStreet(value string) *UserQuery {
	q.BaseQuery.Where(ceous.Eq(tests.Schema.User.Address.Street, value))
	return q
}

// ByNumber add a filter by `Number`.
func (q *UserQuery) ByAddressNumber(value string) *UserQuery {
	q.BaseQuery.Where(ceous.Eq(tests.Schema.User.Address.Number, value))
	return q
}

// ByCity add a filter by `City`.
func (q *UserQuery) ByAddressCity(value string) *UserQuery {
	q.BaseQuery.Where(ceous.Eq(tests.Schema.User.Address.City, value))
	return q
}

// ByState add a filter by `State`.
func (q *UserQuery) ByAddressState(value string) *UserQuery {
	q.BaseQuery.Where(ceous.Eq(tests.Schema.User.Address.State, value))
	return q
}

// ByStreet add a filter by `Street`.
func (q *UserQuery) ByWorkStreet(value string) *UserQuery {
	q.BaseQuery.Where(ceous.Eq(tests.Schema.User.Work.Street, value))
	return q
}

// ByNumber add a filter by `Number`.
func (q *UserQuery) ByWorkNumber(value string) *UserQuery {
	q.BaseQuery.Where(ceous.Eq(tests.Schema.User.Work.Number, value))
	return q
}

// ByCity add a filter by `City`.
func (q *UserQuery) ByWorkCity(value string) *UserQuery {
	q.BaseQuery.Where(ceous.Eq(tests.Schema.User.Work.City, value))
	return q
}

// ByState add a filter by `State`.
func (q *UserQuery) ByWorkState(value string) *UserQuery {
	q.BaseQuery.Where(ceous.Eq(tests.Schema.User.Work.State, value))
	return q
}

// ByCreatedAt add a filter by `CreatedAt`.
func (q *UserQuery) ByCreatedAt(value time.Time) *UserQuery {
	q.BaseQuery.Where(ceous.Eq(tests.Schema.User.CreatedAt, value))
	return q
}

// ByUpdatedAt add a filter by `UpdatedAt`.
func (q *UserQuery) ByUpdatedAt(value time.Time) *UserQuery {
	q.BaseQuery.Where(ceous.Eq(tests.Schema.User.UpdatedAt, value))
	return q
}

// Select defines what fields should be selected from the database.
func (q *UserQuery) Select(fields ...ceous.SchemaField) *UserQuery {
	q.BaseQuery.Select(fields...)
	return q
}

// ExcludeFields defines what fields should not be selected from the database.
func (q *UserQuery) ExcludeFields(fields ...ceous.SchemaField) *UserQuery {
	q.BaseQuery.ExcludeFields(fields...)
	return q
}

// Where defines the conditions for
func (q *UserQuery) Where(pred interface{}, args ...interface{}) *UserQuery {
	q.BaseQuery.Where(pred, args...)
	return q
}

func (q *UserQuery) Limit(limit uint64) *UserQuery {
	q.BaseQuery.Limit(limit)
	return q
}

func (q *UserQuery) Offset(offset uint64) *UserQuery {
	q.BaseQuery.Offset(offset)
	return q
}

// One results only one record matching the query.
func (q *UserQuery) One() (m tests.User, err error) {
	q.Limit(1).Offset(0)

	query, err := q.RawQuery()
	if err != nil {
		return
	}

	rs := NewUserResultSet(query, nil)
	defer rs.Close()

	if rs.Next() {
		err = rs.ToModel(&m)
		if err != nil {
			return
		}

		for _, rel := range q.BaseQuery.Relations {
			err = rel.Aggregate(&m)
			if err != nil {
				return tests.User{}, err // TODO(jota): Shall this error be wrapped? At first, yes.
			}
		}
	} else {
		err = ceous.ErrNotFound
	}

	if err == nil {
		for _, rel := range q.BaseQuery.Relations {
			err = rel.Realize()
			if err != nil {
				return tests.User{}, err // TODO(jota): Shall this error be wrapped? At first, yes.
			}
		}
	}

	return
}

// All return all records that match the query.
func (q *UserQuery) All() ([]*tests.User, error) {
	query, err := q.RawQuery()
	if err != nil {
		return nil, err
	}

	rs := NewUserResultSet(query, nil)
	defer rs.Close()

	ms := make([]*tests.User, 0)
	for rs.Next() {
		m := &tests.User{}
		err = rs.ToModel(m)
		if err != nil {
			return nil, err
		}

		for _, rel := range q.BaseQuery.Relations {
			err = rel.Aggregate(m)
			if err != nil {
				return nil, err // TODO(jota): Shall this error be wrapped? At first, yes.
			}
		}
		ms = append(ms, m)
	}

	for _, rel := range q.BaseQuery.Relations {
		err = rel.Realize()
		if err != nil {
			return nil, err // TODO(jota): Shall this error be wrapped? At first, yes.
		}
	}

	return ms, nil
}

func (q *UserQuery) OrderBy(fields ...interface{}) *UserQuery {
	q.BaseQuery.OrderBy(fields...)
	return q
}

// GroupQuery is the query for the model `GroupQuery`.
type GroupQuery struct {
	*ceous.BaseQuery
}

// NewGroupQuery creates a new query for model `GroupQuery`.
func NewGroupQuery(options ...ceous.CeousOption) *GroupQuery {
	bq := ceous.NewBaseQuery(options...)
	if bq.Schema == nil {
		bq.Schema = tests.Schema.Group.BaseSchema
	}
	return &GroupQuery{
		BaseQuery: bq,
	}
}

// ByID add a filter by `ID`.
func (q *GroupQuery) ByID(value int) *GroupQuery {
	q.BaseQuery.Where(ceous.Eq(tests.Schema.Group.ID, value))
	return q
}

// ByName add a filter by `Name`.
func (q *GroupQuery) ByName(value string) *GroupQuery {
	q.BaseQuery.Where(ceous.Eq(tests.Schema.Group.Name, value))
	return q
}

// Select defines what fields should be selected from the database.
func (q *GroupQuery) Select(fields ...ceous.SchemaField) *GroupQuery {
	q.BaseQuery.Select(fields...)
	return q
}

// ExcludeFields defines what fields should not be selected from the database.
func (q *GroupQuery) ExcludeFields(fields ...ceous.SchemaField) *GroupQuery {
	q.BaseQuery.ExcludeFields(fields...)
	return q
}

// Where defines the conditions for
func (q *GroupQuery) Where(pred interface{}, args ...interface{}) *GroupQuery {
	q.BaseQuery.Where(pred, args...)
	return q
}

func (q *GroupQuery) Limit(limit uint64) *GroupQuery {
	q.BaseQuery.Limit(limit)
	return q
}

func (q *GroupQuery) Offset(offset uint64) *GroupQuery {
	q.BaseQuery.Offset(offset)
	return q
}

// One results only one record matching the query.
func (q *GroupQuery) One() (m tests.Group, err error) {
	q.Limit(1).Offset(0)

	query, err := q.RawQuery()
	if err != nil {
		return
	}

	rs := NewGroupResultSet(query, nil)
	defer rs.Close()

	if rs.Next() {
		err = rs.ToModel(&m)
		if err != nil {
			return
		}

		for _, rel := range q.BaseQuery.Relations {
			err = rel.Aggregate(&m)
			if err != nil {
				return tests.Group{}, err // TODO(jota): Shall this error be wrapped? At first, yes.
			}
		}
	} else {
		err = ceous.ErrNotFound
	}

	if err == nil {
		for _, rel := range q.BaseQuery.Relations {
			err = rel.Realize()
			if err != nil {
				return tests.Group{}, err // TODO(jota): Shall this error be wrapped? At first, yes.
			}
		}
	}

	return
}

// All return all records that match the query.
func (q *GroupQuery) All() ([]*tests.Group, error) {
	query, err := q.RawQuery()
	if err != nil {
		return nil, err
	}

	rs := NewGroupResultSet(query, nil)
	defer rs.Close()

	ms := make([]*tests.Group, 0)
	for rs.Next() {
		m := &tests.Group{}
		err = rs.ToModel(m)
		if err != nil {
			return nil, err
		}

		for _, rel := range q.BaseQuery.Relations {
			err = rel.Aggregate(m)
			if err != nil {
				return nil, err // TODO(jota): Shall this error be wrapped? At first, yes.
			}
		}
		ms = append(ms, m)
	}

	for _, rel := range q.BaseQuery.Relations {
		err = rel.Realize()
		if err != nil {
			return nil, err // TODO(jota): Shall this error be wrapped? At first, yes.
		}
	}

	return ms, nil
}

func (q *GroupQuery) OrderBy(fields ...interface{}) *GroupQuery {
	q.BaseQuery.OrderBy(fields...)
	return q
}

// UserGroupQuery is the query for the model `UserGroupQuery`.
type UserGroupQuery struct {
	*ceous.BaseQuery
}

// NewUserGroupQuery creates a new query for model `UserGroupQuery`.
func NewUserGroupQuery(options ...ceous.CeousOption) *UserGroupQuery {
	bq := ceous.NewBaseQuery(options...)
	if bq.Schema == nil {
		bq.Schema = tests.Schema.UserGroup.BaseSchema
	}
	return &UserGroupQuery{
		BaseQuery: bq,
	}
}

// ByUserID add a filter by `UserID`.
func (q *UserGroupQuery) ByIDUserID(value int) *UserGroupQuery {
	q.BaseQuery.Where(ceous.Eq(tests.Schema.UserGroup.ID.UserID, value))
	return q
}

// ByGroupID add a filter by `GroupID`.
func (q *UserGroupQuery) ByIDGroupID(value int) *UserGroupQuery {
	q.BaseQuery.Where(ceous.Eq(tests.Schema.UserGroup.ID.GroupID, value))
	return q
}

// ByAdmin add a filter by `Admin`.
func (q *UserGroupQuery) ByAdmin(value bool) *UserGroupQuery {
	q.BaseQuery.Where(ceous.Eq(tests.Schema.UserGroup.Admin, value))
	return q
}

// Select defines what fields should be selected from the database.
func (q *UserGroupQuery) Select(fields ...ceous.SchemaField) *UserGroupQuery {
	q.BaseQuery.Select(fields...)
	return q
}

// ExcludeFields defines what fields should not be selected from the database.
func (q *UserGroupQuery) ExcludeFields(fields ...ceous.SchemaField) *UserGroupQuery {
	q.BaseQuery.ExcludeFields(fields...)
	return q
}

// Where defines the conditions for
func (q *UserGroupQuery) Where(pred interface{}, args ...interface{}) *UserGroupQuery {
	q.BaseQuery.Where(pred, args...)
	return q
}

func (q *UserGroupQuery) Limit(limit uint64) *UserGroupQuery {
	q.BaseQuery.Limit(limit)
	return q
}

func (q *UserGroupQuery) Offset(offset uint64) *UserGroupQuery {
	q.BaseQuery.Offset(offset)
	return q
}

// One results only one record matching the query.
func (q *UserGroupQuery) One() (m tests.UserGroup, err error) {
	q.Limit(1).Offset(0)

	query, err := q.RawQuery()
	if err != nil {
		return
	}

	rs := NewUserGroupResultSet(query, nil)
	defer rs.Close()

	if rs.Next() {
		err = rs.ToModel(&m)
		if err != nil {
			return
		}

		for _, rel := range q.BaseQuery.Relations {
			err = rel.Aggregate(&m)
			if err != nil {
				return tests.UserGroup{}, err // TODO(jota): Shall this error be wrapped? At first, yes.
			}
		}
	} else {
		err = ceous.ErrNotFound
	}

	if err == nil {
		for _, rel := range q.BaseQuery.Relations {
			err = rel.Realize()
			if err != nil {
				return tests.UserGroup{}, err // TODO(jota): Shall this error be wrapped? At first, yes.
			}
		}
	}

	return
}

// All return all records that match the query.
func (q *UserGroupQuery) All() ([]*tests.UserGroup, error) {
	query, err := q.RawQuery()
	if err != nil {
		return nil, err
	}

	rs := NewUserGroupResultSet(query, nil)
	defer rs.Close()

	ms := make([]*tests.UserGroup, 0)
	for rs.Next() {
		m := &tests.UserGroup{}
		err = rs.ToModel(m)
		if err != nil {
			return nil, err
		}

		for _, rel := range q.BaseQuery.Relations {
			err = rel.Aggregate(m)
			if err != nil {
				return nil, err // TODO(jota): Shall this error be wrapped? At first, yes.
			}
		}
		ms = append(ms, m)
	}

	for _, rel := range q.BaseQuery.Relations {
		err = rel.Realize()
		if err != nil {
			return nil, err // TODO(jota): Shall this error be wrapped? At first, yes.
		}
	}

	return ms, nil
}

func (q *UserGroupQuery) OrderBy(fields ...interface{}) *UserGroupQuery {
	q.BaseQuery.OrderBy(fields...)
	return q
}

type UserGroupUserRelation struct {
	_runner ceous.DBRunner
	keys    []interface{}
	records map[int][]*tests.UserGroup
}

func NewUserGroupUserRelation(runner ceous.DBRunner) *UserGroupUserRelation {
	return &UserGroupUserRelation{
		_runner: runner,
		keys:    make([]interface{}, 0),
		records: make(map[int][]*tests.UserGroup),
	}
}

func (relation *UserGroupUserRelation) Aggregate(record ceous.Record) error {
	ugRecord, ok := record.(*tests.UserGroup)
	if !ok {
		return ceous.ErrInvalidRecordType
	}
	if rs, ok := relation.records[ugRecord.ID.UserID]; ok {
		relation.records[ugRecord.ID.UserID] = append(rs, ugRecord)
		// No need to add the key here, since its will be already in the `keys`.
	} else {
		relation.records[ugRecord.ID.UserID] = append(rs, ugRecord)
		relation.keys = append(relation.keys, ugRecord.ID.UserID)
	}
	return nil
}

func (relation *UserGroupUserRelation) Realize() error {
	records, err := NewUserQuery(ceous.WithRunner(relation._runner)).Where(ceous.Eq(tests.Schema.User.ID, relation.keys)).All()
	if err != nil {
		return err // TODO(jota): Shall this be wrapped into a custom error?
	}
	for _, record := range records {
		masterRecords, ok := relation.records[record.ID]
		if !ok {
			return ceous.ErrInconsistentRelationResult
		}
		for _, r := range masterRecords {
			r.AssignUser(record)
		}
	}
	return nil
}

func (q *UserGroupQuery) WithUser() *UserGroupQuery {
	q.BaseQuery.Relations = append(q.BaseQuery.Relations, NewUserGroupUserRelation(q.BaseQuery.Runner))
	return q
}

// UserStore is the query for the store `User`.
type UserStore struct {
	*ceous.BaseStore
}

// NewUserStore creates a new query for model `User`.
func NewUserStore(options ...ceous.CeousOption) *UserStore {
	return &UserStore{
		BaseStore: ceous.NewStore(tests.Schema.User, options...),
	}
}

func (store *UserStore) Insert(record *tests.User, fields ...ceous.SchemaField) error {
	return store.BaseStore.Insert(record, fields...)
}

func (store *UserStore) Update(record *tests.User, fields ...ceous.SchemaField) (int64, error) {
	return store.BaseStore.Update(record, fields...)
}

// GroupStore is the query for the store `Group`.
type GroupStore struct {
	*ceous.BaseStore
}

// NewGroupStore creates a new query for model `Group`.
func NewGroupStore(options ...ceous.CeousOption) *GroupStore {
	return &GroupStore{
		BaseStore: ceous.NewStore(tests.Schema.Group, options...),
	}
}

func (store *GroupStore) Insert(record *tests.Group, fields ...ceous.SchemaField) error {
	return store.BaseStore.Insert(record, fields...)
}

func (store *GroupStore) Update(record *tests.Group, fields ...ceous.SchemaField) (int64, error) {
	return store.BaseStore.Update(record, fields...)
}

// UserGroupStore is the query for the store `UserGroup`.
type UserGroupStore struct {
	*ceous.BaseStore
}

// NewUserGroupStore creates a new query for model `UserGroup`.
func NewUserGroupStore(options ...ceous.CeousOption) *UserGroupStore {
	return &UserGroupStore{
		BaseStore: ceous.NewStore(tests.Schema.UserGroup, options...),
	}
}

func (store *UserGroupStore) Insert(record *tests.UserGroup, fields ...ceous.SchemaField) error {
	return store.BaseStore.Insert(record, fields...)
}

func (store *UserGroupStore) Update(record *tests.UserGroup, fields ...ceous.SchemaField) (int64, error) {
	return store.BaseStore.Update(record, fields...)
}

type userGroupPKResultSet struct {
	*ceous.RecordResultSet
}

func NewUserGroupPKResultSet(rs ceous.ResultSet, err error) *userGroupPKResultSet {
	return &userGroupPKResultSet{
		RecordResultSet: ceous.NewRecordResultSet(rs, err),
	}
}

type userAddressResultSet struct {
	*ceous.RecordResultSet
}

func NewUserAddressResultSet(rs ceous.ResultSet, err error) *userAddressResultSet {
	return &userAddressResultSet{
		RecordResultSet: ceous.NewRecordResultSet(rs, err),
	}
}

type userWorkResultSet struct {
	*ceous.RecordResultSet
}

func NewUserWorkResultSet(rs ceous.ResultSet, err error) *userWorkResultSet {
	return &userWorkResultSet{
		RecordResultSet: ceous.NewRecordResultSet(rs, err),
	}
}

type userResultSet struct {
	*ceous.RecordResultSet
}

func NewUserResultSet(rs ceous.ResultSet, err error) *userResultSet {
	return &userResultSet{
		RecordResultSet: ceous.NewRecordResultSet(rs, err),
	}
}

type groupResultSet struct {
	*ceous.RecordResultSet
}

func NewGroupResultSet(rs ceous.ResultSet, err error) *groupResultSet {
	return &groupResultSet{
		RecordResultSet: ceous.NewRecordResultSet(rs, err),
	}
}

type userGroupIDResultSet struct {
	*ceous.RecordResultSet
}

func NewUserGroupIDResultSet(rs ceous.ResultSet, err error) *userGroupIDResultSet {
	return &userGroupIDResultSet{
		RecordResultSet: ceous.NewRecordResultSet(rs, err),
	}
}

type userGroupResultSet struct {
	*ceous.RecordResultSet
}

func NewUserGroupResultSet(rs ceous.ResultSet, err error) *userGroupResultSet {
	return &userGroupResultSet{
		RecordResultSet: ceous.NewRecordResultSet(rs, err),
	}
}
