// Code generated by ent, DO NOT EDIT.

package ent

import (
	"app/repository/ent/predicate"
	"app/repository/ent/user"
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeUser = "User"
)

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *int
	create_time   *time.Time
	update_time   *time.Time
	hash_id       *string
	mobile        *string
	nickname      *string
	bio           *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*User, error)
	predicates    []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id int) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreateTime sets the "create_time" field.
func (m *UserMutation) SetCreateTime(t time.Time) {
	m.create_time = &t
}

// CreateTime returns the value of the "create_time" field in the mutation.
func (m *UserMutation) CreateTime() (r time.Time, exists bool) {
	v := m.create_time
	if v == nil {
		return
	}
	return *v, true
}

// OldCreateTime returns the old "create_time" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldCreateTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreateTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreateTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreateTime: %w", err)
	}
	return oldValue.CreateTime, nil
}

// ResetCreateTime resets all changes to the "create_time" field.
func (m *UserMutation) ResetCreateTime() {
	m.create_time = nil
}

// SetUpdateTime sets the "update_time" field.
func (m *UserMutation) SetUpdateTime(t time.Time) {
	m.update_time = &t
}

// UpdateTime returns the value of the "update_time" field in the mutation.
func (m *UserMutation) UpdateTime() (r time.Time, exists bool) {
	v := m.update_time
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdateTime returns the old "update_time" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldUpdateTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdateTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdateTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdateTime: %w", err)
	}
	return oldValue.UpdateTime, nil
}

// ResetUpdateTime resets all changes to the "update_time" field.
func (m *UserMutation) ResetUpdateTime() {
	m.update_time = nil
}

// SetHashID sets the "hash_id" field.
func (m *UserMutation) SetHashID(s string) {
	m.hash_id = &s
}

// HashID returns the value of the "hash_id" field in the mutation.
func (m *UserMutation) HashID() (r string, exists bool) {
	v := m.hash_id
	if v == nil {
		return
	}
	return *v, true
}

// OldHashID returns the old "hash_id" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldHashID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldHashID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldHashID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldHashID: %w", err)
	}
	return oldValue.HashID, nil
}

// ResetHashID resets all changes to the "hash_id" field.
func (m *UserMutation) ResetHashID() {
	m.hash_id = nil
}

// SetMobile sets the "mobile" field.
func (m *UserMutation) SetMobile(s string) {
	m.mobile = &s
}

// Mobile returns the value of the "mobile" field in the mutation.
func (m *UserMutation) Mobile() (r string, exists bool) {
	v := m.mobile
	if v == nil {
		return
	}
	return *v, true
}

// OldMobile returns the old "mobile" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldMobile(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldMobile is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldMobile requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMobile: %w", err)
	}
	return oldValue.Mobile, nil
}

// ResetMobile resets all changes to the "mobile" field.
func (m *UserMutation) ResetMobile() {
	m.mobile = nil
}

// SetNickname sets the "nickname" field.
func (m *UserMutation) SetNickname(s string) {
	m.nickname = &s
}

// Nickname returns the value of the "nickname" field in the mutation.
func (m *UserMutation) Nickname() (r string, exists bool) {
	v := m.nickname
	if v == nil {
		return
	}
	return *v, true
}

// OldNickname returns the old "nickname" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldNickname(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldNickname is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldNickname requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldNickname: %w", err)
	}
	return oldValue.Nickname, nil
}

// ResetNickname resets all changes to the "nickname" field.
func (m *UserMutation) ResetNickname() {
	m.nickname = nil
}

// SetBio sets the "bio" field.
func (m *UserMutation) SetBio(s string) {
	m.bio = &s
}

// Bio returns the value of the "bio" field in the mutation.
func (m *UserMutation) Bio() (r string, exists bool) {
	v := m.bio
	if v == nil {
		return
	}
	return *v, true
}

// OldBio returns the old "bio" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldBio(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldBio is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldBio requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldBio: %w", err)
	}
	return oldValue.Bio, nil
}

// ResetBio resets all changes to the "bio" field.
func (m *UserMutation) ResetBio() {
	m.bio = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the UserMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *UserMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.User, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *UserMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 6)
	if m.create_time != nil {
		fields = append(fields, user.FieldCreateTime)
	}
	if m.update_time != nil {
		fields = append(fields, user.FieldUpdateTime)
	}
	if m.hash_id != nil {
		fields = append(fields, user.FieldHashID)
	}
	if m.mobile != nil {
		fields = append(fields, user.FieldMobile)
	}
	if m.nickname != nil {
		fields = append(fields, user.FieldNickname)
	}
	if m.bio != nil {
		fields = append(fields, user.FieldBio)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldCreateTime:
		return m.CreateTime()
	case user.FieldUpdateTime:
		return m.UpdateTime()
	case user.FieldHashID:
		return m.HashID()
	case user.FieldMobile:
		return m.Mobile()
	case user.FieldNickname:
		return m.Nickname()
	case user.FieldBio:
		return m.Bio()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldCreateTime:
		return m.OldCreateTime(ctx)
	case user.FieldUpdateTime:
		return m.OldUpdateTime(ctx)
	case user.FieldHashID:
		return m.OldHashID(ctx)
	case user.FieldMobile:
		return m.OldMobile(ctx)
	case user.FieldNickname:
		return m.OldNickname(ctx)
	case user.FieldBio:
		return m.OldBio(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldCreateTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreateTime(v)
		return nil
	case user.FieldUpdateTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdateTime(v)
		return nil
	case user.FieldHashID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetHashID(v)
		return nil
	case user.FieldMobile:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMobile(v)
		return nil
	case user.FieldNickname:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetNickname(v)
		return nil
	case user.FieldBio:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetBio(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldCreateTime:
		m.ResetCreateTime()
		return nil
	case user.FieldUpdateTime:
		m.ResetUpdateTime()
		return nil
	case user.FieldHashID:
		m.ResetHashID()
		return nil
	case user.FieldMobile:
		m.ResetMobile()
		return nil
	case user.FieldNickname:
		m.ResetNickname()
		return nil
	case user.FieldBio:
		m.ResetBio()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown User edge %s", name)
}
