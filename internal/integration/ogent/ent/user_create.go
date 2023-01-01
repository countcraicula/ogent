// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"ariga.io/ogent/internal/integration/ogent/ent/pet"
	"ariga.io/ogent/internal/integration/ogent/ent/schema"
	"ariga.io/ogent/internal/integration/ogent/ent/user"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (uc *UserCreate) SetName(s string) *UserCreate {
	uc.mutation.SetName(s)
	return uc
}

// SetAge sets the "age" field.
func (uc *UserCreate) SetAge(u uint) *UserCreate {
	uc.mutation.SetAge(u)
	return uc
}

// SetHeight sets the "height" field.
func (uc *UserCreate) SetHeight(u uint) *UserCreate {
	uc.mutation.SetHeight(u)
	return uc
}

// SetNillableHeight sets the "height" field if the given value is not nil.
func (uc *UserCreate) SetNillableHeight(u *uint) *UserCreate {
	if u != nil {
		uc.SetHeight(*u)
	}
	return uc
}

// SetFavoriteCatBreed sets the "favorite_cat_breed" field.
func (uc *UserCreate) SetFavoriteCatBreed(ucb user.FavoriteCatBreed) *UserCreate {
	uc.mutation.SetFavoriteCatBreed(ucb)
	return uc
}

// SetFavoriteDogBreed sets the "favorite_dog_breed" field.
func (uc *UserCreate) SetFavoriteDogBreed(udb user.FavoriteDogBreed) *UserCreate {
	uc.mutation.SetFavoriteDogBreed(udb)
	return uc
}

// SetNillableFavoriteDogBreed sets the "favorite_dog_breed" field if the given value is not nil.
func (uc *UserCreate) SetNillableFavoriteDogBreed(udb *user.FavoriteDogBreed) *UserCreate {
	if udb != nil {
		uc.SetFavoriteDogBreed(*udb)
	}
	return uc
}

// SetFavoriteFishBreed sets the "favorite_fish_breed" field.
func (uc *UserCreate) SetFavoriteFishBreed(sb schema.FishBreed) *UserCreate {
	uc.mutation.SetFavoriteFishBreed(sb)
	return uc
}

// SetNillableFavoriteFishBreed sets the "favorite_fish_breed" field if the given value is not nil.
func (uc *UserCreate) SetNillableFavoriteFishBreed(sb *schema.FishBreed) *UserCreate {
	if sb != nil {
		uc.SetFavoriteFishBreed(*sb)
	}
	return uc
}

// AddPetIDs adds the "pets" edge to the Pet entity by IDs.
func (uc *UserCreate) AddPetIDs(ids ...int) *UserCreate {
	uc.mutation.AddPetIDs(ids...)
	return uc
}

// AddPets adds the "pets" edges to the Pet entity.
func (uc *UserCreate) AddPets(p ...*Pet) *UserCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uc.AddPetIDs(ids...)
}

// SetBestFriendID sets the "best_friend" edge to the User entity by ID.
func (uc *UserCreate) SetBestFriendID(id int) *UserCreate {
	uc.mutation.SetBestFriendID(id)
	return uc
}

// SetNillableBestFriendID sets the "best_friend" edge to the User entity by ID if the given value is not nil.
func (uc *UserCreate) SetNillableBestFriendID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetBestFriendID(*id)
	}
	return uc
}

// SetBestFriend sets the "best_friend" edge to the User entity.
func (uc *UserCreate) SetBestFriend(u *User) *UserCreate {
	return uc.SetBestFriendID(u.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	return withHooks[*User, UserMutation](ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "User.name"`)}
	}
	if _, ok := uc.mutation.Age(); !ok {
		return &ValidationError{Name: "age", err: errors.New(`ent: missing required field "User.age"`)}
	}
	if _, ok := uc.mutation.FavoriteCatBreed(); !ok {
		return &ValidationError{Name: "favorite_cat_breed", err: errors.New(`ent: missing required field "User.favorite_cat_breed"`)}
	}
	if v, ok := uc.mutation.FavoriteCatBreed(); ok {
		if err := user.FavoriteCatBreedValidator(v); err != nil {
			return &ValidationError{Name: "favorite_cat_breed", err: fmt.Errorf(`ent: validator failed for field "User.favorite_cat_breed": %w`, err)}
		}
	}
	if v, ok := uc.mutation.FavoriteDogBreed(); ok {
		if err := user.FavoriteDogBreedValidator(v); err != nil {
			return &ValidationError{Name: "favorite_dog_breed", err: fmt.Errorf(`ent: validator failed for field "User.favorite_dog_breed": %w`, err)}
		}
	}
	if v, ok := uc.mutation.FavoriteFishBreed(); ok {
		if err := user.FavoriteFishBreedValidator(v); err != nil {
			return &ValidationError{Name: "favorite_fish_breed", err: fmt.Errorf(`ent: validator failed for field "User.favorite_fish_breed": %w`, err)}
		}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		}
	)
	if value, ok := uc.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := uc.mutation.Age(); ok {
		_spec.SetField(user.FieldAge, field.TypeUint, value)
		_node.Age = value
	}
	if value, ok := uc.mutation.Height(); ok {
		_spec.SetField(user.FieldHeight, field.TypeUint, value)
		_node.Height = value
	}
	if value, ok := uc.mutation.FavoriteCatBreed(); ok {
		_spec.SetField(user.FieldFavoriteCatBreed, field.TypeEnum, value)
		_node.FavoriteCatBreed = value
	}
	if value, ok := uc.mutation.FavoriteDogBreed(); ok {
		_spec.SetField(user.FieldFavoriteDogBreed, field.TypeEnum, value)
		_node.FavoriteDogBreed = value
	}
	if value, ok := uc.mutation.FavoriteFishBreed(); ok {
		_spec.SetField(user.FieldFavoriteFishBreed, field.TypeEnum, value)
		_node.FavoriteFishBreed = value
	}
	if nodes := uc.mutation.PetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PetsTable,
			Columns: []string{user.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.BestFriendIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.BestFriendTable,
			Columns: []string{user.BestFriendColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_best_friend = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
